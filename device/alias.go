package device

import (
	"fmt"
	"net/http"

	"github.com/Aqiling/jpush-api-go-client/common"
)

type GetAliasUsersResult struct {
	common.ResponseBase

	RegistrationIds []string `json:"registration_ids"`
}

func (result *GetAliasUsersResult) FromResponse(resp *http.Response) error {
	result.ResponseBase = common.NewResponseBase(resp)
	if !result.Ok() {
		return nil
	}
	return common.RespToJson(resp, &result)
}

func (result *GetAliasUsersResult) String() string {
	return fmt.Sprintf("<GetAliasUsersResult> RegistrationIds: %v, %v",
		result.RegistrationIds, result.ResponseBase.String())
}
