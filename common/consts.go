package common

import "errors"

const (
	rateLimitQuotaHeader     = "X-Rate-Limit-Limit"
	rateLimitRemainingHeader = "X-Rate-Limit-Remaining"
	rateLimitResetHeader     = "X-Rate-Limit-Reset"

	push_host   = "https://api.jpush.cn"
	device_host = "https://device.jpush.cn"

	PUSH_URL          = push_host + "/v3/push"
	PUSH_VALIDATE_URL = push_host + "/v3/push/validate"

	// GET /v3/devices/{registration_id}
	DEVICE_URL = device_host + "/v3/devices/%s"

	QUERY_TAGS_URL = device_host + "/v3/tags/"
	// GET /v3/tags/{tag_value}/registration_ids/{registration_id}
	CHECK_TAG_USER_EXISTS_URL = device_host + "/v3/tags/%s/registration_ids/%s"
	// POST /v3/tags/{tag_value}
	UPDATE_TAG_USERS_URL = device_host + "/v3/tags/%s"
	// DELETE /v3/tags/{tag_value}
	DELETE_TAG_URL = device_host + "/v3/tags/%s"

	// GET /v3/aliases/{alias_value}
	QUERY_ALIAS_URL = device_host + "/v3/aliases/%s"
	// DELETE /v3/aliases/{alias_value}
	DELETE_ALIAS_URL = device_host + "/v3/aliases/%s"
)

var (
	ErrInvalidPlatform         = errors.New("<Platform>: invalid platform")
	ErrMessageContentMissing   = errors.New("<Message>: msg_content is required.")
	ErrContentMissing          = errors.New("<PushObject>: notification or message is required")
	ErrIosNotificationTooLarge = errors.New("<IosNotification>: iOS notification too large")
)
