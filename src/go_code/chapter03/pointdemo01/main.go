package main
import (
	"fmt"
)

func main() {

	// var num int = 9
	// fmt.Printf("num address=%v\n", num)

	// var ptr *int
	// ptr = &num
	// *ptr = 10 //这里修改时，会修改num的值
	// fmt.Printf("num =%v\n", num)

	slice := []int{1,3,5,6}

	for index, val := range slice{
		fmt.Printf("下标：%v  数值：%v\n", index, val)
	}
}