package utils
import (
	"fmt"
)

type FimaryAccount struct {
	
	//定义一个key字段用来选择功能模块
	key  string
	//定义一个loop字段，用来退出for死循环
	loop bool
	//定义一个字段，用来重复确认是否退出for死循环
	sure string

	//余额字段
	balance float32
	//收入金额字段
	monery float32
	//收入支出说明字段
	note string
	//账单序号字段
	num int
	//家庭收支字段
	details string

	zhanghao string
	mima string
}

func NewfimaryAccount() *FimaryAccount {
	return &FimaryAccount{
		key : "",
		loop : false,
		balance : 10000.0,
		monery : 0.0,
		note : "",
		num : 0,
		details : "账单序号\t账单类型\t账户余额\t账户收支\t说  明",
		zhanghao : "",
		mima : "",
	}
}
func (this *FimaryAccount) XianShi() {
		fmt.Println("............家庭收支明细...........\n")
		//要是没有收支的话，不用打印收支信息，并提示
		if this.num != 0 {
			fmt.Println(this.details)
		}else {
			fmt.Println("当前没有收支，请添加收支情况。")
		}
}

func (this *FimaryAccount) Income() {
		fmt.Print("请输入收入的金额：")
		fmt.Scanln(&this.monery)
		this.balance += this.monery
		this.num++
		fmt.Print("输入收入的说明：")
		fmt.Scanln(&this.note)
		//Sprintln()函数可以将输出的内容转化为字符串类型
		//这样就可以很好的与details字符串进行拼接
		this.details += fmt.Sprintf("\n%v         \t收入        \t%v       \t%v       \t%v", this.num, this.balance, this.monery, this.note )
}

func (this *FimaryAccount) Out() {
		fmt.Print("请输入支出的金额：")
		fmt.Scanln(&this.monery)
		//在这里弄一个判断，若是支出的金额大于余额，则退出switch
		if this.monery > this.balance {
			fmt.Println("您的余额不足，请重新输入")
			//break
		}
		this.balance -= this.monery
		this.num++
		fmt.Print("请输入支出的说明：")
		fmt.Scanln(&this.note)
		this.details += fmt.Sprintf("\n%v         \t支出        \t%v       \t%v       \t%v", this.num, this.balance, this.monery, this.note )
}

func (this *FimaryAccount) Exit() {
		fmt.Print("您确定要退出吗？(y/n):")
		fmt.Scanln(&this.sure)
		//通过一个死循环判断输入的值是否是y/n
		//要是y则直接退出程序
		//要是n则不退出该程序
		//输入的值不是y/n中的一个，则系统提示输入错误，请重新输入，直到输入的值正确为止
		for {
		if this.sure == "y" {
			this.loop = true
			break
		}else {
			if this.sure == "n" {
				break
			}else {
				fmt.Print("输入错误！请重新输入(y/n):")
				fmt.Scanln(&this.sure)
			}
		}
	}
}

func (this *FimaryAccount) Maxmain() {
	for {
		fmt.Println("\n............家庭收支软件...........\n")
		fmt.Println("            1.家庭收支明细\n")
		fmt.Println("            2.家庭收入\n")
		fmt.Println("            3.家庭支出\n")
		//Println()函数默认是加换行的，而Print()函数是没有加换行的
		fmt.Print("请选择<1-4>:")
	
		//键盘输入变量值，得使用变量指针传入
		fmt.Scanln(&this.key)
	
		switch this.key {
			case "1" :
				this.XianShi()
			case "2" :
				this.Income()
			case "3" :
				this.Out()
			case "4" :
				this.Exit()
			default :
				fmt.Println("输入功能错误请重新输入! ! !\n")
		}
		if this.loop  {
			break
		}
	}
	fmt.Println("您已成功退出家庭收支软件\n")
}

func (this *FimaryAccount) Maxmax() {
		for {
			fmt.Print("请输入您的账号：")
			fmt.Scanln(&this.zhanghao)
			fmt.Print("请输入您的密码：")
			fmt.Scanln(&this.mima)
			if this.zhanghao == "llh" && this.mima == "061810" {
				this.Maxmain()
			}else {
				fmt.Println("您输入的账号或者密码错误！！请重新输入！")
			}
			if this.loop {
				break
			}
		}
		
}

