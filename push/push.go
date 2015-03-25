package push

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/DeanThompson/jpush-api-go-client/common"
)

type Validator interface {
	Validate() error
}

// 一个推送对象，表示一条推送相关的所有信息。
type PushObject struct {
	platform     *Platform     `json:"platform"`
	audience     *Audience     `json:"audience"`
	notification *Notification `json:"notification"`
	message      *Message      `json:"message"`
	options      *Options      `json:"options"`
}

func (po *PushObject) Validate() error {
	if po.notification == nil && po.message == nil {
		return common.ErrContentMissing
	}

	for _, v := range []Validator{po.notification, po.message, po.options} {
		if v != nil {
			if err := v.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}

type pushObjectWrapper struct {
	Platform     interface{}   `json:"platform"`
	Audience     interface{}   `json:"audience"`
	Notification *Notification `json:"notification,omitempty"`
	Message      *Message      `json:"message,omitempty"`
	Options      *Options      `json:"options,omitempty"`
}

// 实现 Marshaler interface
func (po *PushObject) MarshalJSON() ([]byte, error) {
	if err := po.Validate(); err != nil {
		return nil, err
	}

	wrapper := pushObjectWrapper{
		Platform:     po.platform.Value(),
		Audience:     po.audience.Value(),
		Notification: po.notification,
		Message:      po.message,
		Options:      po.options,
	}
	return json.Marshal(wrapper)
}

type PushError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PushResult struct {
	// HTTP 状态码
	StatusCode int

	// 频率限制相关
	RateLimitQuota     int
	RateLimitRemaining int
	RateLimitReset     int

	// 成功时返回的 body
	MsgId  string `json:"msg_id"`
	SendNo string `json:"sendno"`

	// 失败时返回的 body
	Error PushError `json:"error"`
}

// 成功： {"sendno":"18","msg_id":"1828256757"}
// 失败：
//    {
//        "error": {
//            "code": 2002,
//            "message": "Rate limit exceeded"
//        }
//    }
//
// 所有的 HTTP API Response Header 里都加了三项频率控制信息：
//
// X-Rate-Limit-Limit：    当前 AppKey 一个时间窗口内可调用次数
// X-Rate-Limit-Remaining：当前时间窗口剩余的可用次数
// X-Rate-Limit-Reset：    距离时间窗口重置剩余的秒数
func (pr *PushResult) FromResponse(resp *http.Response) error {
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	// 成功或失败时解析出返回的数据
	// 实际上只有当 StatusCode = 200 时，才有 msg_id 和 sendno
	// 其他情况下只有 error 数据。 error 和 (msg_id, sendno) 不会同时存在
	err = json.Unmarshal(data, &pr)
	if err != nil {
		return err
	}

	pr.StatusCode = resp.StatusCode
	pr.RateLimitQuota, _ = common.GetIntHeader(resp, rateLimitQuotaHeader)
	pr.RateLimitRemaining, _ = common.GetIntHeader(resp, rateLimitRemainingHeader)
	pr.RateLimitReset, _ = common.GetIntHeader(resp, rateLimitResetHeader)

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
