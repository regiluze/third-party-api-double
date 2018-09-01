package third_party_api_double

func NewService() *Service {

	service := Service{}

	return &service
}

func NewRequest(path string, method string, data map[string]string) Request {
	request := Request{path, method, data}
	return request
}

type Request struct {
	Path   string
	Method string
	Data   map[string]string
}

type Service struct {
	Requests []Request
}

func (service *Service) Store(path string, method string, data map[string]string) {
	request := NewRequest(path, method, data)
	newRequestSlice := append(service.Requests, request)
	service.Requests = newRequestSlice
}
