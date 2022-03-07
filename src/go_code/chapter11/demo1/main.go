package main
import (
	"fmt"
)

//定义一个结构体存放银行卡的信息
type Account struct{
	Id int
	Monery float64
	Pwd int
}

//存钱方法
func (acc *Account)SaveMonery(pwd int, monery float64) {
	if pwd != acc.Pwd {
		fmt.Println("输入密码不正确")
		return
	}
	if monery > 0 {
		acc.Monery += monery
	}
	fmt.Println("存钱成功，目前的余额为", acc.Monery)
}

//取钱方法
func (acc *Account)WithDraw(pwd int, monery float64) {
	if pwd != acc.Pwd {
		fmt.Println("输入密码不正确")
		return
	}
	if monery > 0 && monery <= acc.Monery{
		acc.Monery -= monery
	} else {
		fmt.Println("金额超出余额，请重新输入金额")
		return
	}
	fmt.Println("取钱成功，目前的余额为", acc.Monery)
}

//查询方法
func (acc *Account)Query(pwd int) {
	if pwd != acc.Pwd {
		fmt.Println("输入密码不正确")
	}
	fmt.Println("余额为", acc.Monery)
}


func main() {
	acc := Account{
		Id : 1234561234,
		Pwd : 6666,
		Monery : 100.0,
	}

	// acc.Query(6666)
	// acc.WithDraw(6666, 50)
	// acc.SaveMonery(6666, 50)
	
	//还可以改进一下的，就是在最开始的时候输入银行卡号，和密码，先对他们进行判断再往后执行
	for {
		fmt.Println("请输入根据功能选项选择功能模块：1：存钱，2：取钱，3：查询，4：退出 ")
		var n int
		var monery float64
		fmt.Scanln(&n)
		switch n {
		case 1:
			fmt.Println("请输入你所想要存储的金额")
			fmt.Scanln(&monery)
			acc.SaveMonery(6666, monery)
		case 2:
			fmt.Println("请输入你所想要取出的金额")
			fmt.Scanln(&monery)
			acc.WithDraw(6666, monery)
		case 3:
			acc.Query(6666)
		case 4:
			//好像使用break是不能退出for死循环的

		}
		if n == 4 {
			break
		}

	}
}