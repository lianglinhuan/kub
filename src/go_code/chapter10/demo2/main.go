package main
import "fmt"

type Persion struct {
	Name string
	Age int
}

func main() {

	var p3 *Persion = new(Persion)
	//因为p3是一个指针，因此标准的给字段赋值方式
	//(*p3).Name = "smith" 也可以这样写字段 p3.Name = "smith"
	//因为：go的设计者，为了程序员方便，底层会对p3.Name进行处理
	//会给p3加上取值运算(*p3).Nmae 实现一样的效果

	(*p3).Name = "smith"
	p3.Name = "marry"

	(*p3).Age = 30
	p3.Age = 100

	fmt.Println(*p3)

	var p4 *Persion = &Persion{}
	//因为peision是一个指针，因此访问字段的方法同时上
	//原理也是一模一样的
	//其中，开括号里面也还是可以对结构体字段赋值
	(*p4).Name = "smith"
	p4.Name = "marry"

	(*p4).Age = 30
	p4.Age = 100

	fmt.Println(*p4)
}