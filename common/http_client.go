package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type IHttpClient interface {
	Request(method, url string, params []KeyValue, body interface{}, header []KeyValue) (error, int, []byte)
	Get(url string, params []KeyValue, body interface{}, header []KeyValue) (error, int, []byte)
	Post(url string, params []KeyValue, body interface{}, header []KeyValue) (error, int, []byte)
	Delete(url string, params []KeyValue, body interface{}, header []KeyValue) (error, int, []byte)
	Put(url string, params []KeyValue, body interface{}, header []KeyValue) (error, int, []byte)
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var HttpCli *HttpClient

func init() {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 500 * time.Millisecond,
		}).DialContext,
		TLSHandshakeTimeout: 1000 * time.Millisecond,
		MaxIdleConns:        100,
		IdleConnTimeout:     30 * time.Second,
	}
	HttpCli = NewHttpClient(transport)
}

func NewHttpClient(transport *http.Transport) *HttpClient {
	if transport == nil {
		transport = http.DefaultTransport.(*http.Transport)
	}

	return &HttpClient{&http.Client{Transport: transport}}
}

type HttpClient struct {
	client *http.Client
}

const ErrorCode = -1

func (h *HttpClient) Request(method, u string, params []KeyValue, requestBody interface{}, headers []KeyValue) (err error, code int, respBody []byte) {
	start := time.Now().UTC()
	code = ErrorCode
	respBody = []byte{}
	err = nil
	defer func() {
		logger.Debugf("###[%s]### cost[%.0f] %s %v %v %v ###[%d]###response###%s", method, float64(time.Since(start))/1000000, u, requestBody, params, headers, code, string(respBody))
	}()

	if len(u) == 0 {
		err = fmt.Errorf("url is empty")
		logger.Debugf(err.Error())
		return
	}

	_, err = url.Parse(u)
	if err != nil {
		logger.Debugf("parse url %s failed, error: %v", u, err)
		return
	}

	var buffer bytes.Buffer
	buffer.WriteString(u)
	if len(params) > 0 && !strings.HasSuffix(u, "?") {
		buffer.WriteString("?")
	}
	for i, param := range params {
		buffer.WriteString(url.QueryEscape(param.Key))
		buffer.WriteString("=")
		buffer.WriteString(url.QueryEscape(param.Value))
		if i < len(params)-1 {
			buffer.WriteString("&")
		}
	}

	var bodyBytes []byte
	if requestBody != nil {
		bodyBytes, err = json.Marshal(requestBody)
		if err != nil {
			logger.Debugf("build request error: %v", err.Error())
			return
		}
	}

	req, err := http.NewRequest(method, buffer.String(), bytes.NewBuffer(bodyBytes))
	if err != nil {
		logger.Debugf("build request error: %v", err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")
	for _, header := range headers {
		req.Header.Set(header.Key, header.Value)
	}

	resp, err := h.client.Do(req)
	if err != nil {
		logger.Debugf("http call error: %v", err.Error())
		return
	}

	bodyBytes, err = ioutil.ReadAll(resp.Body)
	defer closeBody(resp)
	if err != nil {
		return
	} else {
		return nil, resp.StatusCode, bodyBytes
	}
}

func (h *HttpClient) Get(url string, params []KeyValue, body interface{}, header []KeyValue) (error, int, []byte) {
	return h.Request(http.MethodGet, url, params, body, header)
}

func (h *HttpClient) Post(url string, params []KeyValue, body interface{}, header []KeyValue) (error, int, []byte) {
	return h.Request(http.MethodPost, url, params, body, header)
}

func (h *HttpClient) Delete(url string, params []KeyValue, body interface{}, header []KeyValue) (error, int, []byte) {
	return h.Request(http.MethodDelete, url, params, body, header)
}

func (h *HttpClient) Put(url string, params []KeyValue, body interface{}, header []KeyValue) (error, int, []byte) {
	return h.Request(http.MethodPut, url, params, body, header)
}

func closeBody(resp *http.Response) {
	if resp != nil && resp.Body != nil {
		err := resp.Body.Close()
		if err != nil {
			logger.Debugf("response body close error: %v, req: %v", err.Error(), resp.Request)
		}
	}
}
