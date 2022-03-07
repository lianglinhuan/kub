package main
import "fmt"

func main() {
	
	//字符串遍历方式1-传统方式
	//这个方法运行起来有问题打印完字符串后会强制结束程序，不能继续执行
	// var str string = "hello,world!"
	// for i := 0; 1 < len(str); i++ {
	// 	fmt.Printf("%c", str[i]) //使用到下标
	// }

	//传统方式，若是包含中文，那么传统遍历字符串会出现错误，因为
	//传统方式对字符串是按照字节来遍历，而一个汉字在utf8编码是对应3个字节
	//解决方法，需要将str转成[]rune切片

	// var str string = "hello,world!北京"
	// str2 := []rune(str)
	// for i := 0; 1 < len(str); i++ {
	// 	fmt.Printf("%c", str2[i]) //使用到下标
	// }


	//字符串遍历方式2-for-rang
	//按照字符来遍历的，与传统的区别
	str1 := "abc~ok上海"
	for index, val := range str1 {
		fmt.Printf("index=%d, val=%c \n", index, val)
	}

	var sum int = 0
	fmt.Printf("在1~100之间所有可以被9整除的整数\n")
	for i := 1; i <= 100; i++ {
		// if i >= 9 {  //取模并不需要担心之前的也会被选出出来
			if i % 9 == 0 {
				fmt.Printf("%v\n", i)
				sum += i
			}
		// }
	}
	fmt.Printf("其中可以被9整除的整数的和=%v\n", sum)

	var j1 int = 6
	for j :=0; j <= 6; j++ {
		fmt.Printf("%v + %v = %v\n", j, j1, j+j1)
		j1--
	}
}