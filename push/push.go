package push

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DeanThompson/jpush-api-go-client/common"
)

type Validator interface {
	Validate() error
}

// 一个推送对象，表示一条推送相关的所有信息。
type PushObject struct {
	Platform     *Platform     `json:"platform"`
	Audience     *Audience     `json:"audience"`
	Notification *Notification `json:"notification"`
	Message      *Message      `json:"message"`
	Options      *Options      `json:"options"`
}

func NewPushObject() *PushObject {
	return &PushObject{}
}

func (po *PushObject) Validate() error {
	if po.Notification == nil && po.Message == nil {
		return common.ErrContentMissing
	}

	for _, v := range []Validator{po.Notification, po.Message, po.Options} {
		if v != nil {
			if err := v.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}

// 实现 Marshaler interface
func (po *PushObject) MarshalJSON() ([]byte, error) {
	if err := po.Validate(); err != nil {
		return nil, err
	}

	wrapper := pushObjectWrapper{
		Platform:     po.Platform.Value(),
		Audience:     po.Audience.Value(),
		Notification: po.Notification,
		Message:      po.Message,
		Options:      po.Options,
	}
	return json.Marshal(wrapper)
}

type pushObjectWrapper struct {
	Platform     interface{}   `json:"platform"`
	Audience     interface{}   `json:"audience"`
	Notification *Notification `json:"notification,omitempty"`
	Message      *Message      `json:"message,omitempty"`
	Options      *Options      `json:"options,omitempty"`
}

type PushResult struct {
	common.ResponseBase

	// 成功时 msg_id 是 string 类型。。。
	// 失败时 msg_id 是 int 类型。。。
	MsgId  interface{} `json:"msg_id"`
	SendNo string      `json:"sendno"`
}

// 成功： {"sendno":"18", "msg_id":"1828256757"}
// 失败： {"msg_id": 1035959738, "error": {"message": "app_key does not exist", "code": 1008}}
//
// 所有的 HTTP API Response Header 里都加了三项频率控制信息：
//
// X-Rate-Limit-Limit：    当前 AppKey 一个时间窗口内可调用次数
// X-Rate-Limit-Remaining：当前时间窗口剩余的可用次数
// X-Rate-Limit-Reset：    距离时间窗口重置剩余的秒数
func (pr *PushResult) FromResponse(resp *http.Response) error {
	// 成功或失败时解析出返回的数据
	// 实际上只有当 StatusCode = 200 时，才有 msg_id 和 sendno
	// 其他情况下只有 error 数据。 error 和 (msg_id, sendno) 不会同时存在
	err := common.RespToJson(resp, &pr)
	if err != nil {
		return err
	}

	pr.ResponseBase = common.NewResponseBase(resp)

	return nil
}

// 根据请求返回的 HTTP 状态码判断推送是否成功
// 规范：
// - 200 一定是正确。所有异常都不使用 200 返回码
// - 业务逻辑上的错误，有特别的错误码尽量使用 4xx，否则使用 400。
// - 服务器端内部错误，无特别错误码使用 500。
// - 业务异常时，返回内容使用 JSON 格式定义 error 信息。
//
// 更多细节： http://docs.jpush.io/server/http_status_code/
func (pr *PushResult) Ok() bool {
	return pr.StatusCode == 200
}

func (pr *PushResult) ErrorCode() int {
	if pr.Error != nil {
		return pr.Error.Code
	}
	return 0
}

func (pr *PushResult) ErrorMsg() string {
	if pr.Error != nil {
		return pr.Error.Message
	}
	return ""
}

func (pr *PushResult) String() string {
	f := "<PushResult> msg_id: %v, sendno: %s, %s"
	return fmt.Sprintf(f, pr.MsgId, pr.SendNo, pr.ResponseBase.String())
}
