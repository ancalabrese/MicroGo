package main

import (
	"net"
	"os"

	"github.com/ancalabrese/MicroGo/Currency/data"

	protos "github.com/ancalabrese/MicroGo/Currency/protos/currency"
	"github.com/ancalabrese/MicroGo/Currency/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	rates, err := data.NewExchangeRate(log)
	if err != nil {
		log.Error("Unable to generate rates", "error", err)
		os.Exit(1)
	}

	cc := data.NewCurrencies(log)

	grpcs := grpc.NewServer()
	cs := server.NewCurrencyServer(rates, cc, log)

	protos.RegisterCurrencyServer(grpcs, cs)

	//for testing purposes
	reflection.Register(grpcs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Cannot bind service", "port", err)
		os.Exit(1)
	}

	grpcs.Serve(l)
}
