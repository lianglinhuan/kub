package process2
import (
	"fmt"
	"net"
	"encoding/json"
	"go_code/chatroom1.1/common/message"
	"go_code/chatroom1.1/server/utils"
	"go_code/chatroom1.1/server/model"
	"time"
)

type UserProcess struct {
	Conn net.Conn
	//增加一个字段，表示该Conn是哪个用户
	UserId int
}

//通知所有在线用户，我退出
func (this *UserProcess)NotifyOutOnilneUser( mes *message.Message) (err error) {
	//1.先从mes中取出mes.Data，并直接反序列化成notifyUserStatusMes
	var notifyUserStatusMes message.NotifyUserStatusMes
	err = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err = ", err)
		return
	}

	//2.将得到的notifyUserStatusMes.UserId，即我自己的用户删除
	userMgr.DelOnlineUser(notifyUserStatusMes.UserId)
	for _, up := range userMgr.onlineUsers {
		//开始通知【单独写一个方法】
		up.NotifyMeOutOnline(notifyUserStatusMes.UserId)
	}
	return
}

func (this *UserProcess)NotifyMeOutOnline(userId int) {
	//组装NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOffline

	//将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	//将序列化后的notifyUserStatusMes赋值给mes.Data
	mes.Data = string(data)
	//对mes在此序列化，准备发送
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}

	//发送
	//创建发送的实例
	transf := &utils.Transfer {
		Conn : this.Conn,
	}
	err = transf.WritePkg(data)
	if err != nil {
		fmt.Println("onlineUsers writePkg(conn, data) err = ", err)
		return
	}
	return

}


//通知所有在线用户的方法
//userId要通知其他在线用户，我上线
func (this *UserProcess)NotifyOtherOnlineUser(userId int) {
	//遍历onlineUsers，然后一个一个的发送NotifyUserStatusMes
	for id, up := range userMgr.onlineUsers {
		//过滤自己
		if id == userId {
			return
		}
		//开始通知【单独写一个方法】
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess)NotifyMeOnline(userId int) {
	//组装NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	//将序列化后的notifyUserStatusMes赋值给mes.Data
	mes.Data = string(data)
	//对mes在此序列化，准备发送
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}

	//发送
	//创建发送的实例
	transf := &utils.Transfer {
		Conn : this.Conn,
	}
	err = transf.WritePkg(data)
	if err != nil {
		fmt.Println("onlineUsers writePkg(conn, data) err = ", err)
		return
	}
	return
}





//处理注册请求
func (this *UserProcess)ServerProcessRegister( mes *message.Message) (err error) {
	//1.先从mes中取出mes.Data，并直接反序列化成registerMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err = ", err)
		return
	}

	//2.对获取客户端的信息，做出判断
	//2.1 先声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXITS {
			registerResMes.Code = 505
			registerResMes.Error =  model.ERROR_USER_EXITS.Error()
		}else {
			registerResMes.Code = 506
			registerResMes.Error =  "注册发生未知错误。。。"
		}
	}else {
		registerResMes.Code = 200
	}

	//3. 对服务器要向客户端返回的数据进行序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal([]byte(registerResMes)) err = ", err)
		return
	}

	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("服务器json.Marshal(resMes) err = ", err)
		return
	}

	//4.发送，服务器向客户端发送信息
	transf := &utils.Transfer {
		Conn : this.Conn,
	}
	err = transf.WritePkg(data)
	if err != nil {
		fmt.Println("服务器writePkg(conn, data) err = ", err)
		return
	}
	return
	
}






//处理登录请求
//编写一个函数serverProcessLogin()函数，专门处理登录请求
func (this *UserProcess)ServerProcessLogin( mes *message.Message) (err error) {
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

	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {
		if err == model.ERROR_USER_NOTEXITS {
			loginResMes.Code = 500  //500为用户不存在
			loginResMes.Error = err.Error()
		}else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403  //403为密码错误
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误。。。。"
		}
	}else {
		loginResMes.Code = 200
		fmt.Println(user, "登录成功")
		
		//这里，因为用户登录成功，我们就吧该登录成功的用户放入到userMgr中
		//将登录成功的用户userId赋给this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)


		//将当前在线用户的id放入到loginResMes.UsersId
		//遍历userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
	}

	// //2.3 对客户端输入的Id和Pwd作出正确性的判断
	// if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	// 	//合法
	// 	loginResMes.Code = 200
	// } else {
	// 	//不合法
	// 	loginResMes.Code = 500  //500状态码，表示该用户不存在
	// 	loginResMes.Error = "该用户不存在，请注册再使用。。。。"
	// }

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
	transf := &utils.Transfer {
		Conn : this.Conn,
	}
	err = transf.WritePkg(data)
	if err != nil {
		fmt.Println("服务器writePkg(conn, data) err = ", err)
		return
	}

	//通知其他在线用户，我上线了
	this.NotifyOtherOnlineUser(loginMes.UserId)

	//延时一秒
	time.Sleep(5 * time.Second)
	return
}