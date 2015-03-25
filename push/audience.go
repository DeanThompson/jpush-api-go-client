package push

import "github.com/DeanThompson/jpush-api-go-client/common"

// 推送设备对象，表示一条推送可以被推送到哪些设备列表。
// 确认推送设备对象，JPush 提供了多种方式，比如：别名、标签、注册ID、分群、广播等。
type Audience struct {
	isAll bool // 是否推送给所有对象，如果是，value 无效
	value map[string][]string
}

func NewAudience() *Audience {
	return &Audience{}
}

func (a *Audience) Value() interface{} {
	if a.isAll {
		return ALL
	}
	return a.value
}

func (a *Audience) All() {
	a.isAll = true
}

func (a *Audience) SetTag(tags []string) {
	a.set("tag", tags)
}

func (a *Audience) SetTagAnd(tagAnds []string) {
	a.set("tag_and", tagAnds)
}

func (a *Audience) SetAlias(alias []string) {
	a.set("alias", alias)
}

func (a *Audience) SetRegistrationId(ids []string) {
	a.set("registration_id", ids)
}

func (a *Audience) set(key string, v []string) {
	a.isAll = false
	if a.value == nil {
		a.value = make(map[string][]string)
	}

	a.value[key] = common.UniqString(v)
}
