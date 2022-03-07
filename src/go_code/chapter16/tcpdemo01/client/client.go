package main 
import (
	"fmt"
	"net"
	"bufio"
	"os"
	_"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err = ", err)
		return
	}
	var num int
	// fmt.Println("conn 成功 = ", conn)

	//功能1：客户端可以发送单行数据，然后就退出
	//1.os/Stdin代表标准输入【终端】
	reader := bufio.NewReader(os.Stdin)
	for {
	    //2.从终端读取一行输入，并准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err = ", err)
		}

		//方法一：判断客户端输入exit就退出程序
		if line == "exit\r\n" {
			break
		}
		//方法二：使用strings.Trim()函数去掉字符串后面的回车"\r\n"
		// line = strings.Trim(line, "\r\n")
		// if line == "exit" {
		// 	break
		// }

		//3.将line，即终端输入的内容发送给服务器
		//conn.Write()函数传入的数据是[]byte类型，而line是字符串类型，
		//所以你要将line强制转换为[]byte，切片类型
		n, err := conn.Write([]byte(line))
		// _, err := conn.Write([]byte(line + "\n")) //将字符串删掉的会出补回
		if err != nil {
			fmt.Println("conn.Write err = ", err)
		}
		num = num + n
    }
fmt.Printf("客户端发送送了 %d 字节的数据，并退出", num)
}