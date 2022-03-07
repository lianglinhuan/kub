package main
import (
	"fmt"
)

type Persion struct{
	ptr *int
	slice []int
	map1 map[string]string
}

func main() {

	var p1 Persion
	fmt.Println(p1)

	//无论是结构体还是正常定义，记得区分好引用的使用方式
	//注意：切片在使用前一定需要用make
	p1.slice = make([]int, 10)
	p1.slice[0] = 100

	//注意：map使用前也一定需要使用make
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "tom"

	fmt.Println(p1)

}