package request

import (
	"fmt"
	"time"

	"github.com/thinkgos/gomodbus"
)

type Poine *modbus.RTUClientProvider

func Configing(pio BaseConfig) Poine {
	//调用RTUClientProvider的构造函数,返回结构体指针
	p := modbus.NewRTUClientProvider()
	p.Address = pio.DeviceNode
	p.BaudRate = pio.BaudRate
	p.DataBits = pio.DataBits
	p.Parity = pio.Parity
	p.StopBits = pio.StopBits
	p.Timeout = 100 * time.Millisecond
	fmt.Println("this progame is still \n")
	return p
}
