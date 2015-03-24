package jpush

import (
	"github.com/DeanThompson/jpush-api-go-client/httplib"
)

type Validator interface {
	Validate() error
}

// JPush 的 Golang 推送客户端
// 详情： http://docs.jpush.io/server/rest_api_v3_push/
type PushClient struct {
	appKey       string
	masterSecret string
	headers      map[string]string
	http         *httplib.HTTPClient
}

func NewPushClient(appKey string, masterSecret string) *PushClient {
	client := PushClient{
		appKey:       appKey,
		masterSecret: masterSecret,
	}
	headers := make(map[string]string)
	headers["User-Agent"] = "jpush-api-go-client"
	headers["Connection"] = "keep-alive"
	headers["Authorization"] = "Basic " + basicAuth(appKey, masterSecret)
	client.headers = headers

	client.http = httplib.NewClient()

	return &client
}

// 推送 API
func (pc *PushClient) Push(payload *PushObject) (*PushResult, error) {
	return pc.doPush(PUSH_URL, payload)
}

// 推送校验 API
func (pc *PushClient) PushValidate(payload *PushObject) (*PushResult, error) {
	return pc.doPush(PUSH_VALIDATE_URL, payload)
}

func (pc *PushClient) doPush(url string, payload *PushClient) (*PushResult, error) {
	resp, err := pc.http.PostJson(url, payload, pc.headers)
	if err != nil {
		return nil, err
	}

	result := &PushResult{}
	err = result.FromResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}
