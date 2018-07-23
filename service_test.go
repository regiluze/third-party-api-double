package third_party_api_double_test

import (
	_ "fmt"
	. "github.com/onsi/ginkgo"
	_ "github.com/onsi/gomega"
	. "github.com/regiluze/third-party-api-double"
)

const ()

var _ = Describe("Third party API double service specs", func() {
	var (
		service *Service
	)

	BeforeEach(func() {
		service = NewService()
	})

	Describe("foo", func() {
		Context("var", func() {
			It("kk", func() {
			})
		})
	})

})
