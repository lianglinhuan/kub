package main
import (
	"fmt"
)

type Student struct{
	name string
	gender string
	age int
	id int
	score float64
}

func (student *Student) say() string{
	infoStr := fmt.Sprintf("student的信息 name=[%v], gender=[%v], age=[%v], id=[%v], score=[%v]", student.name, student.gender,
		student.age, student.id, student.score)

		return infoStr
}

//面向对象写一个箱子体积的方法实例
type Box struct{
	len float64
	width float64
	height float64
}

func ( box *Box) getVolumn() float64{
	return box.len * box.width * box.height
}

//面向对象编写一个根据年龄，来确定门票的方法实例
type Vistor struct {
	name string
	age int
}

func (vir *Vistor)needMonery() {
	if vir.age > 18 {
		fmt.Printf("%s年龄为%v,门票价格为：20元", vir.name, vir.age)

	} else{
		fmt.Printf("%s年龄为%v,门票免费", vir.name, vir.age)
	}
}

func main() {

	var stu = Student{
		name : "tom",
		gender : "male",
		age : 18,
		id : 1000,
		score : 99.88,
	}
	fmt.Println(stu.say())


	var box Box
	box.len = 1.1
	box.width = 2.2
	box.height = 3.3
	fmt.Println(box.getVolumn())
	//保留两位小数
	fmt.Printf("%.2f", box.getVolumn())
	fmt.Println()

	
	var vir Vistor、
	//建立一个死循环，检测输入的信息
	for {
	fmt.Print("请输入姓名：")
	//键盘输入的string和int类型一般使用的是Scanfln（）函数，而不是Scanf（）函数
	fmt.Scanln(&vir.name)
	//当输入的名字为n时，退出循环
	if vir.name == "n" {
		break
	}
	fmt.Print("请输入年龄：")
	fmt.Scanln(&vir.age)
	
	vir.needMonery()
	fmt.Println()
	fmt.Println()
	}
}