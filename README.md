JPush API Go Client
====================

[![GoDoc](https://godoc.org/github.com/DeanThompson/jpush-api-go-client?status.svg)](https://godoc.org/github.com/DeanThompson/jpush-api-go-client) [![Build Status](https://travis-ci.org/DeanThompson/jpush-api-go-client.svg?branch=master)](https://travis-ci.org/DeanThompson/jpush-api-go-client)

# 概述

这是 JPush REST API 的 Golang 版本封装开发包，**非官方实现**，只支持 v3 版本。

官方 REST API 文档： [http://docs.jpush.cn/display/dev/REST+API](http://docs.jpush.cn/display/dev/REST+API)

# 安装

使用 go get 安装，无任何第三方依赖：

```sh
go get github.com/DeanThompson/jpush-api-go-client
```

# 使用方法

## 1. 创建 JPushClient

```go
import "github.com/DeanThompson/jpush-api-go-client"

const (
    appKey = ""
    masterSecret = ""
)

jclient := jpush.NewPushClient(appKey, masterSecret)
```

## 2. 逐步构建消息体

与推送有关的数据结构都在 push 包里

```go
import "github.com/DeanThompson/jpush-api-go-client/push"
```

### 2.1 创建 Platform 对象

```go
platform := push.NewPlatform()

// 用 Add() 方法添加具体平台参数，可选: "all", "ios", "android"
platform.Add("ios", "android")

// 或者用 All() 方法设置所有平台
// platform.All()
```

### 2.2 创建 Audience 对象

```go
audience := push.NewAudience()
audience.SetTag([]string{"广州", "深圳"})   // 设置 tag
// audience.SetTagAnd([]string{"北京", "女"}) // 设置 tag_and
// audience.SetAlias([]string{"alias1", "alias2"})  // 设置 alias
// audience.SetRegistrationId([]string{"id1", "id2"})   // 设置 registration_id
// audience.All()   和 platform 一样，可以调用 All() 方法设置所有受众
```

### 2.3 创建 Notification 对象

#### 2.3.1 创建 AndroidNotification 对象

```go
// android 平台专有的 notification，用 alert 属性初始化
androidNotification := push.NewAndroidNotification("Android Notification Alert")
androidNotification.Title = "title"
// androidNotification.BuilderId = 10086
androidNotification.AddExtra("key", "value")
```

#### 2.3.2 创建 IosNotification 对象

```go
// iOS 平台专有的 notification，用 alert 属性初始化
iosNotification := push.NewIosNotification("iOS Notification Alert")
// iosNotification.Sound = "/paht/to/sound"
iosNotification.Badge = 1   // 只支持 int 类型的 badge
// iosNotification.ContentAvailable = true
// iosNotification.Category = "category_name"
// iosNotification.AddExtra("key", "value")

// Validate 方法可以验证 iOS notification 是否合法
// 一般情况下，开发者不需要直接调用此方法，这个方法会在构造 PushObject 时自动调用
// iosNotification.Validate()
```

#### 2.3.3 创建 WinphoneNotification 对象

```go
// Windows Phone 平台专有的 notification，用 alert 属性初始化
wpNotification := push.NewWinphoneNotification("Winphone Notification Alert")
// wpNotification.Title = "Winphone Notification Title"
// wpNotification.OpenPage = "some page"

// 所有平台的专有 notification 都有 AddExtra 方法，用于添加 extra 信息
wpNotification.AddExtra("extra_key", "extra_value")
```

#### 2.3.4 创建 Notification 对象

AndroidNotification, IosNotification, WinphoneNotification 三个分别是三种平台专有的 notification。

Notification 是极光推送的“通知”，包含一个 alert 属性，和可选的三个平台属性。

```go
// notification 对象，表示 通知，传递 alert 属性初始化
notification := push.NewNotification("Notification Alert")
notification.Android = androidNotification
notification.Ios = iosNotification
notification.Winphone = wpNotification
```

### 2.4 创建 Message 对象

Message 是极光推送的“消息”，也叫透传消息

```go
// message 对象，表示 透传消息，用 content 属性初始化
message := push.NewMessage("Message Content must not be empty")
// message.Title = "Message Title"
// message.ContentType = "application/json"

// 可以调用 AddExtra 方法，添加额外信息
// message.AddExtra("key", 123)
```

### 2.5 创建 Options 对象

```go
// option 对象，表示推送可选项
options := push.NewOptions()
// options.SendNo = 12345
// options.OverrideMsgId = 123333333

// Options 的 Validate 方法会对 time_to_live 属性做范围限制，以满足 JPush 的规范
options.TimeToLive = 10000000

// iOS 平台，是否推送生产环境，false 表示开发环境；如果不指定，就是生产环境
options.ApnsProduction = true

// Options 的 Validate 方法会对 big_push_duration 属性做范围限制，以满足 JPush 的规范
options.BigPushDuration = 1500

// Options 对象有 Validate 方法，但实际上这里并不会返回错误，
// 而是对 time_to_live 和 big_push_duration 两个值做了范围限制
// 开发者不需要手动调用此方法，在构建 PushObject 时会自动调用
// err := options.Validate()
```

### 2.6 创建 PushObject 对象

```go
payload := push.NewPushObject()
payload.Platform = platform
payload.Audience = audience
payload.Notification = notification
payload.Message = message
payload.Options = options

// 打印查看 json 序列化的结果，也就是 POST 请求的 body
// data, err := json.Marshal(payload)
// if err != nil {
//    fmt.Println("json.Marshal PushObject failed:", err)
// } else {
//    fmt.Println("payload:", string(data), "\n")
// }
```

## 3. 推送/推送验证

```go
// result, err := jclient.Push(payload)
result, err := jclient.PushValidate(payload)
if err != nil {
    fmt.Print("PushValidate failed:", err)
} else {
    fmt.Println("PushValidate result:", result)
}
```

## 4. 完整示例

完整的示例可以直接运行，代码在这里：[examples/push.go](examples/push.go)
