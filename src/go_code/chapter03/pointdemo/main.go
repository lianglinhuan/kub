package main
import (
	"fmt"
)

//演示golang中的指针类型
func main() {
	
	//基本数据类型在内存布局
	var i int = 10
	//i的地址是多少来着？
	fmt.Printf("i的地址=%v\n", &i)

	//ptr是一个int（整形）指针变量，值为&i（i的地址）
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	//ptr本身也有一个地址
	fmt.Printf("ptr的地址=%v\n", &ptr)
	//取出ptr指正变量所指向地址的值
	fmt.Printf("ptr指向的值=%v\n", *ptr)
}