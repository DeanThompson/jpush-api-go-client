package report

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DeanThompson/jpush-api-go-client/common"
)

type ReceivedReportNode struct {
	MsgId           uint64 `json:"msg_id"`
	AndroidReceived int    `json:"android_received"` // Android 送达。如果无此项数据则为 null
	IosApnsSent     int    `json:"ios_apns_sent"`    // iOS 推送成功。如果无此项数据则为 null
	IosMsgReceived  int    `json:"ios_msg_received"` // iOS 自定义消息送达数。如果无此项数据则为null
	WpMpnsSent      int    `json:"wp_mpns_sent"`     // winphone 通知送达。如果无此项数据则为 null
}

func (node *ReceivedReportNode) String() string {
	v, _ := json.Marshal(node)
	return string(v)
}

type ReceiveReport struct {
	common.ResponseBase

	Report []ReceivedReportNode
}

func (report *ReceiveReport) FromResponse(resp *http.Response) error {
	report.ResponseBase = common.NewResponseBase(resp)
	if !report.Ok() {
		return nil
	}
	return common.RespToJson(resp, &report.Report)
}

func (report *ReceiveReport) String() string {
	s, _ := json.Marshal(report.Report)
	return fmt.Sprintf("<ReceiveReport> report: %v, %v", string(s),
		report.ResponseBase.String())
}
