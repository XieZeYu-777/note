package importStruct

import "fmt"

type Master struct {
	Name string
	Age int
}

func TestStruct()  {
	nStruct := Master{Name:"bali", Age:10}
	fmt.Println(nStruct)
}