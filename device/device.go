package device

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jukylin/jpush-api-go-client/common"
)

type QueryDeviceResult struct {
	common.ResponseBase

	// 设备的所有属性，包含tags, alias
	Tags  []string `json:"tags"`
	Alias string   `json:"alias"`
}

func (dir *QueryDeviceResult) FromResponse(resp *http.Response) error {
	dir.ResponseBase = common.NewResponseBase(resp)
	if !dir.Ok() {
		return nil
	}
	return common.RespToJson(resp, &dir)
}

func (dir *QueryDeviceResult) String() string {
	return fmt.Sprintf("<QueryDeviceResult> tags: %v, alias: \"%s\", %v",
		dir.Tags, dir.Alias, dir.ResponseBase.String())
}

/////////////////////////////////////////////////////

type tags struct {
	Add    []string `json:"add,omitempty"`
	Remove []string `json:"remove,omitempty"`
	Clear  bool     `json:"-"`
}

type DeviceUpdate struct {
	// 支持 add, remove 或者空字符串。
	// 当 tags 参数为空字符串的时候，表示清空所有的 tags；
	// add/remove 下是增加或删除指定的 tags
	Tags tags

	// 更新设备的别名属性；当别名为空串时，删除指定设备的别名；
	Alias string

	// 手机号码
	Mobile string
}

type deviceUpdateWrapper struct {
	Tags   interface{} `json:"tags"`
	Alias  string      `json:"alias"`
	Mobile string      `json:"mobile"`
}

func NewDeviceUpdate() *DeviceUpdate {
	return &DeviceUpdate{
		Tags: tags{},
	}
}

func (du *DeviceUpdate) MarshalJSON() ([]byte, error) {
	wrapper := deviceUpdateWrapper{}
	if du.Tags.Clear {
		wrapper.Tags = ""
	} else {
		wrapper.Tags = du.Tags
	}
	wrapper.Alias = du.Alias
	wrapper.Mobile = du.Mobile
	return json.Marshal(wrapper)
}

func (du *DeviceUpdate) AddTags(tags ...string) {
	du.Tags.Clear = false
	tags = common.UniqString(tags)
	du.Tags.Add = common.UniqString(append(du.Tags.Add, tags...))
}

func (du *DeviceUpdate) RemoveTags(tags ...string) {
	du.Tags.Clear = false
	tags = common.UniqString(tags)
	du.Tags.Remove = common.UniqString(append(du.Tags.Remove, tags...))
}

func (du *DeviceUpdate) ClearAllTags() {
	du.Tags.Clear = true
}

func (du *DeviceUpdate) SetAlias(alias string) {
	du.Alias = alias
}

func (du *DeviceUpdate) SetMobile(mobile string) {
	du.Mobile = mobile
}
