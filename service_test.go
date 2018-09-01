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
		requestHub *RequestHub
	)

	BeforeEach(func() {
		requestHub = NewRequestHub()
	})

	Describe("Managing request data", func() {
		Context("when storing a request data", func() {
			It("stores the request info", func() {
				data := map[string]string{
					"data": "foo",
				}
				requestHub.Store("path", "method", data)

				Expect(requestHub.Requests[0].Path).To(Equal("path"))
				Expect(requestHub.Requests[0].Method).To(Equal("method"))
				Expect(requestHub.Requests[0].Data).To(Equal(data))
			})
		})
		Context("when reseting the request hub", func() {
			It("set requests list to empty list", func() {
				data := map[string]string{
					"data": "foo",
				}
				request := NewRequest("path", "method", data)
				requestHub.Requests = append(requestHub.Requests, request)

				requestHub.Reset()

				Expect(requestHub.Requests).To(BeEmpty())
			})
		})
		Context("when finding the request by path", func() {
			It("returns a slice of match requests", func() {
				data := map[string]string{
					"data": "foo",
				}
				request1 := NewRequest("path1", "method", data)
				request2 := NewRequest("path2", "method", data)
				requestHub.Requests = append(requestHub.Requests, request1)
				requestHub.Requests = append(requestHub.Requests, request2)

				requests := requestHub.Find_by_path("path1")

				Expect(requests).NotTo(BeEmpty())
				Expect(requests[0].Path).To(Equal("path1"))
			})
		})
	})

})
