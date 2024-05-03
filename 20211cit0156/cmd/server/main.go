package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/nikimanoledaki/calculator-microservice/pkg/server"
	protos "github.com/nikimanoledaki/calculator-microservice/protos/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	calculatorService := server.NewComputation(log)
	grpcServer := grpc.NewServer()

	protos.RegisterCalculatorServer(grpcServer, calculatorService)

	reflection.Register(grpcServer)

	var port int
	flag.IntVar(&port, "p", 9092, "Specify port to use. Defaults to 9092.")
	flag.Parse()

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Error("Failed to listen to port %d", "error", port, err)
		os.Exit(1)
	}

	if err := grpcServer.Serve(l); err != nil {
		log.Error("Failed to serve gRPC server over port 9092", "error", err)
		os.Exit(1)
	}
}
