package main
import (
	"fmt"
	_"os"
	"go_code/chatroom1.1/client/process"
)
var userId int
var userPwd string
var userName string
func main() {
	//设置一个变量存放客户的选项内容
	var key int
	//判断是否还继续显示菜单
	var loop = true

	for loop {
		fmt.Println("------------------欢迎登录多人聊天系统---------------")
		fmt.Println("\t\t\t 1 登录聊天系统")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t 请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		//fmt.Scanln(&key)也是可以的
		switch key {
		case 1:
			fmt.Println("登录聊天系统")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
			//也可以使用os.Exit()可以直接终止程序的运行,不过得引用os包
			//os.Exit()
		default:
			fmt.Println("您的输入有误，请重新输入")
		}
	}
	//登录用户，显示登录用户的提示信息
	if key == 1 {
		fmt.Println("请输入用户的id")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户的密码")
		fmt.Scanf("%s\n", &userPwd)
		//先把登录的函数，写到另外一个文件，比如login.go
		//因为在同一个文件夹下面，而且都是打main包，所以可以直接使用
		// err := login(userId, userPwd)
		// if err != nil {
		// 	fmt.Println("登录失败")
		// } else {
		// 	fmt.Println("登录成功")
		// }
		userProc := &process.UserProcess{

		}
		userProc.Login(userId, userPwd)
	} else if key == 2 {
		fmt.Println("进行用户注册。。。。。")
		fmt.Println("请输入注册用户的id:")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户注册的密码:")
		fmt.Scanf("%s\n", &userPwd)
		fmt.Println("请输入用户注册的名字(nickname):")
		fmt.Scanf("%s\n", &userName)

		userProc := &process.UserProcess{

		}
		userProc.Register(userId, userPwd, userName)
	}
}