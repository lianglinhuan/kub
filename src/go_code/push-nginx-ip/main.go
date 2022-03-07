package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"strings"
)

func readfile(filename string) (err error) {
	//打开文件，并得到一个file的句柄，操作文件就是通过该句柄
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("open file error = %v", err)
		return
	}

	//函数调用结束即关闭文件
	defer file.Close()

	//缓存读取文件信息
	reader := bufio.NewReader(file)
	// reader1 := bufio.NewReader(file)
	num := 1
	httpmin := 0
	httpmax := 0
	streammin := 0
	streammax := 0
	// tag := 0


	// for {
	// 	//使用ReadString()函数来读取文件的内容
	// 	//其函数传入的参数为末行的结束符号，用单引号引起来
		
	// 	//func (b *Reader) ReadString(delim byte) (line string, err error)
	// 	//ReadString读取直到第一次遇到delim字节，返回一个包含已读取的数据和delim字节的字符串。
	// 	//如果ReadString方法在读取到delim之前遇到了错误，它会返回在错误之前读取的数据以及该错误（一般是io.EOF）。
	// 	//当且仅当ReadString方法返回的切片不以delim结尾时，会返回一个非nil的错误。+6
	// 	str, err := reader.ReadString('\n')

	// 	//当err接收到的值为io.EOF时，就表示文件读取结束
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Print(str)
	// 	// num++
	// }
	// // fmt.Println(num)
	// // fmt.Println("文件读取结束")
	// for {
	// 	str, err := reader.ReadString('\n')
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	//判断http或者stream的协议作用范围
	// 	//通过strings包下面的func HasSuffix(s, suffix string) bool，函数判断字符串后缀是否与我们期望的值相匹配
	// 	if strings.HasSuffix(str, "http {\r\n") {
	// 		fmt.Print(str)
	// 		httpmin = num
			
	// 	}else if strings.HasSuffix(str, "stream {\r\n") {
	// 		fmt.Print(str)
	// 		streammin = num
	// 	}
	// 	num++
	// }
	// //下面是判断http还是stream协议在前后位置
	// if streammin > httpmin {
	// 	httpmax = streammin - 1
	// 	streammax = num
	// } else {
	// 	streammax = httpmin - 1
	// 	httpmax = num
	// }
	
	fmt.Printf("httpmin = %d\nhttpmax = %d\nstreammin = %d\nstreammax = %d\n",
		httpmin, httpmax, streammin, streammax)
	//fmt.Println(num)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println(err)
			break
		}
		if strings.HasSuffix(str, "server {\r\n") {
			fmt.Println(str)
			// fmt.Println(num)
		}
		num++
	}
	fmt.Println(num)
	return
}

func main() {
	filename := "f:/nginx_ip/peizhi.txt"

	err := readfile(filename)
	if err != nil {
		fmt.Println("文件读取错误  ", err)
	}
	
}