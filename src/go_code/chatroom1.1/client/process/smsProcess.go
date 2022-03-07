package process
import (
	"fmt"
	"go_code/chatroom1.1/common/message"
	"encoding/json"
	"go_code/chatroom1.1/client/utils"
	_"net"
)

type SmsProcess struct {

}

//发送群聊的消息
func (this *SmsProcess)SendGroupMes(content string)(err error) {
	
		//因为conn直接使用的是login登录时的conn，
		//CurUser.Conn = conn，CurUser是全局变量，因此可以直接使用而不需要再连接
		//体现了curUser.go文件下CurUser结构体的作用(维护客户端与服务器之间的连接)

		// //1.连接到服务器端
		// conn, err := net.Dial("tcp", "0.0.0.0:8889")
		// if err != nil {
		// 	fmt.Println("net.Dial err = ", err)
		// 	return
		// }
		// //延时关闭!!!!!
		// defer conn.Close()
	
		//2.准备通过conn发送给服务器
		var mes message.Message
		mes.Type = message.SmsMesType
		
		//3.创建一个notifyUserStatusMes结构体
		var smsMes message.SmsMes
		smsMes.Content = content
		smsMes.UserId = CurUser.UserId
		smsMes.UserStatus = CurUser.UserStatus
	
		//4.notifyUserStatusMes
		data, err := json.Marshal(smsMes)
		if err != nil {
			fmt.Println("json.Marshal err = ", err)
			return
		}
	
		//5.吧data传递给mes.Data字段
		mes.Data = string(data)

		//6.将mes进行序列化
		data, err = json.Marshal(mes)
		if err != nil {
			fmt.Println("SendGroupMes json.Marshal err = ", err)
			return
		}

		//7.创建一个实例
		transf := &utils.Transfer{
			Conn : CurUser.Conn,
		}
		//8.客户端写入数据调用
		err = transf.WritePkg(data)
		if err != nil{
			fmt.Println("群发发送消息 err = ", err)
			return
		}
		return
}