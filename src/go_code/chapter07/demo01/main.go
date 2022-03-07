package main

import (
	"fmt"
	"time"
	"math/rand"
)


func main() {
	var arr [26]int
	var n int = 65
	for i := 0; i < 26; i++ {
		arr[i] = n
		n++
	}

	//for - range遍历数组模式，比传统遍历模式方便
	for _, value := range arr{

		//将整型数据类型，转换成字符串类型
		// n := string(value) 
		// fmt.Printf("%s\n", n)

		//%c格式可以输出ASSIC码值
		fmt.Printf("%c\n", value)
	}

	// var num [4]int
	// for i := 0; i < 4; i++ {
	// 	fmt.Printf("请输入第%v个数组的值\n", i)
	// 	fmt.Scanf(&num[i])
	// }

	// for i := 0; i < 3; i++ {
	// 	if num[i] > num[i+1] {
	// 		num[i+1] = num[i]
	// 	}
	// 	if i + 1 == 3 {
	// 		fmt.Printf("最大值为%v\n", num[i+1])
	// 	}
	// }

	//要求：随机生成五个数，并将其翻转打印
	//1.随机生成五个数，rand.Intn()函数
	//1.当我们得到随机数后，就放到一个数组，int数组
	//3.反转打印

	var intArr3 [5]int
	//为了每次生成的随机数不一样，需要给一个seed（种子）
	//不然的话，第一次生成的随机数会被保留下来

	len := len(intArr3)
	//将len()函数的结果导出给一个变量，方便以后使用
	//因为该函数是一个内置的函数，每一次调用都是需要调动资源来进行运算，这样会耗费资源

	rand.Seed(time.Now().Unix())
	for i := 0; i < len; i++{
		//通过函数参数限定生成随机数的范围
		intArr3[i] = rand.Intn(100)
	}
	fmt.Println(intArr3)

	//通过len()检测出数组的大小
	for i := 0; i < len/2; i++ {
		t := intArr3[i]
		intArr3[i] = intArr3[len - i -1]
		intArr3[len - i -1] = t
	}
	fmt.Println(intArr3)


}