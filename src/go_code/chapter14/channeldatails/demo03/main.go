package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello, world")
	}
}

func test() {
	//这里使用一defer + recover机制，采集错误panic
	//避免了因为协程错误，导致主线程也退出
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("test() 发生错误 = ",err)
		}
	}()

	//定义了一个map
	var myMap map[int]string
	//因为map使用之前是需要make的，所以这里会报错，产生一个panic
	//协程发生错误而导致主线程也退出
	myMap[0] = "golang"
}

func main() {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hi, tom")
	}

}