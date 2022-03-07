package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age int
}

func main() {
	//定义一个存放任意数据类型的管道3个数据
	//var allChan chan interface{}
	allChan := make(chan interface{}, 3)

	allChan<- 10
	allChan<- "tom jack"
	cat := Cat{"小花猫", 4}
	allChan<- cat

	//我们希望得到的是第三个元素，则先将前两个推出
	<-allChan
	<-allChan

	newCat := <-allChan

	fmt.Printf("newCat=%T , newCat=%v\n", newCat, newCat)

	//下面的写法是错误的！编译不能够通过
	//因为channel是interface{}，空接口类型，而不是原先的Cat结构体类型
	//fmt.Printf("newCat.Name=%v\n", newCat.Name)

	//使用类型断言，将数据类型(空接口类型)转换回Cat结构体类型
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v\n", a.Name)
}