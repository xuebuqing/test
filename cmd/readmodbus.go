package readmodbus

import (
	"fmt"
	"time"

	modbus "github.com/thinkgos/gomodbus"
	"github.com/xuebuqing/test/pkg"
)

func Testing(ConfigConf request.RequestPara, p *modbus.RTUClientProvider) {

	client := modbus.NewClient(p)
	client.LogMode(true)
	err := client.Connect()
	if err != nil {
		fmt.Println("start err,", err)
		return
	}
	numBytes := ConfigConf.Data     //寄存器
	cnumBytes := []byte{0xff, 0xff} //线圈
	fcode := ConfigConf.Fcode
	slave := ConfigConf.Slave
	length := ConfigConf.Length
	address := ConfigConf.InitAddress
	fmt.Println("this ismaster num :", slave, fcode)

	switch {
	case fcode == 1:
		//读线圈   	   功能码 0x01

		value, err := client.ReadCoils(slave, address, length)
		if err != nil {
			fmt.Println("readHoldErr,", err)
		} else {
			fmt.Printf("功能码 0x01\n")
			fmt.Printf("%#v\n", value)
		}
		time.Sleep(time.Second * 3)
	case fcode == 2:
		//读离散线圈   	   功能码 0x02
		value2, err2 := client.ReadDiscreteInputs(17, 0xC4, 4)
		if err2 != nil {
			fmt.Println("readHoldErr,", err2)
		} else {
			fmt.Printf("功能码 0x02\n")
			fmt.Printf("%#v\n", value2)
		}
		time.Sleep(time.Second * 3)

	case fcode == 3:
		//读保持器  	   功能码 0x03
		value3, err3 := client.ReadHoldingRegisters(17, 0x6B, 0x03)
		if err3 != nil {
			fmt.Println("readHoldErr,", err3)
		} else {
			fmt.Printf("功能码 0x03\n")
			fmt.Printf("%#v\n", value3)
		}
		time.Sleep(time.Second * 3)
	case fcode == 4:
		value4, err4 := client.ReadHoldingRegistersBytes(17, 0x6B, 0x03)
		if err4 != nil {
			fmt.Println("readHoldErr,", err4)
		} else {
			fmt.Printf("功能码 0x03\n")
			fmt.Printf("%#v\n", value4)
		}
	case fcode == 5:
		//读保持器字节  	   功能码 0x03
		value5, err5 := client.ReadInputRegistersBytes(17, 0x08, 1)
		if err5 != nil {
			fmt.Println("readHoldErr,", err5)
		} else {
			fmt.Printf("功能码 0x04\n")
			fmt.Printf("%#v\n", value5)
		}
		time.Sleep(time.Second * 3)
	case fcode == 6:
		err6 := client.WriteSingleCoil(17, 2, true)
		if err6 != nil {
			fmt.Println("readHoldErr,", err6)
		} else {
			fmt.Printf("功能码 0x05\n")
		}
		time.Sleep(time.Second * 3)
	case fcode == 7:
		err7 := client.WriteSingleRegister(17, 2, 13)
		if err7 != nil {
			fmt.Println("readHoldErr,", err7)
		} else {
			fmt.Printf("功能码 0x06\n")
		}
		time.Sleep(time.Second * 3)
	case fcode == 8:
		err8 := client.WriteMultipleRegisters(17, 3, 3, numBytes)
		if err8 != nil {
			fmt.Println("readHoldErr,", err8)
		} else {
			fmt.Printf("功能码 0x10\n")
		}
		time.Sleep(time.Second * 3)
	case fcode == 9:
		err9 := client.WriteMultipleCoils(17, 3, 16, cnumBytes)
		if err9 != nil {
			fmt.Println("readHoldErr,", err9)
		} else {
			fmt.Printf("功能码 0x0F\n")
		}
		time.Sleep(time.Second * 3)
	case fcode == 10:
		value10, err10 := client.ReadWriteMultipleRegisters(17, 1, 4, 1, 3, numBytes)
		if err10 != nil {
			fmt.Println("readHoldErr,", err10)
		} else {
			fmt.Printf("功能码 0x17\n")
			fmt.Printf("%#v\n", value10)
		}
		time.Sleep(time.Second * 3)
	case fcode == 11:
		value11, err11 := client.ReadWriteMultipleRegistersBytes(17, 1, 3, 1, 1, cnumBytes)
		if err11 != nil {
			fmt.Println("readHoldErr,", err11)
		} else {
			fmt.Printf("功能码 0x17\n")
			fmt.Printf("%#v\n", value11)
		}
		time.Sleep(time.Second * 3)
	default:
		fmt.Printf("功能码错误\n")
	}

}
