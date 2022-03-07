package main 

import (
	"fmt"
	"net"
	"encoding/json"
	"encoding/binary"
	"go_code/chatroom/common/message"
	_"errors"
	"io"
	"time"
)

func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()

	//循环读取的客户发送的信息,调用封装函数模式
	for {

		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端正常退出。。。。。")
				return
			} else {
				fmt.Println("readPkg err = ", err)
				return
			}
		}
		fmt.Println("mes = ", mes)
		err = serverProcessMes(conn, &mes)
		if err != nil {
			fmt.Println("serverProcessMes err = ", err)
			return
		}
	}

}


//定义一个读取客户端数据的函数
func readPkg(conn net.Conn) (mes message.Message, err error) {
	//1.获取客户端发送信息的数据字节长度
	buf := make([]byte, 8096)
	fmt.Println("....等待读取客户端的信息。。。。")
	//这里conn.Raed在conn没有关闭的情况下，才会阻塞
	//如果客户端关闭了conn，则就不会阻塞，要是不阻塞，程序就会往下执行
	//所以在客户端关闭前来个延时，等待服务器处理好数据
	_, err = conn.Read(buf[:4])
	if err != nil {
		if err == io.EOF {
			return
		}else {
			fmt.Println("conn.Read err = ", err)
			//自定义错误
			// err = errors.New("read pkg header error")
			return
		}
	}
	// fmt.Println("读取到的buf = ", buf[:4])

	//2.根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])

	//3.根据pkgLen读取消息内容
	//这里做一个特殊说明！！！！，conn.read()函数不是从buf[:pkgLen]读取数据的，
	//而是从conn读取数据后，在将读取到的数据信息传入buf[:pkgLen]切片里面的
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read fail err = ", err)
		// err = errors.New("read pkg body error")
		return
	}

	//4.将获取到的数据进行序列化
	//4.1 先将数据转换成message.Message的结构体数据类型
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}
	return

}


//编写一个ServerProcessMes函数
//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func serverProcessMes (conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType :
		//处理登录
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType :
		//处理注册
	default :
		fmt.Println("消息类型不存在，无法处理。。。。。")
	}
	return 
}


//编写一个函数serverProcessLogin()函数，专门处理登录请求
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	//1.先从mes中取出mes.Data，并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err = ", err)
		return
	}

	//2.对获取客户端的信息，做出判断
	//2.1 先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//2.2 再声明一个LoginResMes，并完成赋值
	var loginResMes message.LoginResMes

	//2.3 对客户端输入的Id和Pwd作出正确性的判断
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//合法
		loginResMes.Code = 200
	} else {
		//不合法
		loginResMes.Code = 500  //500状态码，表示该用户不存在
		loginResMes.Error = "该用户不存在，请注册再使用。。。。"
	}

	//3. 对服务器要向客户端返回的数据进行序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal([]byte(loginResMes)) err = ", err)
		return
	}

	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("服务器json.Marshal(resMes) err = ", err)
		return
	}

	//4.发送，服务器向客户端发送信息
	err = writePkg(conn, data)
	if err != nil {
		fmt.Println("服务器writePkg(conn, data) err = ", err)
		return
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	//1. 到这时候data就是我们要发送的消息
	//1.1 先把data的长度发送给服务器
	//先获取data的长度，然后将其转换成一个表示长度的byte切片，
	//因为conn.Write([]byte)所传递的数据是[]byte切片
	var pakLen uint32
	pakLen = uint32(len(data))
	var buf [4]byte
	//binary.BigEndian.PutUint32([]byte, uint32)
	//1.2 该方法可以将uint32类型的数据转换为[]byte类型，并存放在切片当中
	binary.BigEndian.PutUint32(buf[0:4], pakLen)

	//2. 发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf[:4]) err = ", err)
		return
	}
	// fmt.Printf("客户端，发送消息的长度为%d\n 内容是%s\n", len(data), string(data))

	//3.发送消息本身
	n, err = conn.Write(data)
	if n != int(pakLen) || err != nil {
		fmt.Println("conn.Write(data) err = ", err)
		return
	}
	return
}

func main() {
	fmt.Println("服务器在8889端口监听。。。。")
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