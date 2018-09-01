package third_party_api_double

func NewRequestHub() *RequestHub {

	requestHub := RequestHub{}

	return &requestHub
}

func NewRequest(path string, method string, data map[string]string) Request {
	request := Request{path, method, data}
	return request
}

type compationFunc func(Request) bool

type Request struct {
	Path   string
	Method string
	Data   map[string]string
}

type RequestHub struct {
	Requests []Request
}

func (requestHub *RequestHub) Store(path string, method string, data map[string]string) {
	request := NewRequest(path, method, data)
	newRequestSlice := append(requestHub.Requests, request)
	requestHub.Requests = newRequestSlice
}

func (requestHub *RequestHub) Reset() {
	requestHub.Requests = []Request{}
}

func (requestHub RequestHub) Find_by_path(path string) []Request {
	return requestHub.find_by(func(request Request) bool { return request.Path == path })
}

func (requestHub RequestHub) Find_by_path_and_method(path string, method string) []Request {
	return requestHub.find_by(func(request Request) bool { return request.Path == path && request.Method == method })
}

func (requestHub RequestHub) find_by(compFunc compationFunc) []Request {
	matchRequests := requestHub.Requests[:0]
	for _, request := range requestHub.Requests {
		if compFunc(request) {
			matchRequests = append(matchRequests, request)
		}
	}
	return matchRequests
}
