package main

import "fmt"

type USB interface {
	Name() string
	Connecter // 嵌入接口

}

type Connecter interface {
	Connect()
}

type Phone struct {
	name string
}

func (pc Phone) Name() string {
	return pc.name
}

func (pc Phone) Connect() {
	fmt.Println("连接成功", pc.name)
}

func main() {
	//var a USB
	//a := Phone{"phone"}
	// 强制转换usb 转换为connect
	// TODO 接口转换
	pc := Phone{"pc"} // 初始化

	var a Connecter
	a = Connecter(pc)
	a.Connect()
	//Disconnect(a)
}

//func Disconnect (usb USB) {
//	// 简单的断言 判断
//	if pc,ok := usb.(Phone);ok {
//		fmt.Println("succ", pc.name)
//		return
//	}
//	fmt.Println("Error")
//}

func Disconnect (usb interface {}) { // 可以传递任何类型
	switch v:= usb.(type) { // v 局部变量
	case Phone:
		fmt.Println("succ", v.name)
	default:
		fmt.Println("Error")
	}
}