package main
import (
	"fmt"
	"net"
	"go_code/chatroom1.1/common/message"
	"go_code/chatroom1.1/server/process"
	//其实这里使用的是process文件夹下面把的的process2包
	//不然process包名与process协程的函数名一样的话，会发生编译错误
	"go_code/chatroom1.1/server/utils"
	"io"

)

//先创建一个processor的结构体
type Processor struct{
	Conn net.Conn
}

//编写一个ServerProcessMes函数
//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *Processor)serverProcessMes (mes *message.Message) (err error) {
	//fmt.Println(mes)
	switch mes.Type {
	case message.LoginMesType :		//处理登录
		//创建一个实例
		userPr := &process2.UserProcess {
			Conn : this.Conn,
		}
		//调用
		err = userPr.ServerProcessLogin(mes)
	case message.RegisterMesType :		//处理注册
		//创建一个实例
		userPr := &process2.UserProcess {
			Conn : this.Conn,
		}
		//调用
		err = userPr.ServerProcessRegister(mes)
	case message.NotifyUserStatusMesType :		//处理用户退出消息群发
		//创建一个实例
		userPr := &process2.UserProcess {
			Conn : this.Conn,
		}
		//调用
		err = userPr.NotifyOutOnilneUser(mes)
	case message.SmsMesType : //处理群发消息
		//创建一个实例
		smsProcess := &process2.SmsProcess{}
		err = smsProcess.SendGroupMes(mes)
	default :
		fmt.Println("消息类型不存在，无法处理。。。。。")
	}
	return 
}

//函数名不要与被程序调用的包名一样，会报错的
func (this *Processor)process2() (err error) {

	//循环读取的客户发送的信息,调用封装函数模式
	for {
		//创建一个Transfer实例
		transf := &utils.Transfer {
			Conn : this.Conn,
		}
		mes, err := transf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端正常退出。。。。。")
				return err
			} else {
				fmt.Println("readPkg err = ", err)
				return err 
			}
		}
		fmt.Println("mes = ", mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			fmt.Println("serverProcessMes err = ", err)
			return err 
		}
	}

}