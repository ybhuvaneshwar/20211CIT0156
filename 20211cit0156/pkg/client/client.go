package client

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/nikimanoledaki/calculator-microservice/protos/calculator"
	protos "github.com/nikimanoledaki/calculator-microservice/protos/calculator"
)

// ParseArguments manages the error handling for the unhappy paths of the client.
func ParseArguments(args []string) (string, []string, error) {
	if len(args) != 3 {
		return "", args, fmt.Errorf("expected 'sum' or 'average' with 2 numeric values")
	}

	operation := args[0]
	if operation != "sum" && operation != "average" {
		return "", args, fmt.Errorf("operation not recognized")
	}

	numbers := args[1:]

	return operation, numbers, nil
}

// NewRequest filters the new requests by operation and returns their response.
func NewRequest(operation string, client protos.CalculatorClient, args []string) error {
	if operation == "sum" {
		response, err := PrintSum(client, args)
		if err != nil {
			return err
		}
		fmt.Println(response)
	} else {
		response, err := PrintAverage(client, args)
		if err != nil {
			return err
		}
		fmt.Println(response)
	}
	return nil
}

// PrintSum receives a type CalculatorClient and command-line arguments to create an SumRequest then log the SumResponse.
func PrintSum(client protos.CalculatorClient, args []string) (*calculator.SumResponse, error) {

	numbers := make([]int32, 2)
	for i, arg := range args {
		number, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalf("failed to convert %d to integer", err)
		}
		numbers[i] = int32(number)
	}

	sumReq := &protos.SumRequest{
		FirstNumber:  numbers[0],
		SecondNumber: numbers[1],
	}

	return client.GetSum(context.Background(), sumReq)
}

// PrintAverage receives a type CalculatorClient and command-line arguments to create an AverageRequest then log the AverageResponse.
func PrintAverage(client protos.CalculatorClient, args []string) (*calculator.AverageResponse, error) {
	numbers := make([]float32, 2)
	for i, arg := range args {
		number, err := strconv.ParseFloat(arg, 32)
		if err != nil {
			log.Fatalf("failed to convert %d to float", err)
		}
		numbers[i] = float32(number)
	}

	avgReq := &protos.AverageRequest{
		FirstNumber:  numbers[0],
		SecondNumber: numbers[1],
	}

	return client.GetAverage(context.Background(), avgReq)
}
