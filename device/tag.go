package device

import (
	"fmt"
	"net/http"

	"github.com/DeanThompson/jpush-api-go-client/common"
)

// 查询标签列表请求结果
type GetTagsResult struct {
	common.ResponseBase

	Tags []string `json:"tags"`
}

func (tir *GetTagsResult) FromResponse(resp *http.Response) error {
	tir.ResponseBase = common.NewResponseBase(resp)
	if !tir.Ok() {
		return nil
	}
	return common.RespToJson(resp, &tir)
}

func (tir *GetTagsResult) String() string {
	return fmt.Sprintf("<GetTagsResult> tags: %v, %v", tir.Tags, tir.ResponseBase.String())
}

// 判断设备与标签的绑定请求结果
type CheckTagUserExistsResult struct {
	common.ResponseBase

	Result bool `json:"result"`
}

func (result *CheckTagUserExistsResult) FromResponse(resp *http.Response) error {
	result.ResponseBase = common.NewResponseBase(resp)
	if !result.Ok() {
		return nil
	}
	return common.RespToJson(resp, &result)
}

func (result *CheckTagUserExistsResult) String() string {
	return fmt.Sprintf("<CheckTagUserExistsResult> result: %v, %v", result.Result, result.ResponseBase)
}

// 更新标签（与设备的绑定的关系）请求参数
type UpdateTagUsersArgs struct {
	RegistrationIds map[string][]string `json:"registration_ids"`
}

func NewUpdateTagUsersArgs() *UpdateTagUsersArgs {
	return &UpdateTagUsersArgs{
		RegistrationIds: make(map[string][]string),
	}
}

func (args *UpdateTagUsersArgs) AddRegistrationIds(ids ...string) {
	args.updateRegistrationIds(actionAdd, ids...)
}

func (args *UpdateTagUsersArgs) RemoveRegistrationIds(ids ...string) {
	args.updateRegistrationIds(actionRemove, ids...)
}

func (args *UpdateTagUsersArgs) updateRegistrationIds(action string, ids ...string) {
	if action != actionAdd && action != actionRemove {
		return
	}
	ids = common.UniqString(ids)
	current := args.RegistrationIds[action]
	if current == nil {
		current = make([]string, 0, len(ids))
	}
	merged := common.UniqString(append(current, ids...))
	if len(merged) > maxAddOrRemoveRegistrationIds {
		merged = merged[:maxAddOrRemoveRegistrationIds]
	}
	args.RegistrationIds[action] = merged
}
