package main
import "fmt"

func main() {
	var i = 1
	var j = 2
	var r = i + j //做加法运算
	fmt.Println("r=", r)

	var(
		str1 = "hello"
		str2 = "world"
		res = str1 + str2 //做拼接操作
	)
	fmt.Println("res=", res)
}