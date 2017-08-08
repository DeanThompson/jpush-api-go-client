package push

import (
	"encoding/json"

	"github.com/jukylin/jpush-api-go-client/common"
)

// “通知”对象，是一条推送的实体内容对象之一（另一个是“消息”）
type Notification struct {
	Alert    string                `json:"alert,omitempty"`
	Android  *AndroidNotification  `json:"android,omitempty"`
	Ios      *IosNotification      `json:"ios,omitempty"`
	Winphone *WinphoneNotification `json:"winphone,omitempty"`
}

func NewNotification(alert string) *Notification {
	return &Notification{Alert: alert}
}

func (n *Notification) Validate() error {
	if n.Ios != nil {
		return n.Ios.Validate()
	}
	return nil
}

// 平台通用的通知属性
type platformNotification struct {
	Alert  string                 `json:"alert"` // required
	Extras map[string]interface{} `json:"extras,omitempty"`
}

func (nc *platformNotification) AddExtra(key string, value interface{}) {
	if nc.Extras == nil {
		nc.Extras = make(map[string]interface{})
	}
	nc.Extras[key] = value
}

// Android 平台上的通知。
type AndroidNotification struct {
	platformNotification

	Title     string `json:"title,omitempty"`
	BuilderId int    `json:"builder_id,omitempty"`
	Style int    `json:"style,omitempty"`
	BigText string    `json:"big_text,omitempty"`
	Inbox string    `json:"inbox,omitempty"`
	BigPicPath string    `json:"big_pic_path,omitempty"`
}

func NewAndroidNotification(alert string) *AndroidNotification {
	n := &AndroidNotification{}
	n.Alert = alert
	return n
}

// iOS 平台上 APNs 通知。
type IosNotification struct {
	platformNotification

	Sound            string `json:"sound,omitempty"`
	Badge            int    `json:"badge,omitempty"`
	ContentAvailable bool   `json:"content-available,omitempty"`
	Category         string `json:"category,omitempty"`
}

func NewIosNotification(alert string) *IosNotification {
	a := &IosNotification{}
	a.Alert = alert
	return a
}

// APNs 协议定义通知长度为 2048 字节。
// JPush 因为需要重新组包，并且考虑一点安全冗余，
// 要求"iOS":{ } 及大括号内的总体长度不超过：2000 个字节。
func (in *IosNotification) Validate() error {
	data, err := json.Marshal(in)
	if err != nil {
		return err
	}

	if len("iOS:{}")+len(string(data)) >= IosNotificationMaxSize {
		return common.ErrIosNotificationTooLarge
	}

	return nil
}

// Windows Phone 平台上的通知。
type WinphoneNotification struct {
	platformNotification

	Title    string `json:"title,omitempty"`
	OpenPage string `json:"_open_page,omitempty"`
}

func NewWinphoneNotification(alert string) *WinphoneNotification {
	w := &WinphoneNotification{}
	w.Alert = alert
	return w
}
