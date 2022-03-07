package main
import (
	"fmt"
)

func main() {

	//有两个变量a，b，现在要将其两者进行交换，但是不允许使用中间变量
	var a int = 10
	var b int = 20
	fmt.Printf("交换前a=%v, b=%v\n", a, b)

	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("交换后a=%v, b=%v\n", a, b)
}