package jpush

import "errors"

const (
	ALL = "all"

	PLATFORM_IOS     = "ios"
	PLATFORM_ANDROID = "android"
	PLATFORM_WP      = "winphone"

	PUSH_URL          = "https://api.jpush.cn/v3/push"
	PUSH_VALIDATE_URL = "https://api.jpush.cn/v3/push/validate"

	IosNotificationMaxSize = 2000

	rateLimitQuotaHeader     = "X-Rate-Limit-Limit"
	rateLimitRemainingHeader = "X-Rate-Limit-Remaining"
	rateLimitResetHeader     = "X-Rate-Limit-Reset"

	MaxTimeToLive      = 10 * 24 * 60 * 60 // 10 天
	MaxBigPushDuration = 1400
)

var (
	ErrInvalidPlatform         = errors.New("<Platform>: invalid platform")
	ErrMessageContentMissing   = errors.New("<Message>: msg_content is required.")
	ErrContentMissing          = errors.New("<PushObject>: notification or message is required")
	ErrIosNotificationTooLarge = errors.New("<IosNotification>iOS notification too large")
)
