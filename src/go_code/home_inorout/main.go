package main
import (
	"fmt"
)

func main() {

	//定义一个key变量用来选择功能模块
	key  := " "
	//定义一个loop变量，用来退出for死循环
	loop := true
	//定义一个变量，用来重复确认是否退出for死循环
	var sure string

	//余额变量
	balance := 10000.0
	//收入金额变量
	monery := 0.0
	//收入支出说明变量
	note := " "
	//账单序号变量
	num := 0
	//家庭收支说明
	details := "账单序号\t账单类型\t账户余额\t账户收支\t说  明"

	var zhanghao string
	var mima string

	for {
		fmt.Print("请输入您的账号：")
		fmt.Scanln(&zhanghao)
		fmt.Print("请输入您的密码：")
		fmt.Scanln(&mima)
		if zhanghao == "llh" && mima == "061810" {
			for {
				fmt.Println("\n............家庭收支软件...........\n")
				fmt.Println("            1.家庭收支明细\n")
				fmt.Println("            2.家庭收入\n")
				fmt.Println("            3.家庭支出\n")
				//Println()函数默认是加换行的，而Print()函数是没有加换行的
				fmt.Print("请选择<1-4>:")
		
				//键盘输入变量值，得使用变量指针传入
				fmt.Scanln(&key)
		
				switch key {
					case "1" :
						fmt.Println("............家庭收支明细...........\n")
						//要是没有收支的话，不用打印收支信息，并提示
						if num != 0 {
							fmt.Println(details)
						}else {
							fmt.Println("当前没有收支，请添加收支情况。")
						}
					case "2" :
						fmt.Print("请输入收入的金额：")
						fmt.Scanln(&monery)
						balance += monery
						num++
						fmt.Print("输入收入的说明：")
						fmt.Scanln(&note)
						//Sprintln()函数可以将输出的内容转化为字符串类型
						//这样就可以很好的与details字符串进行拼接
						details += fmt.Sprintf("\n%v         \t收入        \t%v       \t%v       \t%v", num, balance, monery, note )
					case "3" :
						fmt.Print("请输入支出的金额：")
						fmt.Scanln(&monery)
						//在这里弄一个判断，若是支出的金额大于余额，则退出switch
						if monery > balance {
							fmt.Println("您的余额不足，请重新输入")
							break
						}
						balance -= monery
						num++
						fmt.Print("请输入支出的说明：")
						fmt.Scanln(&note)
						details += fmt.Sprintf("\n%v         \t支出        \t%v       \t%v       \t%v", num, balance, monery, note )
					case "4" :
						fmt.Print("您确定要退出吗？(y/n):")
						fmt.Scanln(&sure)
						//通过一个死循环判断输入的值是否是y/n
						//要是y则直接退出程序
						//要是n则不退出该程序
						//输入的值不是y/n中的一个，则系统提示输入错误，请重新输入，直到输入的值正确为止
						for {
						if sure == "y" {
							loop = false
							break
						}else {
							if sure == "n" {
								break
							}else {
								fmt.Print("输入错误！请重新输入(y/n):")
								fmt.Scanln(&sure)
							}
						}
					}
					default :
						fmt.Println("输入功能错误请重新输入! ! !\n")
				}
				if !loop  {
					break
				}
			}
		}
		if !loop {
			break
		}
	}
	
	fmt.Println("您已成功退出家庭收支软件\n")
	
}