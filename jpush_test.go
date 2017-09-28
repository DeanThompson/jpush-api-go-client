package jpush

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jukylin/jpush-api-go-client/device"
	"github.com/jukylin/jpush-api-go-client/push"
)

const (
	appKey       = "8b7127870ccae51a2c2e6da4"
	masterSecret = "55df2bc707d65fb39ca01325"
)

var client = NewJPushClient(appKey, masterSecret)

func init() {
	client.SetDebug(true)
}

func showResultOrError(method string, result interface{}, err error) {
	if err != nil {
		fmt.Printf("%s failed: %v\n\n", method, err)
	} else {
		fmt.Printf("%s result: %v\n\n", method, result)
	}
}

///////////////////// Push /////////////////////

func test_Push(t *testing.T) {
	// platform 对象
	platform := push.NewPlatform()
	// 用 Add() 方法添加具体平台参数，可选: "all", "ios", "android"
	platform.Add("ios", "android")
	// 或者用 All() 方法设置所有平台
	// platform.All()

	// audience 对象，表示消息受众
	audience := push.NewAudience()
	audience.SetTag([]string{"广州", "深圳"})   // 设置 tag
	audience.SetTagAnd([]string{"北京", "女"}) // 设置 tag_and
	// 和 platform 一样，可以调用 All() 方法设置所有受众
	// audience.All()

	// notification 对象，表示 通知，传递 alert 属性初始化
	notification := push.NewNotification("Notification Alert")

	// android 平台专有的 notification，用 alert 属性初始化
	androidNotification := push.NewAndroidNotification("Android Notification Alert")
	androidNotification.Title = "title"
	androidNotification.AddExtra("key", "value")

	notification.Android = androidNotification

	// iOS 平台专有的 notification，用 alert 属性初始化
	iosNotification := push.NewIosNotification("iOS Notification Alert")
	iosNotification.Badge = 1
	// Validate 方法可以验证 iOS notification 是否合法
	// 一般情况下，开发者不需要直接调用此方法，这个方法会在构造 PushObject 时自动调用
	// iosNotification.Validate()

	notification.Ios = iosNotification

	// Windows Phone 平台专有的 notification，用 alert 属性初始化
	wpNotification := push.NewWinphoneNotification("Winphone Notification Alert")
	// 所有平台的专有 notification 都有 AddExtra 方法，用于添加 extra 信息
	wpNotification.AddExtra("key", "value")
	wpNotification.AddExtra("extra_key", "extra_value")

	notification.Winphone = wpNotification

	// message 对象，表示 透传消息，用 content 属性初始化
	message := push.NewMessage("Message Content must not be empty")
	message.Title = "Message Title"

	// option 对象，表示推送可选项
	options := push.NewOptions()
	// iOS 平台，是否推送生产环境，false 表示开发环境；如果不指定，就是生产环境
	options.ApnsProduction = true
	// Options 的 Validate 方法会对 time_to_live 属性做范围限制，以满足 JPush 的规范
	options.TimeToLive = 10000000
	// Options 的 Validate 方法会对 big_push_duration 属性做范围限制，以满足 JPush 的规范
	options.BigPushDuration = 1500

	payload := push.NewPushObject()
	payload.Platform = platform
	payload.Audience = audience
	payload.Notification = notification
	payload.Message = message
	payload.Options = options

	data, err := json.Marshal(payload)
	if err != nil {
		t.Error("json.Marshal PushObject failed:", err)
	}
	fmt.Println("payload:", string(data), "\n")

	// Push 会推送到客户端
	// result, err := client.Push(payload)
	//	showErrOrResult("Push", result, err)

	// PushValidate 的参数和 Push 完全一致
	// 区别在于，PushValidate 只会验证推送调用成功，不会向用户发送任何消息
	result, err := client.PushValidate(payload)
	showResultOrError("PushValidate", result, err)
}

///////////////////// Device /////////////////////

func test_QueryDevice(t *testing.T) {
	registrationId := "123456"
	info, err := client.QueryDevice(registrationId)
	showResultOrError("QueryDevice", info, err)
}

func test_UpdateDevice(t *testing.T) {
	update := device.NewDeviceUpdate()
	update.AddTags("tag1", "tag2")
	update.SetMobile("13800138000")
	registrationId := "123456"
	result, err := client.UpdateDevice(registrationId, update)
	showResultOrError("UpdateDevice", result, err)
}

///////////////////// Tags /////////////////////

func test_GetTags(t *testing.T) {
	result, err := client.GetTags()
	showResultOrError("GetTags", result, err)
}

func test_CheckTagUserExists(t *testing.T) {
	tag := "tag1"
	registrationId := "090c1f59f89"
	result, err := client.CheckTagUserExists(tag, registrationId)
	showResultOrError("CheckTagUserExists", result, err)
}

func test_UpdateTagUsers(t *testing.T) {
	args := device.NewUpdateTagUsersArgs()
	args.AddRegistrationIds("123", "234", "345")
	args.RemoveRegistrationIds("abc", "bcd")
	fmt.Println("UpdateTagUsersArgs", args.RegistrationIds)
	result, err := client.UpdateTagUsers("tag1", args)
	showResultOrError("UpdateTagUsers", result, err)
}

func test_DeleteTag(t *testing.T) {
	result, err := client.DeleteTag("tag1", nil)
	showResultOrError("DeleteTag", result, err)

	result, err = client.DeleteTag("tag2", []string{"ios", "android"})
	showResultOrError("DeleteTag", result, err)
}

///////////////////// Alias /////////////////////

func test_GetAliasUsers(t *testing.T) {
	result, err := client.GetAliasUsers("alias1", nil)
	showResultOrError("GetAliasUsers", result, err)

	result, err = client.GetAliasUsers("alias1", []string{"ios", "android"})
	showResultOrError("GetAliasUsers", result, err)
}

func test_DeleteAlias(t *testing.T) {
	result, err := client.DeleteAlias("alias1", nil)
	showResultOrError("DeleteAlias", result, err)

	result, err = client.DeleteAlias("alias1", []string{"ios", "android"})
	showResultOrError("DeleteAlias", result, err)
}

///////////////////// Report /////////////////////

func test_GetReceivedReport(t *testing.T) {
	msgIds := []uint64{1613113584, 1229760629}
	result, err := client.GetReceivedReport(msgIds)
	showResultOrError("GetReceivedReport", result, err)
}

func Test_Starter(t *testing.T) {
	test_Push(t)

	test_QueryDevice(t)
	test_UpdateDevice(t)

	test_GetTags(t)
	test_CheckTagUserExists(t)
	test_UpdateTagUsers(t)
	test_DeleteTag(t)

	test_GetAliasUsers(t)
	test_DeleteAlias(t)

	test_GetReceivedReport(t)
}
