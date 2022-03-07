package service

import (
	_"fmt"
	"go_code/customerManage/model"
)

type CustomerService struct {
	//完成对Customer的操作，增删改查
	customers []model.Customer
	//声明一个字段，表示当前切片含有多少个客户
	//该字段后面+1，表示新用户的id
	CustomerNum int
}

func NewCustomerService() *CustomerService{
	CustomerService := &CustomerService{}
	//初始化测试的内容而已，
	CustomerService.CustomerNum = 1
	customer := model.NewCustomer( "张三", "男", 20, "1511", "168@")
	customer.Id = CustomerService.CustomerNum
	CustomerService.customers = append(CustomerService.customers, customer)
	return CustomerService
}

//返回一个切片
//可是，为什么不直接用上面哪个方法返回的类型来定位使用切片呢？还要特地拿出来单独实现？
//只是为了方便往下的操作吗？应该是的, 没错就是这样子
func (this *CustomerService) List() []model.Customer{
	return this.customers
}

//向切片添加数据，数据类型是一个结构体类型
func (this *CustomerService) Add(customer model.Customer) bool {
	//通过一个指针保存之前的的数据，可以保证编号是按顺序排序的
	this.CustomerNum++
	customer.Id = this.CustomerNum
	this.customers = append(this.customers, customer)
	return true
}

//定义一个方法来检测，所要删除的Id的客户信息是否存在
func (this *CustomerService) FindbyId(Id int) int{
	//要是没有输入Id参数是，直接回车默认的是-1
	index := -1
	for i := 0; i < this.CustomerNum; i++ {
		//这里的this.customer[i].Id的Id不是传入的Id参数，而是原先结构体定义的一个字段
		//在这里吃了大亏！！！
		if this.customers[i].Id == Id {
			index = i
		}
	}
	return index
}

//定义一个删除的方法
func (this *CustomerService) Delete(Id int) bool {
	index := this.FindbyId(Id)
	if index == -1 {
		return false
	}
	//在切片中[:index]，表示从0到index-1的所有内容，
	//[index+1:]则表示从index+1到最后一个，两者并接在一起刚好省略了index的内容
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
	
}

//定义一个修改客户信息的方法
func (this *CustomerService) Update(id int, customer model.Customer) {
	this.customers[id].Name = customer.Name
	this.customers[id].Gender = customer.Gender
	this.customers[id].Age = customer.Age
	this.customers[id].Phone = customer.Phone
	this.customers[id].Email = customer.Email

}