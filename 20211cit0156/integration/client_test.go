package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	protos "github.com/nikimanoledaki/calculator-microservice/protos/calculator"
)

var _ = Describe("Client", func() {

	Context("Function GetSum", func() {

		var (
			sumResp *protos.SumResponse
			err     error
			sumReq  *protos.SumRequest
		)

		BeforeEach(func() {
			sumReq = &protos.SumRequest{FirstNumber: 1, SecondNumber: 2}
			sumResp, err = clt.GetSum(ctx, sumReq)
		})

		It("given two integers, it should return no error and a sum", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(sumResp.Result).Should(BeNumerically("==", 3))
		})

	})

	Context("Function GetAverage", func() {

		var (
			avgResp *protos.AverageResponse
			err     error
			avgReq  *protos.AverageRequest
		)

		BeforeEach(func() {
			avgReq = &protos.AverageRequest{FirstNumber: 1, SecondNumber: 2}
			avgResp, err = clt.GetAverage(ctx, avgReq)
		})

		It("given two floats, it should return no error and a sum", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(avgResp.Result).Should(BeNumerically("==", 1.5))
		})

	})

})
