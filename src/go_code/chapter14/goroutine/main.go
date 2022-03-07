package main

import (
	"fmt"
	"time"
	"strconv"
)

func test() {
	for i := 0; i <= 5; i++ {
		fmt.Println("hello, world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	
	//在函数前面直接使用go，即可实现协程
	//协程
	go test()

	//主协程
	for i := 0; i <= 10; i++{
		fmt.Printf("hello, goland%v\n", i)
		time.Sleep(time.Second)
	}
	//主协程和协程打印开始不是固定的，但是都是跟着顺序间隔打印出来
}