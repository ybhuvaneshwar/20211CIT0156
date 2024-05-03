package server_test

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/nikimanoledaki/calculator-microservice/pkg/server"
	protos "github.com/nikimanoledaki/calculator-microservice/protos/calculator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server", func() {

	var (
		cs     *server.CalculatorServer
		logger hclog.Logger
	)

	BeforeEach(func() {
		logger = hclog.Default()
		cs = server.NewComputation(logger)
	})

	Context("Function NewComputation", func() {

		It("returns the operator and no error", func() {
			var cs2 *server.CalculatorServer
			cs2 = new(server.CalculatorServer)
			Expect(cs).Should(BeAssignableToTypeOf(cs2))
		})
	})

	Context("Function GetSum", func() {

		It("it returns an AverageResponse", func() {
			response, err := cs.GetSum(context.TODO(), &protos.SumRequest{FirstNumber: 1, SecondNumber: 2})
			Expect(response.Result).Should(BeNumerically("==", 3))
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Function GetAverage", func() {

		It("it returns an AverageResponse", func() {
			response, err := cs.GetAverage(context.TODO(), &protos.AverageRequest{FirstNumber: 1, SecondNumber: 2})
			Expect(response.Result).Should(BeNumerically("==", 1.5))
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
