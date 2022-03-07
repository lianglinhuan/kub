package main
import (
	"fmt"
)

//在函数定义的时候，定义变量数组是要定义好数组的长度
func number(arr *[5]int) {
	temp := 0
	for i := 0; i < len(*arr) - 1; i++ {
		for j := 0; j < len(*arr) -1 -i; j++ {
			if (*arr)[j] > (*arr)[j + 1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j + 1]
				(*arr)[j + 1] = temp
			}
		}
	}
}

func main() {

	var name string
	fmt.Println("请输入一个名字....")

	//这里用fmt.Scanf()函数是不行的，说是类型出现问题，以后少用看看
	fmt.Scanln(&name)
	fmt.Println(name)

	//数组定义的方法
	//var arr [5]int = [5]int{89, 23, 48, 8, 78}
	arr := [5]int{89, 23, 48, 8, 78}
	fmt.Println("arr=", arr)

	//要是想要改变数组的内容，需要地址进行传递
	number(&arr)
	fmt.Println("arr=", arr)

}