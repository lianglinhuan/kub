package process

import (
	"fmt"
	"os"
	"go_code/chatroom1.1/common/message"
	"go_code/chatroom1.1/client/utils"
	"encoding/json"
	"net"
	"time"

)

//显示登录成功后的界面。。
func ShowMenu(userId int) {
	fmt.Println("---------恭喜xxx登录成功---------")
	fmt.Println("-------您可以选择的功能如下---------")
	fmt.Println("------1. 显示在线用户列表---------")
	fmt.Println("------2. 发送消息----------------")
	fmt.Println("------3. 信息列表----------------")
	fmt.Println("------4. 退出系统----------------")
	fmt.Println("请选择(1-4):")
	var key int

	//发送消息
	var content string
	smsProcess := &SmsProcess{}

	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表-")
		outputOnlineUser()
	case 2:
		//发送消息
		fmt.Println("你想对大家说些什么：")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("消息列表")
	case 4:
		fmt.Println("你选择退出了系统。。。。")
		//向服务器发送消息，userId该用户退出
		useroutOnlineUser(userId)
		os.Exit(0) //强制退出
	default :
		fmt.Println("你输入的选项不正确。。。")
	}
}

//和服务器保持通讯
func serverProcessMes(conn net.Conn) {
	transf := &utils.Transfer {
		Conn : conn,
	}
	for {
		//延时一秒
		time.Sleep(1 * time.Second)

		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := transf.ReadPkg()
		if err != nil {
			fmt.Println("transf.ReadPkg err = ", err)
			return
		}
		//如果读取到消息，进行下一步处理
		switch mes.Type {
		case message.NotifyUserStatusMesType : //有人上线
			//处理
			//1.取出NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			//2.把这个用户的信息，状态保存到客户map[int]*User中
			if notifyUserStatusMes.Status == message.UserOnline { //有人上线
				//1.通知在线的用户，有一个用户上线
				fmt.Printf("用户%d, 上线。。\n", notifyUserStatusMes.UserId)
				updateUserStatus(&notifyUserStatusMes) 
			} else if notifyUserStatusMes.Status == message.UserOffline { //有人下线
				//1.通知在线用户，有一个用户已退出
				fmt.Printf("用户%d, 已下线。\n", notifyUserStatusMes.UserId)
				delUserStatus(&notifyUserStatusMes)
			}
		case message.SmsMesType : //有人群发消息
			outputGroupMes(&mes)
		default :
			fmt.Println("服务器端返回的是未知的消息类型")
		}
		//fmt.Printf("mes = %v\n", mes)
	}
}