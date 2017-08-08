package jpush

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jukylin/jpush-api-go-client/common"
	"github.com/jukylin/jpush-api-go-client/device"
	"github.com/jukylin/jpush-api-go-client/httplib"
	"github.com/jukylin/jpush-api-go-client/push"
	"github.com/jukylin/jpush-api-go-client/report"
)

// JPush 的 Golang 推送客户端
// 详情： http://docs.jpush.io/server/rest_api_v3_push/
type JPushClient struct {
	appKey       string
	masterSecret string
	headers      map[string]string
	http         *httplib.HTTPClient
}

func NewJPushClient(appKey string, masterSecret string) *JPushClient {
	client := JPushClient{
		appKey:       appKey,
		masterSecret: masterSecret,
	}
	headers := make(map[string]string)
	headers["User-Agent"] = "jpush-api-go-client"
	headers["Connection"] = "keep-alive"
	headers["Authorization"] = common.BasicAuth(appKey, masterSecret)
	client.headers = headers

	client.http = httplib.NewClient()

	return &client
}

// 设置调试模式，调试模式下，会输出日志
func (jpc *JPushClient) SetDebug(debug bool) {
	jpc.http.SetDebug(debug)
}

// 推送 API
func (jpc *JPushClient) Push(payload *push.PushObject) (*push.PushResult, error) {
	return jpc.doPush(common.PUSH_URL, payload)
}

// 推送校验 API， 只用于验证推送调用是否能够成功，与推送 API 的区别在于：不向用户发送任何消息。
func (jpc *JPushClient) PushValidate(payload *push.PushObject) (*push.PushResult, error) {
	return jpc.doPush(common.PUSH_VALIDATE_URL, payload)
}

func (jpc *JPushClient) doPush(url string, payload *push.PushObject) (*push.PushResult, error) {
	resp, err := jpc.http.PostJson(url, payload, jpc.headers)
	if err != nil {
		return nil, err
	}

	result := &push.PushResult{}
	err = result.FromResponse(resp)
	return result, err
}

// 查询设备(设备的别名与标签)
func (jpc *JPushClient) QueryDevice(registrationId string) (*device.QueryDeviceResult, error) {
	url := fmt.Sprintf(common.DEVICE_URL, registrationId)
	resp, err := jpc.http.Get(url, nil, jpc.headers)
	if err != nil {
		return nil, err
	}

	result := &device.QueryDeviceResult{}
	err = result.FromResponse(resp)
	return result, err
}

// 更新设备 （设置的别名与标签）
func (jpc *JPushClient) UpdateDevice(registrationId string, payload *device.DeviceUpdate) (*common.ResponseBase, error) {
	url := fmt.Sprintf(common.DEVICE_URL, registrationId)
	resp, err := jpc.http.PostJson(url, payload, jpc.headers)
	return common.ResponseOrError(resp, err)
}

// 查询标签列表
func (jpc *JPushClient) GetTags() (*device.GetTagsResult, error) {
	resp, err := jpc.http.Get(common.QUERY_TAGS_URL, nil, jpc.headers)
	if err != nil {
		return nil, err
	}

	result := &device.GetTagsResult{}
	err = result.FromResponse(resp)
	return result, err
}

// 判断设备与标签的绑定
func (jpc *JPushClient) CheckTagUserExists(tag string, registrationId string) (*device.CheckTagUserExistsResult, error) {
	url := fmt.Sprintf(common.CHECK_TAG_USER_EXISTS_URL, tag, registrationId)
	resp, err := jpc.http.Get(url, nil, jpc.headers)
	if err != nil {
		return nil, err
	}

	result := &device.CheckTagUserExistsResult{}
	err = result.FromResponse(resp)
	return result, err
}

// 更新标签 （与设备的绑定的关系）
func (jpc *JPushClient) UpdateTagUsers(tag string, payload *device.UpdateTagUsersArgs) (*common.ResponseBase, error) {
	url := fmt.Sprintf(common.UPDATE_TAG_USERS_URL, tag)
	resp, err := jpc.http.PostJson(url, payload, jpc.headers)
	return common.ResponseOrError(resp, err)
}

// 删除标签 (与设备的绑定关系)
func (jpc *JPushClient) DeleteTag(tag string, platforms []string) (*common.ResponseBase, error) {
	url := fmt.Sprintf(common.DELETE_TAG_URL, tag)
	params := addPlatformsToParams(platforms)
	resp, err := jpc.http.Delete(url, params, jpc.headers)
	return common.ResponseOrError(resp, err)
}

// 查询别名 （与设备的绑定关系）
func (jpc *JPushClient) GetAliasUsers(alias string, platforms []string) (*device.GetAliasUsersResult, error) {
	url := fmt.Sprintf(common.QUERY_ALIAS_URL, alias)
	params := addPlatformsToParams(platforms)
	resp, err := jpc.http.Get(url, params, jpc.headers)
	if err != nil {
		return nil, err
	}

	result := &device.GetAliasUsersResult{}
	err = result.FromResponse(resp)
	return result, err
}

// 删除别名 （与设备的绑定关系）
func (jpc *JPushClient) DeleteAlias(alias string, platforms []string) (*common.ResponseBase, error) {
	url := fmt.Sprintf(common.DELETE_ALIAS_URL, alias)
	params := addPlatformsToParams(platforms)
	resp, err := jpc.http.Delete(url, params, jpc.headers)
	return common.ResponseOrError(resp, err)
}

// 送达统计
func (jpc *JPushClient) GetReceivedReport(msgIds []uint64) (*report.ReceiveReport, error) {
	ids := make([]string, 0, len(msgIds))
	for _, msgId := range msgIds {
		ids = append(ids, strconv.FormatUint(msgId, 10))
	}
	params := map[string]interface{}{"msg_ids": strings.Join(ids, ",")}

	resp, err := jpc.http.Get(common.RECEIVED_REPORT_URL, params, jpc.headers)
	if err != nil {
		return nil, err
	}

	result := &report.ReceiveReport{}
	err = result.FromResponse(resp)
	return result, err
}

////////////////////////////////////////////////////////////////////////////////

func addPlatformsToParams(platforms []string) map[string]interface{} {
	if platforms == nil {
		return nil
	}
	params := make(map[string]interface{})
	params["platform"] = strings.Join(platforms, ",")
	return params
}
