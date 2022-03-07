package main 

import (
	"fmt"
	"net"
	"time"
	"go_code/chatroom1.1/server/model"
)

func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()

	//这里调用总控
	//因为在一个main包下面，可以直接调用Processor结构体
	processor := &Processor {
		Conn : conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误 err = ", err)
		return
	}
}

//这里我们编写一个函数，完成对UserDao的初始化任务
func initUserDao() {
	//这里的pool本身就是一个全局的变量

	model.MyUserDao = model.NewUserDao(pool)
}



func main() {
	//当服务器启动时，我们就去初始化我们的redis连接池
	//这里需要注意一个初始化顺序的问题
	//initPool要在initUserDao之前，因为后面会被使用到
	initPool("localhost:6379", 16, 0, 300 * time.Second)
	initUserDao()
	fmt.Println("新的结构。。。服务器在8889端口监听。。。。")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
	}
	//延时关闭
	defer listen.Close()

	//一旦监听成功，就等待客户端连接服务器
	for {
		fmt.Println()
		fmt.Println("等待客户端来连接服务器。。。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err = ", err)
		}
		//一旦连接成功，则启动一个协程与客户端保持通讯。。。
		go process(conn)

		//延时一秒
		time.Sleep(1 * time.Second)
	}
}