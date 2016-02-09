package standard

import (
	"io"
	"net/http"

	"github.com/labstack/echo/engine"
)

type (
	Request struct {
		request *http.Request
		url     engine.URL
		header  engine.Header
	}
)

func NewRequest(r *http.Request) *Request {
	return &Request{
		request: r,
		url:     &URL{url: r.URL},
		header:  &Header{r.Header},
	}
}

func (r *Request) Object() interface{} {
	return r.request
}

func (r *Request) URL() engine.URL {
	return r.url
}

func (r *Request) Header() engine.Header {
	return r.header
}

func (r *Request) RemoteAddress() string {
	return r.request.RemoteAddr
}

func (r *Request) Method() string {
	return r.request.Method
}

func (r *Request) URI() string {
	return r.request.RequestURI
}

func (r *Request) Body() io.ReadCloser {
	return r.request.Body
}

func (r *Request) FormValue(name string) string {
	return r.request.FormValue(name)
}

func (r *Request) reset(req *http.Request, h engine.Header, u engine.URL) {
	r.request = req
	r.header = h
	r.url = u
}
