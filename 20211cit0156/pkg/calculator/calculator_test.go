package calculator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nikimanoledaki/calculator-microservice/pkg/calculator"
)

var _ = Describe("Calculator", func() {
	Context("Can do basic arithmetic operations such as", func() {
		It("adding two numbers", func() {
			Expect(calculator.Sum(1, 1)).Should(BeNumerically("==", 2))
		})
		It("finding the average of two floats", func() {
			Expect(calculator.Average(0, 1)).Should(BeNumerically("<", 1))
		})
	})
})
