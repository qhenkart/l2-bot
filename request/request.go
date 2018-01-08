package request

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/qhenkart/l2bot/errs"
)

// Request handles request information for inter-service requests
type Request struct {
	Header   http.Header
	Method   string
	Body     interface{}
	response interface{}
	Endpoint string
	ctx      context.Context
}

// Response handles response information for inter-service requests
type Response struct {
	Body       interface{}
	StatusCode int
	Header     http.Header
}

// NewRequest initializes the Request struct and primes a new request
// These methods are meant to be chained together as needed
func NewRequest(rawURL string, header ...http.Header) *Request {
	req := &Request{Endpoint: rawURL, Method: "GET"}
	if len(header) != 0 {
		req.Header = header[0]
	}
	return req
}

// WithMethod defines the method of the Request, the Default is GET
func (r *Request) WithMethod(method string) *Request {
	r.Method = method
	return r
}

// WithBody provides a body for the payload if necessary
func (r *Request) WithBody(in interface{}) *Request {
	r.Body = in
	return r
}

// WithResponse provides an interface for the response if necessary
func (r *Request) WithResponse(out interface{}) *Request {
	r.response = out
	return r
}

// Do performs the actual request
func (r *Request) Do() error {
	resp := &Response{Body: r.response}
	return HandleRequest(r, resp)
}

// DoWithResponse performs the actual request and returns the response struct which contains metadata information
// such as StatusCode and header
func (r *Request) DoWithResponse() (*Response, error) {
	resp := &Response{Body: r.response}
	return resp, HandleRequest(r, resp)
}

// HandleRequest handles inter-service http requests
func HandleRequest(req *Request, res *Response) error {

	var encodedData io.Reader

	// Marshal the payload if one is included
	if req.Body != nil {
		byteIn, err := json.Marshal(req.Body)
		if err != nil {
			return errs.ErrorWithContext(errs.ErrMarshalData, err)
		}
		encodedData = bytes.NewBuffer(byteIn)
	}

	r, err := http.NewRequest(req.Method, req.Endpoint, encodedData)
	if err != nil {
		return errs.ErrServiceRequest.Context(err)
	}

	if req.Header != nil {
		r.Header = req.Header
	}

	client := &http.Client{}

	resp, err := client.Do(r)
	if err != nil {
		return errs.ErrServiceRequest.Context(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode >= 500 {
		return errs.ErrServiceRequest
	}

	res.StatusCode = resp.StatusCode

	decoder := json.NewDecoder(resp.Body)
	if resp.StatusCode > 399 {
		// fail
		svcError := new(errs.ErrInfo)
		svcError.StatusCode = resp.StatusCode

		return svcError
	}

	if res.Body != nil {
		// success
		if err := decoder.Decode(res.Body); err != nil {
			return errs.ErrorWithContext(errs.ErrDecoder, err)
		}
	}

	res.Header = resp.Header
	return nil
}
