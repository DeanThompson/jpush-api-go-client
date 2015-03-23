package httplib

import (
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// HTTP methods we support
const (
	GET  = "GET"
	POST = "POST"
)

type HTTPClient struct {
	client    *http.Client
	transport *http.Transport
	debug     bool
}

func New() *HTTPClient {
	c := &HTTPClient{
		client:    &http.Client{},
		transport: &http.Transport{MaxIdleConnsPerHost: 10},
	}
	c.client.Transport = c.transport
	return c
}

func (c *HTTPClient) Debug(debug bool) *HTTPClient {
	c.debug = debug
	return c
}

func (c *HTTPClient) SetTimeout(connTimeout, rwTimeout time.Duration) *HTTPClient {
	dialer := &net.Dialer{
		Timeout:   connTimeout,
		KeepAlive: 30 * time.Second,
	}

	c.transport.Dial = func(network, addr string) (net.Conn, error) {
		conn, err := dialer.Dial(network, addr)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, nil
	}
	return c
}

func (c *HTTPClient) do(method string, url string, headers map[string]string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	ret, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return ret, nil
}

//func (c *HTTPClient) Get(url string, params map[string]interface{}, headers map[string]string) ([]byte, error) {
//
//}

func valueToString(data interface{}) string {
	value := reflect.ValueOf(data)
	switch value.Kind() {
	case reflect.Bool:
		return strconv.FormatBool(value.Bool())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(value.Uint(), 10)

	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(value.Float(), 'g', -1, 64)

	case reflect.String:
		return value.String()
	}

	return ""
}

func mapToURLValues(data map[string]interface{}) url.Values {
	values := url.Values{}
	for k, v := range data {
		values.Set(k, valueToString(v))
	}
	return values
}

func formatUrl(base string, params map[string]interface{}) string {
	if params == nil || len(params) == 0 {
		return base
	}

	if !strings.Contains(base, "?") {
		base += "?"
	}

	queryString := mapToURLValues(params).Encode()

	if strings.HasSuffix(base, "?") || strings.HasSuffix(url, "&") {
		return base + queryString
	}

	return base + "&" + queryString
}
