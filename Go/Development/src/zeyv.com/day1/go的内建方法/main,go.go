package main

import (
	"errors"
	"fmt"
	"reflect"
)
// todo make new 是创建变量的
// todo append delete copy 是用来操作数据类型（变量
func goMake () {
	// todo make 创建 slice 切片（类似数组 map 键值对 channel（chan // 管道 两个线程之间数据通信 返回的是引用类型
	// make 创建切片
	// todo 返回的是引入类型 new返回的是指针类型
	mSlice := make([]string,3) // make(切片类型, 切片的长度
	mSlice[0] = "cat"
	mSlice[1] = "dog"
	mSlice[2] = "tiger"

	fmt.Println(mSlice, "切片")

	// 创建map[keyType]valueType
	mMap := make(map[int]string)
	mMap[10] = "men"
	mMap[100] = "wmen"
	// {10: "men", 100: "wmen"}
	fmt.Println(mMap, "map")

	// map 创建channel 管道
	mChannel := make(chan int, 3) // make(chan 类型, 缓存容量
	// mChannel := make(chan int) // make(chan 类型, 如果不填的话就是无缓存的channel
	fmt.Println(mChannel, "Channel")
}

func goNew () {
	// todo new 内存置零 返回传入类型的指针地址
	// 创建map
	nMap := new(map[int]string)
	makeMap := make(map[int]string)

	fmt.Println(reflect.TypeOf(nMap)) // *map[int]string 指针类型
	fmt.Println(reflect.TypeOf(makeMap)) // map[int]string 引用类型
}
func sliceAppend () {
	sSlice := make([]string, 2)
	sSlice[0] = "lol"
	sSlice[1] = "cf"

	sSlice = append(sSlice, "asphyxia")
	fmt.Println(sSlice, "APPEND")
}
// copy
func sliceCopy () {
	// 切片1
	cSlice1 := make([]int, 3)
	cSlice1[0] = 33
	cSlice1[1] = 44
	cSlice1[2] = 55
	// 切片2
	cSlice2 := make([]int, 3)
	cSlice2[0] = 99
	cSlice2[1] = 98
	cSlice2[2] = 97
	copy(cSlice1,cSlice2) // copy(copy到那个切片, copy切片源
	fmt.Println(cSlice1) // [99 98 97]
}
func mapDelete () {
	mDelete := make(map[int]string)
	mDelete[0] = "id-1"
	mDelete[1] = "id-2"
	mDelete[2] = "id-3"
	// 删除map
	delete(mDelete, 0)
	fmt.Println(mDelete) // todo map[1:id-2 2:id-3]
}
func errPanic() {
	defer errRecover()
	panic(errors.New("err")) // 程序无法运行 err 类型的异常触发
	//panic(1) // 程序无法运行 默认 类型的异常触发
	//panic("err") // 程序无法运行 string 类型的异常触发
}
func errRecover()  {
	message := recover()
	switch message.(type) {
	case string:
		fmt.Println("string Err ", message)
	case error:
		fmt.Println("error Err", message)
	default:
		fmt.Println("default Err", message)
	}
}

func getLenCap () {
	tLen := make([]string, 2, 8)
	tLen[0] = "stringOne"
	tLen[1] = "stringTwo"
	tLen = append(tLen, "wwd")

	fmt.Println("长度->", len(tLen), "容量->", cap(tLen))
}

func getClose () {
	tChannel := make(chan int, 2)
	tChannel <- 1 // 写
	tChannel <- 12 // 写
	fmt.Println(tChannel)
	//todo defer 当前方法完成后执行的
	defer close(tChannel) // 关闭channel
}
// 内建方法
func main()  {
	// todo golang -> make new
	// goMake() // make
	//goNew() // new
	// todo slice -> append & copy
	// todo map -> delete
	//sliceAppend()
	//sliceCopy()
	//mapDelete()
	// todo 处理异常 panic recover
	// todo panic 抛出异常(程序无法继续运行 recover 捕获异常(错误信息提示 捕获
	//errPanic()
	//errRecover()
	// todo len 支持 string array slice map chan
	// todo cap 支持 array slice chan
	// todo close 支持 chan
	//getLenCap()
	getClose()

}
