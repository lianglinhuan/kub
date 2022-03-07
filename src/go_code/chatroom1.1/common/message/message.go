package message

const (
	LoginMesType    =  "LoginMes"
	LoginResMesType =  "LoginResMes"
	RegisterMesType =  "RegisterMes"
	RegisterResMesType =  "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes" //用户上线
	//NotifyUserOutMesType = "NotifyUserOutMes" //用户下线
	SmsMesType = "SmsMes"
)

//这里定义几个用户状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusystatus 
)

type Message struct {
	Type string `json:"type`//消息类型
	Data string `json:"data"`//消息的数据
}

//定义客户端要发送的结构体类型数据
type LoginMes struct {
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}

//定义一个服务器要返回的数据类型
type LoginResMes struct {
	Code int `json:"code"`//返回状态码 500表示该用户未注册 200表示登录成功
	UsersId []int  //定义切片存放登录用户的状态
	Error string `json:"error"`
}

//定义一个注册用户的结构体类型数据
type RegisterMes struct {
	User User `json:"user"` //类型就是User结构体
}

type RegisterResMes struct {
	Code int `json:"code"`//返回状态码 400表示该用户已存在 200表示注册成功
	Error string `json:"error"`
}

//为了配合服务器端推送用户状态变化的信息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`  //用户id
	Status int `json:"status"`  //用户的状态
}

//增加一个Sms   //发送的消息
type SmsMes struct {
	Content string `json:"content`
	User //使用匿名结构体，继承
}