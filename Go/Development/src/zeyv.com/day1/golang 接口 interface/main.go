package main

import "fmt"

// todo 接口是公共方法的集合
type MethodsInterfaces interface {
	Run() string
	Eat() bool
	Jump() string
}

// 定义Dog的结构体
type Dog struct {
	Name string
	Age int
}

type Cat struct {
	Name string
	Age int
}
func (d *Dog) Run ()string {
	fmt.Println("Dog Run Run Runs")
	return " Dog Run Run Runs"
}


func (d *Dog) Eat ()bool {
	return false
}

func (d *Dog) Jump ()string {
	return "Jumping"
}

func (d *Cat) Run ()string {
	fmt.Println("Cat Run Run Runs")
	return " Cat Run Run Runs"
}


func (d *Cat) Eat ()bool {
	return false
}

func (d *Cat) Jump ()string {
	return "Jumping"
}


func main () {
	//dog := Dog{Name:"JoMs", Age:88}
	//dog.Run()
	//fmt.Println(dog, "struct")
	// 接口定义变量
	var dog MethodsInterfaces
	dog = new(Dog)
	dog.Run()
}
