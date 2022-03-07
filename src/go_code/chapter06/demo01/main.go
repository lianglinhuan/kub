package main
import (
	"fmt"
	"go_code/chapter06/utils"
)
//递归求斐波那契系数 1,1,2,3,5,8,13....
func test(n3 int) int{
	if n3 == 1 || n3 ==2 {
		return 1
	} else {
		return test(n3 - 1) + test(n3 - 2)				
	}

}
//一天吃一半再加多一个桃子，第十天就只剩下一个桃子
func test1(n4 int) int{
	var tao int
	if n4 == 1 {
		return 1
	} else {
		tao = (test1(n4 - 1) + 1) * 2
	}
	return tao
}
//计算桃子的方法2
func test2(n4 int) int{
	var tao int
	if n4 == 10 {
		return 1
	} else {
		tao = (test2(n4 + 1) + 1) * 2
	}
	return tao
}

func main() {

	var n1 float64 = 4.5
	var n2 float64 = 3.3
	var operator byte = '/'
	var result float64

	result = utils.Cal(n1, n2, operator)
	fmt.Println(result)

	var res int
	var n3 int = 8
	res = test(n3)
	fmt.Println(res)

	var tao int
	var n4 int = 10 //输入的天数是最后一天吃桃子的天数
	tao = test1(n4)
	fmt.Println(tao)

	var tao2 int
	var n5 int = 1 //输入的天数是从开始吃桃子那天开始
	tao2 = test2(n5)
	fmt.Println(tao2)

	
}
