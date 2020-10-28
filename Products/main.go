package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ancalabrese/MicroGo/Currency"
	"github.com/ancalabrese/MicroGo/Images"
	"github.com/ancalabrese/MicroGo/Products/handlers"
	"github.com/ancalabrese/MicroGo/Products/middleware"
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProducts(l)

	//CurrecyServer client
	
	//main product router
	r := mux.NewRouter()
	productsRouter := r.NewRoute().PathPrefix("/products").Subrouter()
	middlewareLogger := middleware.NewLogger(l)
	productsRouter.Use(middlewareLogger.LogIncomingReq)

	// Sub-router for each suppoprted method
	getRouter := productsRouter.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("", ph.GetProducts)

	postRouter := productsRouter.Methods(http.MethodPost).Subrouter()
	postRouter.Use(middleware.Validate)
	postRouter.HandleFunc("", ph.AddProducts)

	putRouter := productsRouter.Methods(http.MethodPut).Subrouter()
	putRouter.Use(middleware.Validate)
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	

	// deleteRouter := productsRouter.Methods(http.MethodDelete).Subrouter()

	//CORS
	corsHandler:= gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

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
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, grecefully shout down", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
