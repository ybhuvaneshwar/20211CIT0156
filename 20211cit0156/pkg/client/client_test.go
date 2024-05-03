package client_test

import (
	"github.com/golang/mock/gomock"
	"github.com/nikimanoledaki/calculator-microservice/pkg/client"
	calc_mock "github.com/nikimanoledaki/calculator-microservice/pkg/client/mock"
	protos "github.com/nikimanoledaki/calculator-microservice/protos/calculator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	var (
		err       error
		args      []string
		operation string
	)

	var _ = Describe("Function ParseArguments", func() {

		JustBeforeEach(func() {
			operation, _, err = client.ParseArguments(args)
		})

		Context("If operator is not `sum` or `average`", func() {

			BeforeEach(func() {
				args = []string{"multiply", "1", "2"}
			})

			It("returns an error saying the operation is not recognized", func() {
				Expect(err).Should(MatchError("operation not recognized"))
				Expect(operation).To(Equal(""))
			})
		})

		Context("If the arguments are not an operator with two digits", func() {

			BeforeEach(func() {
				args = []string{"sum", "1", "2", "3"}
			})

			It("it returns an error", func() {
				Expect(err).Should(MatchError("expected 'sum' or 'average' with 2 numeric values"))
				Expect(operation).To(Equal(""))
			})
		})

		Context("If there is a client, an operator, and two digits", func() {

			BeforeEach(func() {
				args = []string{"sum", "1", "2"}
			})

			It("returns no error and the operator", func() {
				Expect(err).ShouldNot(HaveOccurred())
				Expect(operation).To(Equal("sum"))
			})
		})
	})

	var _ = Describe("gRPC functions", func() {

		var (
			ctrl           *gomock.Controller
			mockCalcClient *calc_mock.MockCalculatorClient
		)

		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			mockCalcClient = calc_mock.NewMockCalculatorClient(ctrl)
			args = []string{"1", "2"}
		})

		AfterEach(func() {
			ctrl.Finish()
		})

		Context("When passing a SumRequest and CalculatorClient to function PrintSum", func() {

			var (
				response *protos.SumResponse
			)

			JustBeforeEach(func() {
				mockCalcClient.EXPECT().GetSum(gomock.Any(), gomock.Any()).Return(&protos.SumResponse{Result: 3}, nil)
				response, err = client.PrintSum(mockCalcClient, args)
			})

			It("it returns a SumResponse", func() {
				Expect(response.Result).Should(BeNumerically("==", 3))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("When passing an AverageRequest and CalculatorClient to function PrintAverage", func() {

			var (
				response *protos.AverageResponse
			)

			JustBeforeEach(func() {
				mockCalcClient.EXPECT().GetAverage(gomock.Any(), gomock.Any()).Return(&protos.AverageResponse{Result: 1.5}, nil)
				response, err = client.PrintAverage(mockCalcClient, args)
			})

			It("it returns an AverageResponse", func() {
				Expect(response.Result).Should(BeNumerically("==", 1.5))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("When calling on function NewRequest", func() {

			JustBeforeEach(func() {
				mockCalcClient.EXPECT().GetSum(gomock.Any(), gomock.Any()).Return(&protos.SumResponse{Result: 3}, nil)
				operation = "sum"
				err = client.NewRequest(operation, mockCalcClient, args)
			})

			It("it returns an AverageResponse", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("When calling on function NewRequest", func() {

			JustBeforeEach(func() {
				mockCalcClient.EXPECT().GetAverage(gomock.Any(), gomock.Any()).Return(&protos.AverageResponse{Result: 1.5}, nil)
				operation = "average"
				err = client.NewRequest(operation, mockCalcClient, args)
			})

			It("it returns an AverageResponse", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})
})
