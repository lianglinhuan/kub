package main
import (
	"fmt"
	"go_code/chapter11/encapsulate/model"
)

func main() {
	//将结构体对象进行输出，而且是地址输出，不然改不了原本结构体的参数
	per := model.NewPerson("tom")
	fmt.Println(*per)

	//因为对上面的结构体已经重新定义，且是地址类型的定义
	//（所以使用per结构体可以访问到原结构体的内容）XXXXX这样的理解是错误的！！！

	//工厂模式的本质是通过自定义构造的函数来访问原结构体， 这个才是正解
	//因为model包的set，get方法的首字母是大写的，可以被其他包访问
	//同时，这些方法的是被person结构体类型限制的，NewPerson函数就已经将model包的
	//person结构体类型传递给per变量了，
	//所以，往下的话，就可以直接使用per来使用model包下面的方法了  
	per.SetAge(23)
	per.SetSal(4000.0)
	fmt.Println(*per)

	fmt.Printf("用户的名字为%v,用户的年龄为%v，用户的薪水为%v\n",
	per.Name, per.GetAge(), per.GetSal())
}