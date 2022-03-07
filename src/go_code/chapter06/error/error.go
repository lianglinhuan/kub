package main

import (
	"fmt"
	"errors"
)

func test() {
	//通过defer释放资源的机制，将匿名函数压入堆栈中，等到test函数退出时，在执行匿名函数
	defer func() {
		err := recover() //用recover来回收错误
		if err != nil {
			fmt.Println("err=", err) //若有错误，则输出错误
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=",res)
}

//该函数判断读取文件是否一致
func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil  //文件一致，返回nil
	} else {
		//w文件不一致，返回错误信息
		return errors.New("读取文件错误。。。")
	}
}

func test1() {
	err := readConf("config.ini")
	if err != nil {
		//如果文件不一致，就输出readConf函数的错误信息，并终止程序
		panic(err)
	}
	fmt.Println("test1继续执行")
}

func  main() {

	test()
	fmt.Println("main下的代码")

	test1()
	fmt.Println("main下的代码~~~")

}