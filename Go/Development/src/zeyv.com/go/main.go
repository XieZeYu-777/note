package main

import (
	//'os'
	"fmt"
	"sort"

	//'strings'
)


func main()  {
	//fmt.Println(os.Args[0])
	//fmt.Println(strings.Join(os.Args[1:], '  ')) // 拼接
	//for _,v := range os.Args {
	//	fmt.Println(v)
	//}

	//a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	//sa := a[1:5]
	//sb := sa[0:2] // reslice
	////sl := make([]int, 3,10)
	//fmt.Println(string(sb))
	//
	//
	//s1 := make([]int , 3, 6)
	//
	//fmt.Printf("%p", s1)
	//s1 = append(s1, 33,2)
	//fmt.Printf("%v %p\n", s1,s1)

	s1 := []int {1,2,3,4}
	s2 :=[]int {7,8,9}

	copy(s1[1:2],s2[1:2])
	fmt.Println(s1)



	s3 := []int {1,2,0,4}

	sort.Ints(s3)

	fmt.Println(s3)

}


