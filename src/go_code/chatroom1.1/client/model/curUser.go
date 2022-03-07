package model
import (
	"net"
	"go_code/chatroom1.1/common/message"
)
//该结构体的作用就是，维护客户端与服务器之间的连接
//体现在smsProcess.go文件上
type CurUser struct {
	Conn net.Conn
	message.User
}