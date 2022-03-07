package main

import (
	"fmt"
)

func sum(n1, n2 int) int{

	//当执行defer是，暂时不执行，会将defer后面的语句压入到独立的栈(defer栈)
	//当函数执行完毕后，再凶defer栈按照先入后出方式出栈，执行
	defer fmt.Printf("ok1 n1=%v\n", n1)
	defer fmt.Printf("ok2 n2=%v\n", n2)

	n1++
	n2++
	res := n1 + n2
	fmt.Printf("ok3 res=%v\n", res)
	return res
}

func main() {

	res := sum(10, 20)
	fmt.Printf("res=%v\n", res)

}
