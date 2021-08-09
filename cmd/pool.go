package readmodbus

import (
	"fmt"

	"github.com/thinkgos/gomodbus"
	"github.com/xuebuqing/test/pkg"
)

func Pooling(ShuZu []*request.RequestPara, p *modbus.RTUClientProvider) {
	for {
		for i, x := range ShuZu {
			//testing
			fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
			Testing(*x, p)
		}
	}
}
