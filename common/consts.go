package common

import "errors"

const (
	PUSH_URL          = "https://api.jpush.cn/v3/push"
	PUSH_VALIDATE_URL = "https://api.jpush.cn/v3/push/validate"
)

var (
	ErrInvalidPlatform         = errors.New("<Platform>: invalid platform")
	ErrMessageContentMissing   = errors.New("<Message>: msg_content is required.")
	ErrContentMissing          = errors.New("<PushObject>: notification or message is required")
	ErrIosNotificationTooLarge = errors.New("<IosNotification>: iOS notification too large")
)
