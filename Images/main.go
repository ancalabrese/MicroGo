package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ancalabrese/MicroGo/Images/middleware"

	"github.com/ancalabrese/MicroGo/Images/file"

	"github.com/ancalabrese/MicroGo/Images/handlers"
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

var bindAddress = ":9091"
var logLevel = "DEBUG"
var basePath = "./imgstore"
var maxImgSize = 1 * 1024 * 2

func main() {
	l := hclog.New(&hclog.LoggerOptions{
		Name:  "product-images",
		Level: hclog.LevelFromString("DEBUG"),
	})

	//Logger for the server instance
	sl := l.StandardLogger(&hclog.StandardLoggerOptions{
		InferLevels: true,
	})

	ls, err := file.NewLocalStorage(basePath, maxImgSize)
	if err != nil {
		l.Error("Couldn't create internal storage: ", err)
		os.Exit(1)
	}

	fh := handlers.NewFile(l, ls)

	r := mux.NewRouter()
	gZipperMW := middleware.Gzipper{}

	imageRouter := r.PathPrefix("/image").Subrouter()

	imgPostSubRouter := imageRouter.Methods(http.MethodPost).Subrouter()
	fh = handlers.NewFile(l, ls)
	imgPostSubRouter.HandleFunc("/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", fh.UploadREST)
	imgPostSubRouter.HandleFunc("/", fh.UploadMultiPart)

	imgGetSubRouter := imageRouter.Methods(http.MethodGet).Subrouter()
	imgGetSubRouter.Use(gZipperMW.GzipperMiddleware)
	imgGetSubRouter.Handle("/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}",
		http.StripPrefix("/image/", http.FileServer(http.Dir(basePath))))

	corsHandler := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

	s := &http.Server{
		Addr:         bindAddress,
		Handler:      corsHandler(r),
		ErrorLog:     sl,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		l.Info("Starting server...", "Binding addres", bindAddress)
		err := s.ListenAndServe()
		if err != nil {
			l.Error("Couldn't Start the server: ", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	//Block until signal is received
	sig := <-c
	l.Info("Shutting down server with signal", "sig", sig)
	//Gracefully shutdown server
	context.WithTimeout(context.Background(), 30*time.Second)
}
