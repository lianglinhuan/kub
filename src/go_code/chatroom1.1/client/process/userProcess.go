package process
import (
	"fmt"
	"go_code/chatroom1.1/common/message"
	"go_code/chatroom1.1/client/utils"
	"encoding/json"
	"encoding/binary"
	"net"
	_"time"
)

type UserProcess struct{

}

func (this *UserProcess)Register(userId int, userPwd string, userName string)(err error) {
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
	mes.Type = message.RegisterMesType

	//3.创建一个registerMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4.registerMes
	data, err := json.Marshal(registerMes)
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

	//7.创建一个实例
	transf := &utils.Transfer{
		Conn : conn,
	}
	//8.客户端写入数据调用
	err = transf.WritePkg(data)
	if err != nil{
		fmt.Println("注册发送消息 err = ", err)
		return
	}
	//9.客户端读取数据调用
	mes, err = transf.ReadPkg()
	if err != nil{
		fmt.Println("注册读取信息 err = ", err)
		return
	}

	//将mes的Data部分反序列化成LoginResMes
	var loginResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("注册成功。。。。")
	} else {
		fmt.Println(loginResMes.Error)
	}
	return

}



func (this *UserProcess)Login(userId int, userPwd string) (err error) {
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
	fmt.Printf("客户端，发送消息的长度为%d\n 内容是%s\n", len(data), string(data))

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err = ", err)
		return
	}

	// //休眠20
	// time.Sleep(10 * time.Second)
	// fmt.Println("休眠了10秒")

	//创建一个实例
	transf := &utils.Transfer{
		Conn : conn,
	}
	//调用
	mes, err = transf.ReadPkg()
	if err != nil{
		fmt.Println("客户端readPkg(conn) err = ", err)
		return
	}

	//将mes的Data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		// fmt.Println("登录成功")

		//初始化CurUser, 发送消息
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		//在这里可以显示当前在线用户列表，遍历loginResMes.UsersId
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UsersId {

			//如果我们要求不显示自己在线
			if v == userId {
				continue
			}
			fmt.Printf("用户id:%d\n", v)

			//完成客户端的onlineUsers初始化
			user := &message.User {
				UserId : v,
				UserStatus : message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Printf("\n\n")

		//在这里我们还需要在客户端启动一个协程
		//该协程保持和服务器的通讯，如果服务器有数据推送给客户端
		//则接收并显示在客户的终端
		go serverProcessMes(conn)

		//1. 显示我们登录成功的菜单。。。
		for {
			//因为在同一个包里面，所以可以直接调用
			ShowMenu(userId)
		}
	} else {
		fmt.Println(loginResMes.Error)
	}

	

	return
}