package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc/codes"

	proto "github.com/ancalabrese/MicroGo/Currency/protos/currency"
	cdata "github.com/ancalabrese/MicroGo/Products/data/currency"
	pdata "github.com/ancalabrese/MicroGo/Products/data/product"
	"github.com/ancalabrese/MicroGo/Products/grpcClient"
	cHandler "github.com/ancalabrese/MicroGo/Products/handlers/currencies"
	pHandler "github.com/ancalabrese/MicroGo/Products/handlers/products"
	"github.com/ancalabrese/MicroGo/Products/middleware"
	settings "github.com/ancalabrese/MicroGo/Products/settings"
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	l := hclog.New(&hclog.LoggerOptions{
		Level: hclog.LevelFromString("Debug"),
	})
	//Load server config
	config := settings.NewConfig(l)
	err := config.Load("./config.yml")
	if err != nil {
		l.Debug("Defaults", "Addr", config.SeviceConfig.Url, "Port", config.SeviceConfig.Port, "Base Path", config.SeviceConfig.ApiBasePath)
	}
	l.SetLevel(hclog.LevelFromString(config.GeneralConfig.LogLevel))

	//CurrecyServer client
	//TODO: explore ways of logging retries
	retryPolicy := grpc_retry.BackoffExponential(300 * time.Millisecond)
	callOtp := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(retryPolicy),
		grpc_retry.WithCodes(codes.NotFound, codes.Aborted),
	}
	grpcC := grpcClient.NewClient(l, "localhost", "9092")
	grpcC.WithDialOption(
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(callOtp...)))

	err = grpcC.DialUp()
	if err != nil {
		panic(err)
	}
	defer grpcC.Close()
	cc := proto.NewCurrencyClient(grpcC.ClientConnection)

	pdb := pdata.NewProductsDB(l, cc)
	ph := pHandler.NewProducts(l, pdb)

	cdb := cdata.NewCurrencyDB(l, cc)
	ch := cHandler.NewCurrencyH(l, cdb)

	//main API router
	r := mux.NewRouter()

	//Products router
	productsRouter := r.NewRoute().PathPrefix(config.SeviceConfig.ApiBasePath + "/products").Subrouter()
	middlewareLogger := middleware.NewLogger(l)
	productsRouter.Use(middlewareLogger.LogIncomingReq)

	//Currency router
	currencyRouter := r.NewRoute().PathPrefix(config.SeviceConfig.ApiBasePath + "/currencies").Subrouter()
	currencyRouter.Use(middlewareLogger.LogIncomingReq)

	validator := middleware.NewProductValidator(l)
	// Products: Sub-router for each suppoprted method
	pGetRouter := productsRouter.Methods(http.MethodGet).Subrouter()
	pGetRouter.HandleFunc("", ph.GetProducts).Queries("currency", "{[A-Z]3")
	pGetRouter.HandleFunc("", ph.GetProducts)
	pGetRouter.HandleFunc("/{id:[0-9]+}", ph.GetProduct).Queries("currency", "{[A-Z]3")
	pGetRouter.HandleFunc("/{id:[0-9]+}", ph.GetProduct)

	pPostRouter := productsRouter.Methods(http.MethodPost).Subrouter()
	pPostRouter.Use(validator.Validate)
	pPostRouter.HandleFunc("", ph.AddProducts)

	pPutRouter := productsRouter.Methods(http.MethodPut).Subrouter()
	pPutRouter.Use(validator.Validate)
	pPutRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)

	pDeleteRouter := productsRouter.Methods(http.MethodDelete).Subrouter()
	pDeleteRouter.HandleFunc("/{id:[0-9]+}", ph.DeleteProduct)

	//Currencies: Sub-router for each supported methods
	cGetRouter := currencyRouter.Methods(http.MethodGet).Subrouter()
	cGetRouter.HandleFunc("", ch.GetCurrencies)

	//CORS
	corsHandler := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{config.SeviceConfig.CORSAllowedOrigins}))

	s := &http.Server{
		Addr:         config.SeviceConfig.Url + ":" + config.SeviceConfig.Port,
		Handler:      corsHandler(r),
		ErrorLog:     l.StandardLogger(&hclog.StandardLoggerOptions{}),
		IdleTimeout:  config.SeviceConfig.IdleTimeout * time.Second,
		ReadTimeout:  config.SeviceConfig.ReadTimeout * time.Second,
		WriteTimeout: config.SeviceConfig.WriteTimeout * time.Second,
	}

	go func() {
		l.Info("Starting server", "Address", s.Addr)
		err := s.ListenAndServe()
		if err != nil {
			l.Error("Error starting server", "error", err)
			os.Exit(1)
		}

	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Info("Got system signal", "sig", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
