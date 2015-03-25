package push

import "github.com/DeanThompson/jpush-api-go-client/common"

// 推送可选项。
type Options struct {
	SendNo          int   `json:"sendno,omitempty"`
	TimeToLive      int   `json:"time_to_live,omitempty"`
	OverrideMsgId   int64 `json:"override_msg_id,omitempty"`
	ApnsProduction  bool  `json:"apns_production,omitempty"`
	BigPushDuration int   `json:"big_push_duration,omitempty"`
}

func NewOptions() *Options {
	return &Options{}
}

// 可选项的校验没有那么严格，目前而言并没有出错的情况
// 只是针对某些值做一下范围限制
func (self *Options) Validate() error {
	if self.TimeToLive > 0 {
		self.TimeToLive = common.MinInt(self.TimeToLive, MaxTimeToLive)
	}

	if self.BigPushDuration > 0 {
		self.BigPushDuration = common.MinInt(self.BigPushDuration, MaxBigPushDuration)
	}
	return nil
}
