package third_party_api_double_test

import (
	_ "fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

	Describe("Managing request data", func() {
		Context("when storing a request data", func() {
			It("stores the request info", func() {
				data := map[string]string{
					"data": "foo",
				}
				service.Store("path", "method", data)

				Expect(service.Requests[0].Path).To(Equal("path"))
				Expect(service.Requests[0].Method).To(Equal("method"))
				Expect(service.Requests[0].Data).To(Equal(data))
			})
		})
	})

})
