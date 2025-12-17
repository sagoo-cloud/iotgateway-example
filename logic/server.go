package logic

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gookit/event"
	"github.com/sagoo-cloud/iotgateway"
	"github.com/sagoo-cloud/iotgateway/consts"
)

var Gateway *iotgateway.Gateway // 定义全局网关对象

func Init() {

	// 从这个云网关向平台推送数据，有两种方式：
	// 1. 推送单台设备属性数据（直接普通设备不经过网关处理）
	// 2. 推送多台设备属性数据（批量数据，经过平台的网关处理模块处理）
	// 这里演示这两种方式的定时推送数据。跟据实际业务需求选择使用。

	// 定时向SagooIoT 推送批量设备数据
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {

			//模拟向SagooIoT 发送多个子设备数据。节约网络带宽。
			subDevicesDataList := []map[string]interface{}{
				{
					"ProductKey": "exampleDeviceProductKey",
					"DeviceKey":  "exampleDeviceKey",
					"PropertieDataList": map[string]interface{}{
						"va": grand.N(100, 200),
						"vb": grand.N(100, 200),
						"vc": grand.N(100, 200),
					},
				},
				{
					"ProductKey": "exampleDeviceProductKey2",
					"DeviceKey":  "exampleDeviceKey2",
					"PropertieDataList": map[string]interface{}{
						"va": grand.N(100, 200),
						"vb": grand.N(100, 200),
						"vc": grand.N(100, 200),
					},
				},
				{
					"ProductKey":     "exampleDeviceProductKey3",
					"DeviceKey":      "exampleDeviceKey3",
					consts.EventName: "PowerErrorEvent", // 事件名称，这个值是需要事先在SagooIoT 平台定义好的事件名称
					consts.EventDataList: map[string]interface{}{
						"vababa": grand.N(800, 1000),
						"vacdcd": "error info",
					},
				},
				{
					"ProductKey":     "exampleDeviceProductKey4",
					"DeviceKey":      "exampleDeviceKey4",
					consts.EventName: "PowerError2Event", // 事件名称，这个值是需要事先在SagooIoT 平台定义好的事件名称
					consts.EventDataList: map[string]interface{}{
						"vababa": grand.N(1000, 1500),
						"vacdcd": grand.N(1500, 2000),
					},
				},
			}

			//推送数据
			out := g.Map{
				"ProductKey":         "exampleDeviceProductKey",
				"DeviceKey":          "exampleDeviceKey",
				"SubDevicesDataList": subDevicesDataList,
			}

			event.MustFire(consts.PushAttributePackDataToMQTT, out)
		}
	}()

	// 定时向SagooIoT 推送设备直接透传数据，不经过添加的网关节处理
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {

			//模拟向SagooIoT 发送数据。========== 测试推送数据 ==========
			var propertieData = make(map[string]interface{})
			propertieData["va"] = grand.N(100, 200)
			propertieData["vb"] = grand.N(100, 200)
			propertieData["vc"] = grand.N(100, 200)

			//推送数据
			out := g.Map{
				"ProductKey":             "exampleDeviceProductKey",
				"DeviceKey":              "exampleDeviceKey",
				consts.PropertieDataList: propertieData,
			}

			event.MustFire(consts.PushAttributeDataToMQTT, out)
		}
	}()

	// 定时向SagooIoT 推送设备事件数据
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {

			//模拟向SagooIoT 发送数据。========== 测试推送数据 ==========
			var propertieData = make(map[string]interface{})
			propertieData["vababa"] = grand.N(800, 1000)

			//推送数据
			out := g.Map{
				"ProductKey":         "exampleDeviceProductKey",
				"DeviceKey":          "exampleDeviceKey",
				consts.EventName:     "PowerErrorEvent", // 事件名称，这个值是需要事先在SagooIoT 平台定义好的事件名称
				consts.EventDataList: propertieData,
			}

			event.MustFire(consts.PushAttributeDataToMQTT, out)
		}
	}()

}
