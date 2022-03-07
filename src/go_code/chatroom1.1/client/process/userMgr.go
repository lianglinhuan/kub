package process

import (
	"fmt"
	"go_code/chatroom1.1/common/message"
	"go_code/chatroom1.1/client/utils"
	"go_code/chatroom1.1/client/model"
	"encoding/json"
	"net"

)

//客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser  //我们在用户登录成功后，完成对CurUser初始化

//编写一个方法，处理返回NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	//若发现用户为新登录用户
	if !ok {
		user = &message.User {
			UserId : notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
}

//在客户端显示当前在线用户
func outputOnlineUser() {
	//遍历一把onlineUsers
	fmt.Println("当前在线用户列表：")
	for id, _ := range onlineUsers {
		fmt.Println("用户id：\t", id)
	}
}


//编写一个方法，处理返回NotifyUserStatusMes
func delUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	//1.将下线的用户从onlineUsers中删除
	delete(onlineUsers, notifyUserStatusMes.UserId)
}


//客户端退出，客户端向服务其传递消息
func useroutOnlineUser(userId int) {
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
	mes.Type = message.NotifyUserStatusMesType
	
	//3.创建一个notifyUserStatusMes结构体
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOffline

	//4.notifyUserStatusMes
	data, err := json.Marshal(notifyUserStatusMes)
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
		fmt.Println("用户退出发送消息 err = ", err)
		return
	}
}

