package main

import (
	"encoding/json"
	"fmt"
)

type Servers struct {
	ServerName string `json:"name"`
	ServerIp string `json:"ip"`
	ServerPort int `json:"port"`
}

func main () {
	// todo 结构体序列化
	jsonStruct()
	// todo map 序列化
	jsonMap()

	// todo 反序列化结构体
	unJsonStruct()
	// todo 反序列化map
	unJsonMap()
}

// 反序列化结构体
func unJsonStruct()  {
	jsonString := `{"name": "demo serverName", "ip": "192.168.1.68", "port": 8080}`
	server := new(Servers)

	err := json.Unmarshal([]byte(jsonString), &server)

	if err != err {
		fmt.Println("Unmarshal Struct Err : ", server)
		return
	}
	fmt.Println("Unmarshal Struct Success : ", server)
}

// 反序列化map
func unJsonMap()  {
	jsonString := `{"name": "demo serverName", "ip": "192.168.1.68", "port": 8080}`
	server := make(map[string]interface{})

	err := json.Unmarshal([]byte(jsonString), &server)

	if err != err {
		fmt.Println("Unmarshal Map Err : ", server)
		return
	}
	fmt.Println("Unmarshal Map Success : ", server)
}

func jsonStruct()  {
	server := new(Servers)
	server.ServerName = "demo serverName"
	server.ServerIp = "192.168.1.9"
	server.ServerPort = 8080

	v,err := json.Marshal(server)

	if err != nil {
		fmt.Println("Marshal Struct Err : ", string(v))
		return
	}
	fmt.Println("Marshal Struct Success : ", string(v))
}

func jsonMap()  {
	server := make(map[string]interface{})
	server["ServerName"] = "demo serverName"
	server["ServerIp"] = "192.168.1.9"
	server["ServerPort"] = 8080

	v,err := json.Marshal(server)

	if err != nil {
		fmt.Println("Marshal Map Err : ", string(v))
		return
	}
	fmt.Println("Marshal Map Success : ", string(v))
}

