package device

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DeanThompson/jpush-api-go-client/common"
)

// 设备的所有属性，包含tags, alias
type DeviceInfo struct {
	Tags  []string `json:"tags"`
	Alias string   `json:"alias"`
}

type DeviceInfoResult struct {
	common.ResponseBase
	DeviceInfo
}

func (dir *DeviceInfoResult) FromResponse(resp *http.Response) error {
	err := common.RespToJson(resp, &dir)
	if err != nil {
		return err
	}
	dir.ResponseBase = common.NewResponseBase(resp)
	return nil
}

func (dir *DeviceInfoResult) String() string {
	return fmt.Sprintf("<DeviceInfoResult> tags: %v, alias: %s, %v",
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
}

type deviceUpdateWrapper struct {
	Tags  interface{} `json:"tags"`
	Alias string      `json:"alias"`
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
