package request

import (
	"encoding/json"
	"fmt"
	"os"
)

//配置请求类型定义
type BaseConfig struct {
	DeviceNode string
	BaudRate   int
	DataBits   int
	Parity     string
	StopBits   int
}

//请求格式类型定义
type RequestPara struct {
	MasterId    byte
	Slave       byte
	Fcode       uint16
	Polling     int
	InitAddress uint16
	Length      uint16
	Data        []byte
}
type Request struct {
	BaseConfig
	RequestPara
}

//定义请求数组的格式
type ShuZu []*RequestPara

//给请求数组添加删除方法
func (Shu ShuZu) Del(s *RequestPara) []*RequestPara {
	var shuzu []*RequestPara
	for i := 0; i < len(Shu); i++ {
		fmt.Println("s: ", s)
		fmt.Println("vs[i]: ", Shu[i])
		if s != Shu[i] {
			shuzu = append(shuzu, Shu[i])
		}
	}
	return shuzu
}

//给请求数组添加增加方法
func (Shu ShuZu) Add(s *RequestPara) []*RequestPara {
	Shu = append(Shu, s)
	return Shu
}

func Reciving() Request {
	// 打开request文件
	var Requesting Request
	file1, _ := os.Open("conf/request.json")
	defer file1.Close()
	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder1 := json.NewDecoder(file1)
	RequestConf := RequestPara{}
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	err1 := decoder1.Decode(&RequestConf)
	if err1 != nil {
		fmt.Println("Error2:", err1)
	}

	// 打开config文件
	file2, _ := os.Open("conf/config.json")
	defer file2.Close()
	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder2 := json.NewDecoder(file2)
	ConfigConf := BaseConfig{}
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	err2 := decoder2.Decode(&ConfigConf)
	if err2 != nil {
		fmt.Println("Error2:", err2)
	}
	Requesting.BaseConfig = ConfigConf
	Requesting.RequestPara = RequestConf

	return Requesting
}
