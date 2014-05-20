package utils

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type HttpRequest interface {
	Post(url string, values url.Values) ([]byte, *RequestError)
}

type Request struct{}

func NewRequest() *Request {
	return &Request{}
}

func (ur *Request) Post(url string, values url.Values) ([]byte, *RequestError) {
	client := &http.Client{}
	r, err := http.NewRequest("POST", url, strings.NewReader(values.Encode()))
	if nil != err {
		return nil, NewRequestError(http.StatusInternalServerError, err)
	}

	headers := ur.getHeaders()
	for header, value := range headers {
		r.Header.Set(header, value)
	}

	resp, err := client.Do(r)
	if nil != err {
		return nil, NewRequestError(http.StatusInternalServerError, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return nil, NewRequestError(http.StatusInternalServerError, err)
	}
	if resp.StatusCode != http.StatusOK {
		xErrorCode, _ := strconv.ParseInt(resp.Header.Get("X-Error-Code"), 10, 0)
		xErrorText := resp.Header.Get("X-Error")
		return nil, NewRequestError(int(xErrorCode), errors.New(xErrorText))
	}
	return body, nil
}

func (ur *Request) getHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"X-Accept":     "application/json",
	}
}
