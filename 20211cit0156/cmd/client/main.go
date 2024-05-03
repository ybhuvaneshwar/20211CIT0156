package main

import (
	"flag"
	"log"
	"os"

	client "github.com/nikimanoledaki/calculator-microservice/pkg/client"
	protos "github.com/nikimanoledaki/calculator-microservice/protos/calculator"
	"google.golang.org/grpc"
)

func main() {
	var port string
	flag.StringVar(&port, "p", "localhost:9092", "Specify port to use. Defaults to localhost:9092.")
	flag.Parse()

	operation, numbers, err := client.ParseArguments(flag.Args())
	if err != nil {
		log.Fatalf("Failed to parse arguments: %s", err)
	}

	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to port %s: %s", port, err)
	}
	defer conn.Close()

	clientService := protos.NewCalculatorClient(conn)

	err = client.NewRequest(operation, clientService, numbers)
	if err != nil {
		log.Fatalf("Failed to perform operation %v: %s", operation, err)
		os.Exit(1)
	}

}
