package common

import (
	"encoding/base64"
	"net/http"
	"strconv"
)

func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func GetIntHeader(resp *http.Response, key string) (int, error) {
	v := resp.Header.Get(key)
	return strconv.Atoi(v)
}

func MaxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
