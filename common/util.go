package common

import (
	"encoding/base64"
	"net/http"
	"sort"
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

func MinInt(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func UniqString(a []string) []string {
	seen := make(map[string]bool, len(a))
	ret := make([]string, 0, len(a))
	for _, v := range a {
		if !seen[v] {
			ret = append(ret, v)
			seen[v] = true
		}
	}
	return ret
}

func EqualStringSlice(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
