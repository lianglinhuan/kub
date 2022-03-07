package message

//定义一个用户的结构体

type User struct{
	
	//注意！！！！
	//为了序列化和反序列化成，我们必须保证成功
	//用户信息的json字符串的key和结构体的字段对应的tag名字一样
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
	UserStatus int `json:"userStatus"` //用户状态。。。
}