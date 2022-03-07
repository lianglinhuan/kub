package model
import (
	"fmt"
)

//构建一个首字母小写的结构体，其他包是不能直接访问的
type person struct{
	Name string
	age int //因为字段是小写，所以即使结果工厂模式后，也不能被其他包直接使用
	sal float64
}

//写一个工厂模式的函数，相当于构造函数
func NewPerson(name string) *person {
	return &person{
		Name : name,
	}
}

func (p *person)SetAge(age int) {
	if age > 0 && age < 150{
		p.age = age
	}else {
		fmt.Println("年龄范围不正确")
	}
}

func (p *person)GetAge() int{
	return p.age
}

func(p *person)SetSal(sal float64){
	if sal >= 3000 && sal <= 30000{
		p.sal = sal
	}else {
		fmt.Println("薪水范围不正确。。。")
	}
}

func (p *person)GetSal() float64{
	return p.sal
}
