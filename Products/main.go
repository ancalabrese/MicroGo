package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	cdata "github.com/ancalabrese/MicroGo/Products/data/currency"
	pdata "github.com/ancalabrese/MicroGo/Products/data/product"
	settings "github.com/ancalabrese/MicroGo/Products/settings"
	"github.com/hashicorp/go-hclog"

	proto "github.com/ancalabrese/MicroGo/Currency/protos/currency"
	cHandler "github.com/ancalabrese/MicroGo/Products/handlers/currencies"
	pHandler "github.com/ancalabrese/MicroGo/Products/handlers/products"
	"github.com/ancalabrese/MicroGo/Products/middleware"
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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
	conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	cc := proto.NewCurrencyClient(conn)

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
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
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
