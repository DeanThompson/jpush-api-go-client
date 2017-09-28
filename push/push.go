package push

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jukylin/jpush-api-go-client/common"
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

	if po.Notification != nil {
		if err := po.Notification.Validate(); err != nil {
			return err
		}
	}

	if po.Message != nil {
		if err := po.Message.Validate(); err != nil {
			return err
		}
	}

	if po.Options != nil {
		if err := po.Options.Validate(); err != nil {
			return err
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
func (pr *PushResult) FromResponse(resp *http.Response) error {
	pr.ResponseBase = common.NewResponseBase(resp)
	if pr.ResponseBase.MsgId != nil {
		pr.MsgId = strconv.FormatFloat(pr.ResponseBase.MsgId.(float64), 'g', 64, 64)
	}
	if !pr.Ok() {
		return nil
	}
	return common.RespToJson(resp, &pr)
}

func (pr *PushResult) String() string {
	f := "<PushResult> msg_id: %v, sendno: \"%s\", \"%s\""
	return fmt.Sprintf(f, pr.MsgId, pr.SendNo, pr.ResponseBase.String())
}
