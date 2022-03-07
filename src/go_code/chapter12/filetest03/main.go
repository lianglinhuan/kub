package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	filename := "f:/test01.txt"

	//打开文件
	//func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
	//该函数的传入参数有三个，第一个是字符串类型的文件名称，第二个是整型的以什么方式打开文件（其参数定义在os包里面）
	//O_RDONLY  只读
	//O_WRONLY  只写
	//O_RDWR    读写
	//O_APPEND  文件末尾添加
	//O_CREATE  要是文件不存在则创建文件
	//O_EXCL    与O_CREATE一起使用，文件必须不存在，然后创建
	//O_SYNC    为同步I/O打开
	//O_TRUNC   清除文件的内容
	//第三个是文件的权限（针对Liunx系统的，与windows系统无关）
	//返回的是一个File结构体的指针类型（句柄）
	file, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE , 0666 )
	if err != nil {
		fmt.Println(err)
	}

	//将关闭文件推入defer栈中
	defer file.Close()

	//定义好所要输入给文件的内容
	//有些编辑器识别\r ，有些识别\n为回车，最好两个都写，即\r\n
	str := "hello, world!\r\n"

	//还是以缓冲区的方式来对文件内容进行添加
	//func NewWriter(w io.Writer) *Writer，该函数的作用是，通过传入一个接口(接口可以接收所有的数据类型)
	//从而返回具有默认缓存的一个Reader结构体的指针类型（与NewReader（）函数十分相似，几乎一模一样）
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		//func (b *Writer) WriteString(s string) (int, error)
		//WriteString写入一个字符串。返回写入的字节数。如果返回值nn < len(s)，还会返回一个错误说明原因
		writer.WriteString(str)
	}

	//因为以上的操作都是在缓存中进行的，要是不使用func (b *Writer) Flush() error，
	//Flush方法将缓冲中的数据写入下层的io.Writer接口。
	//即该函数对缓存的数据刷新，而使得缓冲区的数据传入到文件中去。
	writer.Flush()
}