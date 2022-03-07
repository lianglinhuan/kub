package main
import (
	"fmt"
	"go_code/chatroom/common/message"
	"encoding/json"
	"encoding/binary"
	"net"
	_"time"
)

func login(userId int, userPwd string) (err error) {
	//1.连接到服务器端
	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}
	//延时关闭!!!!!
	defer conn.Close()

	//2.准备通过conn发送给服务器
	var mes message.Message
	mes.Type = message.LoginMesType
	//3.创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4.将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	//5.吧data传递给mes.Data字段
	mes.Data = string(data)

	//6.将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	//7.到这时候data就是我们要发送的消息
	//7.1 先把data的长度发送给服务器
	//先获取data的长度，然后将其转换成一个表示长度的byte切片，
	//因为conn.Write([]byte)所传递的数据是[]byte切片
	var pakLen uint32
	pakLen = uint32(len(data))
	var buf [4]byte
	//binary.BigEndian.PutUint32([]byte, uint32)
	//该方法可以将uint32类型的数据转换为[]byte类型，并存放在切片当中
	binary.BigEndian.PutUint32(buf[0:4], pakLen)

	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf[:4]) err = ", err)
		return
	}
	// fmt.Printf("客户端，发送消息的长度为%d\n 内容是%s\n", len(data), string(data))

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err = ", err)
		return
	}

	// //休眠20
	// time.Sleep(10 * time.Second)
	// fmt.Println("休眠了10秒")
	mes, err = readPkg(conn)
	if err != nil{
		fmt.Println("客户端readPkg(conn) err = ", err)
		return
	}

	//将mes的Data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}

	return
}