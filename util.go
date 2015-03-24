package jpush

import (
	"encoding/base64"
	"net/http"
	"strconv"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func getIntHeader(resp *http.Response, key string) (int, error) {
	v := resp.Header.Get(key)
	return strconv.Atoi(v)
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
