package main
// todo err
import (
	"database/sql"
	"log"
	"net/http"
)

// Message 是调用了api约定返回的格式结构体

type Message struct {
	Code string
	Msg string
	Time int64
	UserInfo UserInfo
}

// userinfo查询了数据库之后 用于接受数据的结构体
type UserInfo struct {
	Id int
	Name string
	Age int
}

// 比main执行前触发
func init () {
	// 初始化mysql驱动 获得db
	var db,err = sql.Open("mysql", "root:password@tcp(localhost:3306)/shop?utf8mb4")
	// 设置最大连接数
	db.SetMaxOpenConns(20)
	// 设置空闲数
	db.SetMaxIdleConns(10)
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatalf("Error on opening database connection: %s",pingErr.Error())
	}
	checkErr(err)
}

func queryUserInfo(name string) ([]byte, error)  {
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main () {
	http.HandleFunc("/userinfo", handler)
	http.ListenAndServe("localhost:8899", nil)
}
