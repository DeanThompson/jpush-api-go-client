package main

import (
	"fmt"

	"github.com/DeanThompson/jpush-api-go-client"
	"github.com/DeanThompson/jpush-api-go-client/device"
)

func device_example() {
	const (
		appKey       = "8b7127870ccae51a2c2e6da4"
		masterSecret = "55df2bc707d65fb39ca01325"
	)

	// 创建 JPush 的客户端
	jclient := jpush.NewJPushClient(appKey, masterSecret)
	jclient.SetDebug(true)

	registrationId := "123456"

	infoResult, err := jclient.QueryDevice(registrationId)
	if err != nil {
		fmt.Println("\nQueryDevice failed:", err)
	} else {
		fmt.Println("\nQueryDevice result:", infoResult)
	}

	update := device.NewDeviceUpdate()
	update.AddTags("tag1", "tag2")
	updateResult, err := jclient.UpdateDevice(registrationId, update)
	if err != nil {
		fmt.Println("\nUpdateDevice failed:", err)
	} else {
		fmt.Println("\nUpdateDevice result:", updateResult)
	}
}

//func main() {
//    device_example()
//}
