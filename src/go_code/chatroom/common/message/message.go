package message

const (
	LoginMesType    =  "LoginMes"
	LoginResMesType =  "LoginResMes"
	RegisterMesType =  "RegisterMes"
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
	Error string `json:"error"`
}

//定义一个注册用户的结构体类型数据
type RegisterMes struct {

}