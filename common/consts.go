package common

import "errors"

const (
	rateLimitQuotaHeader     = "X-Rate-Limit-Limit"
	rateLimitRemainingHeader = "X-Rate-Limit-Remaining"
	rateLimitResetHeader     = "X-Rate-Limit-Reset"

	PUSH_URL          = "https://api.jpush.cn/v3/push"
	PUSH_VALIDATE_URL = "https://api.jpush.cn/v3/push/validate"

	DEVICE_URL = "https://device.jpush.cn/v3/devices/%s"
)

var (
	ErrInvalidPlatform         = errors.New("<Platform>: invalid platform")
	ErrMessageContentMissing   = errors.New("<Message>: msg_content is required.")
	ErrContentMissing          = errors.New("<PushObject>: notification or message is required")
	ErrIosNotificationTooLarge = errors.New("<IosNotification>: iOS notification too large")
)
