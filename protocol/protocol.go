package protocol

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gookit/event"
	"github.com/sagoo-cloud/iotgateway/consts"
	"github.com/sagoo-cloud/iotgateway/model"
)

type ChargeProtocol struct {
}

func (c *ChargeProtocol) Init(device *model.Device, data []byte) error {
	g.Log().Debug(context.Background(), "=========== 处理设备初始化，如设备的标识（deviceKey）获取 ===========", string(data))

	//设备初始化
	device.DeviceKey = "1234567890" //初始化设备KEY
	md := make(map[string]interface{})
	md["deviceType"] = "aaabbb" //设备元数据设置
	device.Metadata = md

	return nil
}

func (c *ChargeProtocol) Encode(device *model.Device, data interface{}, param ...string) (res []byte, err error) {
	return []byte(""), nil
}

// Decode 解码 如果是TCP模式的，则需要从buffer中解析出数据，然后返回数据。否则实现自己的方法
func (c *ChargeProtocol) Decode(device *model.Device, buffer []byte) (res []byte, err error) {

	//1，数据解析处理。。。。。

	//2，解析后，触发事件，向SagooIoT 发送数据。
	var propertieData = make(map[string]interface{})
	propertieData["XXX字段1"] = "XXX值1"
	propertieData["XXX字段2"] = "XXX值2"

	//推送数据
	out := g.Map{
		"DeviceKey":         "设备key",
		"PropertieDataList": propertieData,
	}
	event.MustFire(consts.PushAttributeDataToMQTT, out)

	return buffer, nil
}
