package main

// 导入包名
import "fmt"

// 变量和常量

// 单独声明变量
// var str string // 声明一个变量必须要声明数据类型

// 批量声明
var (
	name string
	age  int
	isOk bool
)

// todo 声明变量必须使用
func main() {
	// 函数里赋值
	name = "zeyv"
	age = 22
	isOk = true
	fmt.Printf("name:%s", name) // %s 相当于占位符使用这个name变量的值替换占位符
	// fmt 是引入的库
	fmt.Println("hello world 你he好世界") // 打印完后面加一个换行符
}
