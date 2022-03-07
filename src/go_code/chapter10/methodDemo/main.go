package main

import (
	"fmt"
)

//定义一个结构体类型
type Person struct {
	name string
}

//绑定类型的方法，类型为结构体
//test方法和Person类型绑定，test方法只能通过Person类型来调用
//不能通过其他类型来调用，或者直接调用
func (p Person) test() {
	fmt.Println("test() name=", p.name)
}

func main() {

	var p Person
	p.name = "tom"
	p.test() //调用方法
}