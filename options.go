package jpush

// 推送可选项。
type Options struct {
	SendNo          int   `json:"sendno,omitempty"`
	TimeToLive      int   `json:"time_to_live,omitempty"`
	OverrideMsgId   int64 `json:"override_msg_id,omitempty"`
	ApnsProduction  bool  `json:"apns_production,omitempty"`
	BigPushDuration int   `json:"big_push_duration,omitempty"`
}

func (self *Options) Validate() error {
	if self.TimeToLive > 0 {
		self.TimeToLive = maxInt(self.TimeToLive, MaxTimeToLive)
	}

	if self.BigPushDuration > 0 {
		self.BigPushDuration = maxInt(self.BigPushDuration, MaxBigPushDuration)
	}
	return nil
}
