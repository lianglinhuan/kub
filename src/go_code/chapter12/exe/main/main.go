package main

import (
	"fmt"
	"os"
)

func main() {

	//os包下面的Args切片可以接收执行程序的传入变量
	//1.可以用go build 源文件，生成一个可执行文件，然后在通过执行可执行文件传入变量
	//2.也可以直接使用go run 源文件，后面直接添加传入变量

	fmt.Println("程序传入的参数数量为", len(os.Args))

	for i, v := range os.Args {
		fmt.Printf("args[%v] = %v \n", i, v)
	}
} 