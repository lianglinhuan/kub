package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int, 3)
	intChan<- 100
	intChan<- 200

	//close()函数是内置函数，内置函数都存放在builtin包里面
	close(intChan)
	//intChan<- 300 //会报错
	fmt.Printf("okokok\n")

	//当管道关闭后，读取数据是可以的
	n1 := <-intChan
	fmt.Printf("n1=%v\n", n1)
	
	intChan1 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan1<- i*2
	}

	//在遍历前关闭管道
	close(intChan1)
	//for—range遍历管道，只返回一个值，因为管道是队列的数据结构
	//它是没有下标的，所以值返回一个值
	for v := range intChan1 {
		fmt.Printf("v = %v\n", v)
	}
}