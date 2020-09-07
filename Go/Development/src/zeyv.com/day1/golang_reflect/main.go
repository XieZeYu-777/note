package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age int
}

func (u User) Hello() {
	fmt.Println("Hello Wrold")
}

func main () {
	u := User{Name:"YO", Age:10}
	// 获取目标对象
	t := reflect.TypeOf(u)
	fmt.Println("TypeOf", t)
	// 获取目标对象的值类型
	v := reflect.ValueOf(u)
	fmt.Println("ValueOf", v)

	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i))
		// 从0开始获取User所包含的key
		key := t.Field(i)
		// 通过interface方法获取key所对应的值
		value := v.Field(i).Interface()

		fmt.Printf("第%d个字段是: %s：%v = %v \n",i+1,key.Name,key.Type, value)
	}
	// 获取User里头的方法
	for i := 0; i < t.NumMethod(); i++ {
		// 从0开始获取User所包含的key
		m := t.Method(i)
		fmt.Println(m, "mmm")
		// 通过interface方法获取key所对应的值

		fmt.Printf("第%d个方法是: %s：%v \n",i+1,m.Name,m.Type)
	}

	// todo log data
	// TypeOf main.User
	// ValueOf {YO 10}
	// {Name  string  0 [0] false}
	//  第1个字段是: Name：string = YO
	// {Age  int  16 [1] false}
	//  第2个字段是: Age：int = 10
	// {Hello  func(main.User) <func(main.User) Value> 0} mmm
	//  第1个方法是: Hello：func(main.User)

	//var name string = "人类,是为了恋爱和革命而生的!~"
	//// 返回name的类型 如string float等
	//reflectType := reflect.TypeOf(name)
	//// 返回上文的文字
	//reflectValue := reflect.ValueOf(name)
	//
	//fmt.Println("reflectType :", reflectType)
	//fmt.Println("reflectValue :", reflectValue)
}
