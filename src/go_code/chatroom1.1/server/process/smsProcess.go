package process2
import (
	"fmt"
	"go_code/chatroom1.1/common/message"
	"go_code/chatroom1.1/server/utils"
	"encoding/json"
	"net"
)

type SmsProcess struct {

}

func (this *SmsProcess)SendGroupMes(mes *message.Message)(err error) {
	//遍历服务器端的onlineUsers map[int]*UserProcess
	//将消息转发出去
	//反序列化mes主要是为了得到smsMes.UserId，进而筛选出除了自己本身外所有的用户
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}

	//重新进行序列化
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserId {//过滤掉自己
			continue
		}
		err = this.SendMesToEachOnlineUser(data, up.Conn)
	}
	return
}

func (this *SmsProcess)SendMesToEachOnlineUser(data []byte, conn net.Conn)(err error) {
	//发送
	//创建发送的实例
	transf := &utils.Transfer {
		Conn : conn,
	}
	err = transf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息 writePkg(conn, data) err = ", err)
		return
	}
	return
}