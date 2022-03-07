package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file := "f:/test.txt"

	//一次将整个文件读入到内存中，与带缓冲区的方法不同
	//适用于文件较小的情况
	
	//func ReadFile(filename string) ([]byte, error)，该函数传入的是一个字符串类型，
	//返回的是一个[]byte切片类型，其中，文件的打开和关闭都封装在该函数里面
	read, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println(err)
	}
	//因为read接收的是byte切片类型，所以输出显示的话需要使用string强制转换类型
	//不然会看到一连串的数字
	fmt.Println(read)
	fmt.Println(string(read))
}