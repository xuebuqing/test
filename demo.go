package main

import (
	"fmt"
	"sync"

	//"time"
	"github.com/xuebuqing/test/cmd"
	"github.com/xuebuqing/test/pkg"
)

var wg sync.WaitGroup

func main() {

	/*创建集合 */
	var RequestPooling request.ShuZu
	var requestMap map[byte]request.Request
	requestMap = make(map[byte]request.Request)
	/* reading map */

	RequestStruct := request.Reciving()
	requestMap[RequestStruct.RequestPara.MasterId] = RequestStruct
	//configing

	for key := range requestMap {
		fmt.Println("key是", key)
		Pt := request.Configing(requestMap[key].BaseConfig)
		conf := requestMap[key].RequestPara
		RequestPooling = RequestPooling.Add(&conf)
		go readmodbus.Pooling(RequestPooling, Pt)

	}
	wg.Add(1)

	//RequestPooling = RequestPooling.Del(&RequestConf)
	//time.Sleep(time.Second * 3)
	fmt.Println(RequestPooling)

	wg.Wait()
}
