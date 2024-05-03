package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/nikimanoledaki/calculator-microservice/pkg/calculator"
	protos "github.com/nikimanoledaki/calculator-microservice/protos/calculator"
)

// CalculatorServer is a struct used to create methods that implement
// the gRPC interface.
type CalculatorServer struct {
	log hclog.Logger
}

// NewComputation is a constructor that implements the Calculator service interface.
func NewComputation(l hclog.Logger) *CalculatorServer {
	return &CalculatorServer{l}
}

// GetSum implements the gRPC method GetSum by handling the SumRequest and returning the SumResponse.
func (cs *CalculatorServer) GetSum(ctx context.Context, sumReq *protos.SumRequest) (*protos.SumResponse, error) {
	cs.log.Info("Handle GetSum", "firstNumber", sumReq.GetFirstNumber(), "secondNumber", sumReq.GetSecondNumber())

	calculatedSum := calculator.Sum(sumReq.GetFirstNumber(), sumReq.GetSecondNumber())

	return &protos.SumResponse{Result: calculatedSum}, nil
}

// GetAverage implements the gRPC method GetAverage by handling the AverageRequest and returning the AverageResponse.
func (cs *CalculatorServer) GetAverage(ctx context.Context, avgReq *protos.AverageRequest) (*protos.AverageResponse, error) {
	cs.log.Info("Handle GetAverage", "firstNumber", avgReq.GetFirstNumber(), "secondNumber", avgReq.GetSecondNumber())

	calculatedAvg := calculator.Average(avgReq.GetFirstNumber(), avgReq.GetSecondNumber())

	return &protos.AverageResponse{Result: calculatedAvg}, nil
}
