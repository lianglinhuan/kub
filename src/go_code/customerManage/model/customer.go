package model

import (
	"fmt"
)
type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string

}

//使用一个工厂模式，返回customer实例
func NewCustomer(name string, gender string,
	 age int, phone string, email string) Customer {
	return Customer{
		Name : name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Email : email,
	}
}

func (this Customer)Getinfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",this.Id, this.Name,
		 this.Gender, this.Age, this.Phone, this.Email)
	return info
}