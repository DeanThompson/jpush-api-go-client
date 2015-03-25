package push

import (
	"github.com/DeanThompson/jpush-api-go-client/common"
	"github.com/DeanThompson/jpush-api-go-client/httplib"
	"github.com/DeanThompson/jpush-api-go-client/push"
)

// JPush 的 Golang 推送客户端
// 详情： http://docs.jpush.io/server/rest_api_v3_push/
type JPushClient struct {
	appKey       string
	masterSecret string
	headers      map[string]string
	http         *httplib.HTTPClient
}

func NewPushClient(appKey string, masterSecret string) *JPushClient {
	client := JPushClient{
		appKey:       appKey,
		masterSecret: masterSecret,
	}
	headers := make(map[string]string)
	headers["User-Agent"] = "jpush-api-go-client"
	headers["Connection"] = "keep-alive"
	headers["Authorization"] = "Basic " + common.BasicAuth(appKey, masterSecret)
	client.headers = headers

	client.http = httplib.NewClient()

	return &client
}

// 推送 API
func (jpc *JPushClient) Push(payload *push.PushObject) (*push.PushResult, error) {
	return jpc.doPush(common.PUSH_URL, payload)
}

// 推送校验 API
func (jpc *JPushClient) PushValidate(payload *push.PushObject) (*push.PushResult, error) {
	return jpc.doPush(common.PUSH_VALIDATE_URL, payload)
}

func (jpc *JPushClient) doPush(url string, payload *JPushClient) (*push.PushResult, error) {
	resp, err := jpc.http.PostJson(url, payload, jpc.headers)
	if err != nil {
		return nil, err
	}

	result := &push.PushResult{}
	err = result.FromResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}
