package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ancalabrese/MicroGo/Products/data"
	"github.com/hashicorp/go-hclog"

	proto "github.com/ancalabrese/MicroGo/Currency/protos/currency"
	"github.com/ancalabrese/MicroGo/Products/handlers"
	"github.com/ancalabrese/MicroGo/Products/middleware"
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {
	l := hclog.New(&hclog.LoggerOptions{
		Name:  "Products-API",
		Level: hclog.LevelFromString("DEBUG"),
	})

	log.New(os.Stdout, "product-api", log.LstdFlags)

	//CurrecyServer client
	conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	cc := proto.NewCurrencyClient(conn)

	pdb := data.NewProductsDB(l, cc)
	ph := handlers.NewProducts(l, pdb)

	//main product router
	r := mux.NewRouter()
	productsRouter := r.NewRoute().PathPrefix("/products").Subrouter()
	middlewareLogger := middleware.NewLogger(l)
	productsRouter.Use(middlewareLogger.LogIncomingReq)

	validator := middleware.NewProductValidator(l)
	// Sub-router for each suppoprted method
	getRouter := productsRouter.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("", ph.GetProducts)
	getRouter.HandleFunc("/{id:[0-9]+}", ph.GetProduct)

	postRouter := productsRouter.Methods(http.MethodPost).Subrouter()
	postRouter.Use(validator.Validate)
	postRouter.HandleFunc("", ph.AddProducts)

	putRouter := productsRouter.Methods(http.MethodPut).Subrouter()
	putRouter.Use(validator.Validate)
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)

	// deleteRouter := productsRouter.Methods(http.MethodDelete).Subrouter()

	//CORS
	corsHandler := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

	s := &http.Server{
		Addr:         ":9090",
		Handler:      corsHandler(r),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
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
	log.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
