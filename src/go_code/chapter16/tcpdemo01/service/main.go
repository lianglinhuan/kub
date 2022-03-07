package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	//这里我们循环的接收客户端发送的数据
	defer conn.Close() //关闭conn

	for {
		//1.创建一个新的切片
		buf := make([]byte, 1024)
		// fmt.Println("服务器在等待客户%s 端发送信息" , conn.RemoteAddr(). String() )
		//2.从conn取出数据信息
		//conn.Read(buf),等待客户端通过conn发送信息，然后将信息存放在buf切片中
		//如果客户端没有write【发送】，name协程就会阻塞在这里
		//n,是返回的客户端发送数据信息的字节数
		n, err := conn.Read(buf) 
		if err != nil {
			fmt.Println("服务器的Read err= ", err)
			return
		}
		//3.在终端输出客户端所发送的信息
		//因为buf是切片类型，所以在输出时需要将buf强制转换为string类型
		//    注意！！！要限定切片的输出字节，不然默认切片全输出的话，
		//    切片很多位置因为没有存放数据，就会输出很多奇怪的符号,得不到最终想要的结果
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听。。。。")
	//net.Listen("tcp", "0.0.0.0:8888")
	//1.tcp表示使用网络协议是tcp
	//2.0.0.0.0:8888表示在本地监听8888端口
	//0.0.0.0支持本Ipv4和Ipv6,127.0.0.1支持Ipv4,都是表示本地地址
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	//要是发生端口监听错误直接退出程序
	if err != nil {
		fmt.Println("listen err = ", err)
		return
	}
	// //测试监听端口函数返回的是什么
	// fmt.Printf("listen = %v\n", listen)

	defer listen.Close() //延时关闭listen

	//循环等待客户端连接
	for {
		//等待客户端连接
		fmt.Println("等待客户端连接。。。。")
		//1Accept()该函数等待客户端连接（堵塞），直到客户端连接才退出该函数，程序往下执行
		//返回的conn是一个套接字类型的数据，为引用类型
		conn, err := listen.Accept()
		if err != nil {
			 fmt.Println("Accept() err = ", err)
		} else {
			//conn.RemoteAddr(). String()该函数的作用是返回远程连接客户端的ip
			fmt.Printf("Accept() success comm = %v 客户端ip = %v \n", conn, conn.RemoteAddr(). String())
		}
		go process(conn)
	}
}