package main
import (
	"fmt"
)

func Type(types ...interface{}) {
	for index, v := range types{
		switch v.(type) {  //这里的type是一个关键字，固定写法
			case bool:
				fmt.Printf("第%v个数的类型是%T，值是%v\n", index+1, v, v) 
			case float32:
				fmt.Printf("第%v个数的类型是%T，值是%v\n", index+1, v, v)
			case float64:
				fmt.Printf("第%v个数的类型是%T，值是%v\n", index+1, v, v)
			case int, int32, int64:
				fmt.Printf("第%v个数的类型是%T，值是%v\n", index+1, v, v)
			case string:
				fmt.Printf("第%v个数的类型是%T，值是%v\n", index+1, v, v)
			default:
				fmt.Printf("第%v个数的类型不确定，值是%v\n", index+1, v, v)
		}
	}
}

func main() {
	var n1 int64 = 10;
	var n2 float64 = 1.11;
	var n3 bool
	var n4 string = "tom";
	name := "mary"
	math := 100

	Type(n1, n2, n3, n4, name, math)
}
