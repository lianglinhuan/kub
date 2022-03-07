package main

import (
	"fmt"
	"os" //提供open，和close函数
	"bufio"
	"io"
)

func main() {
	//打开文件
	//1.file 可以叫做file指针
	//2.file 可以叫做file对象
	//3.file 可以叫做file文件句柄
	file, err := os.Open("f:/test.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("file =", file)

	//关闭文件
	//使用defer，将关闭文件操作推入defer栈中，待到函数结束后才释放出来执行
	//在使用文件时，要注意及时关闭文件，否则容易造成内存泄漏
	 defer file.Close()
	// if err != nil {
	// 	fmt.Println(err)
	// }


	
	//带缓冲区的读取数据结构（特点：数据分片段读取）
	//适用于文件较大的情况

	//func NewReader(rd io.Reader) *Reader 该函数的作用是，通过传入一个接口(接口可以接收所有的数据类型)
	//从而返回具有默认缓存的一个Reader结构体的指针类型
	reader := bufio.NewReader(file)

	for {
		//使用ReadString()函数来读取文件的内容
		//其函数传入的参数为末行的结束符号，用单引号引起来
		
		//func (b *Reader) ReadString(delim byte) (line string, err error)
		//ReadString读取直到第一次遇到delim字节，返回一个包含已读取的数据和delim字节的字符串。
		//如果ReadString方法在读取到delim之前遇到了错误，它会返回在错误之前读取的数据以及该错误（一般是io.EOF）。
		//当且仅当ReadString方法返回的切片不以delim结尾时，会返回一个非nil的错误。+6
		str, err := reader.ReadString('\n')

		//当err接收到的值为io.EOF时，就表示文件读取结束
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束")

}