package main

import (
	"net"
	"os"

	protos "github.com/ancalabrese/EXPerimenta/GoMic/Currency/protos/currency"
	"github.com/ancalabrese/EXPerimenta/GoMic/Currency/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	grpcs := grpc.NewServer()
	cs := server.NewCurrencyServer(log)

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
