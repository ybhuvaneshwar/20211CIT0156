package calculator_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCalculatorMicroservice(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CalculatorMicroservice Suite")
}
