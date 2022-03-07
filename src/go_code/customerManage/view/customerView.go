package main

import (
	"fmt"
	"go_code/customerManage/service"
	"go_code/customerManage/model"
)

type customerView struct {
	//接收用户输入
	key string
	//判断循环退出
	loop bool
	//该字段用于判断是否退出系统
	// exit string

	//定义该字段用来接收service包的List（）切片函数
	//因为service.NewCustomerService（）函数返回的是个指针
	customerService *service.CustomerService
}


//1.查询
//list（）函数可以对切片进行遍历
//其中的Getinfo（）函数用来对切片格式化输出
//Getinfo（）函数在model包下，方便代码逻辑的处理

func (this *customerView)list() {
	//首先获取当前在切片中所有的客户信息
	customers := this.customerService.List()
	fmt.Println("--------------客户列表-------------\n")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		//直接通过Getinfo()函数获取客户信息
		fmt.Println(customers[i].Getinfo())
	}
	//fmt.Println()
	fmt.Println("\n-------------客户列表完成-----------\n\n")
}



//2.添加客户信息

func (this *customerView) add() {
	fmt.Println("-------------添加客户--------------\n")
	fmt.Print("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Print("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("邮箱：")
	email := ""
	fmt.Scanln(&email)

	//NewCustomer()函数返回的是一个结构体类型
	customer := model.NewCustomer(name, gender, age, phone, email)

	//fmt.Println()
	//向Add()函数传入结构体参数
	if this.customerService.Add(customer) {
		fmt.Println("\n-------------添加客户成功--------------\n\n")
	}else {
		fmt.Println("\n-------------添加客户失败--------------\n\n")
	}
}



//3.删除客户信息

func (this *customerView) delete() {
	fmt.Println("-------------删除客户信息---------\n")
	fmt.Print("请输入要删除客户的Id, <-1>退出：")
	Id := 0
	fmt.Scanln(&Id)
	//若是输入-1，则直接退出删除
	if Id == -1 {
		return
	}

	del := ""
	for {
		fmt.Print("是否真的需要删除该客户信息<y/n>:")
		fmt.Scanln(&del)
		if del == "y" || del == "Y" || del == "n" || del == "N" {
			break
		}else {
			fmt.Println("输入错误，请重新输入！！！")
		}
	}
	//fmt.Println()
	if del == "y" || del == "Y" {
		if this.customerService.Delete(Id) {
			fmt.Println("\n-------------删除成功----------")
		}else {
			fmt.Println("\n删除失败，该Id客户信息不在系统中! ! !\n")
		}
		
	}
}



//4.修改客户信息

func (this *customerView) update() {
	fmt.Print("请输入要修改客户的Id, <-1>放弃修改：")
	id := 0
	fmt.Scanln(&id)
	//判断是否愿意放弃修改，
	//若是输入-1，则直接退出删除
	if id == -1 {
		return
	}
	//FindbyId()函数要是输入的id存在，则返回customer切片相应的下标
	//然后根据这个下标来修改切片中的内容，以达成修改客户信息的内容效果
	//要是id不存在，则返回-1
	index := this.customerService.FindbyId(id)
	for {
		if index == -1 {
			fmt.Print("修改客户的Id不存在客户信息系统！请重新输入！：")
			fmt.Scanln(&id)
			index = this.customerService.FindbyId(id)
		}else {
			break
		}
	}
	fmt.Print("修改后的名字：")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("修改后的性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Print("修改后的年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("修改后的电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("修改后的邮箱：")
	email := ""
	fmt.Scanln(&email)

	//将修改的内容，转换成一个结构体类型输出给customer
	customer := model.NewCustomer(name, gender, age, phone, email)

	//调用Update()函数，对数据内容就行修改
	this.customerService.Update(index, customer)

}



//5.退出方法

func (this *customerView) exit() {
	fmt.Print("是否确定推出客户信息管理系统<y/n>:")
	test := ""
	for {
		fmt.Scanln(&test)
		if test == "y" || test == "Y" || test == "n" || test == "N" {
			break
		}else {
			fmt.Print("输入错误！请重新输入<y/n>:")
		}
	}
	if test == "y" || test == "Y" {
		this.loop = false
	}
}

func (this *customerView) MaxView() {
	for {
		fmt.Println("\n.............客 户 信 息 管 理 软 件...........\n")
		fmt.Println("              1.添 加 客 户\n")
		fmt.Println("              2.删 除 客 户\n")
		fmt.Println("              3.修 改 客 户\n")
		fmt.Println("              4.查 询 客 户\n")
		fmt.Println("              5.退   出\n")
		fmt.Print("请输入功能选项，<1-5>:")
		fmt.Scanln(&this.key)
		fmt.Println()

		switch this.key {
			case "1":
				//fmt.Println("添 加 客 户\n")
				this.add()
			case "2":
				//fmt.Println("删 除 客 户\n")
				this.delete()
			case "3":
				//fmt.Println("修 改 客 户\n")
				this.update()
			case "4":
				//fmt.Println("查 询 客 户\n")
				this.list()
			case "5":
				this.exit()
			default :
			    fmt.Println("输入错误请重新输入！！")
		}
		if !this.loop {
			break
		}
	}
}

func main() {
	customerView := customerView{
		key : "",
		loop : true,
	}
	customerView.customerService = service.NewCustomerService()
	customerView.MaxView()
}