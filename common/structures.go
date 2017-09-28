package common

import (
	"fmt"
	"net/http"
)

type ErrorResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (pe *ErrorResult) String() string {
	return fmt.Sprintf("{code: %d, message: \"%s\"}", pe.Code, pe.Message)
}

type RateLimitInfo struct {
	RateLimitQuota     int
	RateLimitRemaining int
	RateLimitReset     int
}

//所有的 HTTP API Response Header 里都加了三项频率控制信息：
//
// X-Rate-Limit-Limit：    当前 AppKey 一个时间窗口内可调用次数
// X-Rate-Limit-Remaining：当前时间窗口剩余的可用次数
// X-Rate-Limit-Reset：    距离时间窗口重置剩余的秒数
func NewRateLimitInfo(resp *http.Response) *RateLimitInfo {
	info := &RateLimitInfo{}
	info.RateLimitQuota, _ = GetIntHeader(resp, rateLimitQuotaHeader)
	info.RateLimitRemaining, _ = GetIntHeader(resp, rateLimitRemainingHeader)
	info.RateLimitReset, _ = GetIntHeader(resp, rateLimitResetHeader)
	return info
}

func (info *RateLimitInfo) String() string {
	return fmt.Sprintf("{limit: %d, remaining: %d, reset: %d}",
		info.RateLimitQuota, info.RateLimitRemaining, info.RateLimitReset)
}

type ResponseBase struct {
	// HTTP 状态码
	StatusCode int

	// 频率限制相关
	RateLimitInfo *RateLimitInfo

	MsgId interface{} `json:"msg_id"`

	// 错误相关
	Error *ErrorResult `json:"error"`
}

func NewResponseBase(resp *http.Response) ResponseBase {
	rb := ResponseBase{
		StatusCode:    resp.StatusCode,
		RateLimitInfo: NewRateLimitInfo(resp),
	}
	if !rb.Ok() {
		RespToJson(resp, &rb)
	}
	return rb
}

// 根据请求返回的 HTTP 状态码判断推送是否成功
// 规范：
// - 200 一定是正确。所有异常都不使用 200 返回码
// - 业务逻辑上的错误，有特别的错误码尽量使用 4xx，否则使用 400。
// - 服务器端内部错误，无特别错误码使用 500。
// - 业务异常时，返回内容使用 JSON 格式定义 error 信息。
//
// 更多细节： http://docs.jpush.io/server/http_status_code/
func (rb *ResponseBase) Ok() bool {
	return rb.StatusCode == http.StatusOK
}

func (rb *ResponseBase) String() string {
	return fmt.Sprintf("StatusCode: %d, rateLimit: %v, error: %v",
		rb.StatusCode, rb.RateLimitInfo, rb.Error)
}
