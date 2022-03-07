package main

import (
	"fmt"
)

func main() {

	//演示管道的使用
	//1.设定一个管道可以存放3个int数据
	var intChan chan int
	intChan = make(chan int, 3)

	//2.看输出的结果是什么
	fmt.Printf("intChan 的值=%v intChan 本身的地址=%v\n", intChan, &intChan)
	//返回的结果是一个地址，所以证明channel(管道)是引用类型

	//3.向管道写入数据
	intChan<- 2
	num := 211
	intChan<- num

	//注意：当我们给管道写入数据时，不能超过其容量，会报错

	//4.查看管道的长度len和容量cap
	fmt.Printf("channel len(intChan) = %v  cap(intChan) = %v\n", len(intChan), cap(intChan)) //2, 3

	//5.从管道中读取数据
	var num01 int
	num01 = <-intChan
	fmt.Printf("num01 = %v\n", num01) //2, 先进先出
	fmt.Printf("channel len(intChan) = %v  cap(intChan) = %v\n", len(intChan), cap(intChan)) //1, 3

	//注意：在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告错误deadlock
}