package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan<- i
	}

	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan<- "hello" + fmt.Sprintf("%d", i)
	}

	//要是使用传统的方法遍历管道是，如果不关闭会阻塞而导致deadlock
	
	//关闭管道不失为一种好方法，可是在某些开发中，可能不确定什么时候该关闭管道
	//可以使用select方式来解决
	//label:
	for {
		select {
			//注意：这里，如果intChan一致没有关闭，不会一致阻塞deadlock
			//会自动到下一个case匹配，直到default退出
		case v := <- intChan:
			fmt.Printf("从intChan读取的数据%d\n", v)
			time.Sleep(time.Second)
		case v := <-stringChan :
			fmt.Printf("从stringChan读取到的数据%s\n", v)
			time.Sleep(time.Second)
		default :
			fmt.Printf("都读取不到了\n")
			time.Sleep(time.Second)
			//return直接退出main函数，表示程序终止
			return
			//通过设定label标签，通过break退出标签
			//注意：直接使用break是不能够退出程序的
			//break label
		}
	}
}