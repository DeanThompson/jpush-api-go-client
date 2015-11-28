package httplib

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
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
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"

//	PUT    = "PUT"
//	HEAD   = "HEAD"
)

type HTTPClient struct {
	client    *http.Client
	transport *http.Transport
	debug     bool
}

func NewClient() *HTTPClient {
	c := &HTTPClient{
		client:    &http.Client{},
		transport: &http.Transport{MaxIdleConnsPerHost: 10},
	}
	c.SetTimeout(10*time.Second, 10*time.Second)
	c.client.Transport = c.transport
	return c
}

func (c *HTTPClient) SetDebug(debug bool) *HTTPClient {
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

func (c *HTTPClient) do(method string, url string, headers map[string]string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	return c.client.Do(req)
}

func (c *HTTPClient) Get(url string, params map[string]interface{}, headers map[string]string) (*http.Response, error) {
	url = formatUrl(url, params)
	c.log(GET, url, headers)
	return c.do(GET, url, headers, nil)
}

func (c *HTTPClient) Post(url string, bodyType string, body io.Reader, headers map[string]string) (*http.Response, error) {
	// DON'T modify headers passed-in
	new_headers := make(map[string]string, len(headers))
	for k, v := range headers {
		new_headers[k] = v
	}
	new_headers["Content-Type"] = bodyType

	return c.do(POST, url, new_headers, body)
}

func (c *HTTPClient) PostForm(url string, data map[string]interface{}, headers map[string]string) (*http.Response, error) {
	c.log(POST, url, data, headers)

	body := strings.NewReader(mapToURLValues(data).Encode())
	return c.Post(url, "application/x-www-form-urlencoded", body, headers)
}

func (c *HTTPClient) PostJson(url string, data interface{}, headers map[string]string) (*http.Response, error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	c.log(POST, url, string(payload), headers)

	return c.Post(url, "application/json", bytes.NewReader(payload), headers)
}

func (c *HTTPClient) Delete(url string, params map[string]interface{}, headers map[string]string) (*http.Response, error) {
	url = formatUrl(url, params)
	c.log(DELETE, url, headers)
	return c.do(DELETE, url, headers, nil)
}

func (c *HTTPClient) log(v ...interface{}) {
	if c.debug {
		vs := []interface{}{"[HTTPClient]"}
		log.Println(append(vs, v...)...)
	}
}

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

	if strings.HasSuffix(base, "?") || strings.HasSuffix(base, "&") {
		return base + queryString
	}

	return base + "&" + queryString
}
