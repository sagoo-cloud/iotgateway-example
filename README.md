# Sagoo Iotgateway-example

## 编译与运行

### 本地源码运行

修改 config 下的 config.yaml 文件，并配置相关项。请根据注释进行配置修改，包括服务相关配置，日志相关配置。

注意：在生产环境中要将日志级别设置为 error，避免日志过大。

运行命令：
```sh
go run main.go
```


## 网关与设备相关参数


### 网关相关参数

在/config/config.yaml 中配置的

产品KEY：agoo-example-geteway-productkey

网关设备KEY：sagoo-example-geteway-devicekey


### 内置的模拟子设备相关参数：

在代码中当前是硬编码的，实际过程中可以动态获取。

子设备产品KEY：exampleDeviceProductKey

子设备KEY：exampleDeviceKey

### 子设备属性：
| 属性 | 类型 | 描述   |
|----| --- |------|
| va | int | A相电压 |
| vb | int | B相电压 |
| vc | int | C相电压 |

### 子设备事件：
| 事件          | 参数 | 描述  |
|-------------| --- |-----|
| DeviceEvent | Power | 值on |
|             | WF | 值2  |

这个事件是一个模拟事件，实际过程中可以根据需要定义不同的事件和参数。



