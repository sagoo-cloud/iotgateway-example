package main

import (
	"errors"
	"sagoo-example-gateway/events"
	"sagoo-example-gateway/logic"
	"sagoo-example-gateway/protocol"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gookit/event"
	"github.com/sagoo-cloud/iotgateway"
	"github.com/sagoo-cloud/iotgateway/consts"
	"github.com/sagoo-cloud/iotgateway/version"
)

// 定义编译时的版本信息
var (
	BuildVersion = "0.0"
	BuildTime    = ""
	CommitID     = ""
)

func main() {
	glog.SetDefaultLogger(g.Log())
	//显示版本信息
	version.ShowLogo(BuildVersion, BuildTime, CommitID)
	ctx := gctx.GetInitCtx()

	//创建网关,如果是TCP或UDP设备，则需要传入解析协议，如果为mqtt设备，protocol为nil，设备类型在配置文件中进行配置：netType的值
	//gateway, err := iotgateway.NewGateway(ctx, nil)

	proto := &protocol.ChargeProtocol{}
	err := errors.New("")
	logic.Gateway, err = iotgateway.NewGateway(ctx, proto)
	if err != nil {
		panic(err)
	}

	events.Init() //初始化事件

	//模拟向SagooIoT 发送数据。========== 测试推送数据 ==========
	var propertieData = make(map[string]interface{})
	propertieData["XXX字段1"] = "XXX值1"
	propertieData["XXX字段2"] = "XXX值2"

	//推送数据
	out := g.Map{
		"DeviceKey":         "设备key",
		"PropertieDataList": propertieData,
	}
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			event.MustFire(consts.PushAttributeDataToMQTT, out)
		}
	}()

	// 结束测试推送数据 ==========

	logic.Gateway.Start() //启动网关
}
