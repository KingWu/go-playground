package model_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"kw101/go-playground/api"
)

var _ = Describe("Api/Sum", func() {
	var (
		p, q, m, n, sum1, sum2 int
	)
	BeforeEach(func() {
		p, q, sum1 = 5, 6, 11
		// Putting wrong value of sum2 intentionally 
		m, n, sum2 = 8, 7, 19
	})
	Context("Addition of two digits", func() {
		It("should return sum of the two digits", func() {
			addition_of_two_digits := model.Sum(p, q)
			Expect(addition_of_two_digits).Should(Equal(sum1))
		})
		It("should not return the sum provided", func(){
			addition_of_two_digits := model.Sum(m, n)
			Expect(addition_of_two_digits).ShouldNot(Equal(sum2))
		})
	})
})