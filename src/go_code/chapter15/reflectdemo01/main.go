package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {

	//通过反射获取的传入的变量的type，kind，值
	//1.先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp = ", rTyp)

	//2.获取到reflect.ValueOf
	rVal := reflect.ValueOf(b)
	//这时的rVal返回的值不在是int类型，而是reflect.Value类型
	fmt.Printf("rVal = %v \n rVal 的类型为%T\n", rVal, rVal)

	//3.获取变量对应的kind，有两种方法
	//(1)rVal.Kind()
	//(2)rTyp.Kind()
	//两者之间的效果一样，本质也是一样的
	fmt.Printf("kind = %v  kind = %v\n", rVal.Kind(), rTyp.Kind())

	n2 := 2 + rVal.Int()
	fmt.Println("n2 = ", n2)

	//下面我们将rVal转成interface{}
	iv := rVal.Interface()
	//将interface{}通过断言转成需要的类型
	num2 := iv.(int)
	fmt.Printf("num2 = %v  num2 = %T\n\n\n", num2, num2)
}

func reflectTest02(b interface{}) {
	//通过反射获取的传入的变量的type，kind，值
	//1.先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	//该输出是main.Student
	fmt.Println("rTyp = ", rTyp)

	//2.获取到reflect.ValueOf
	rVal := reflect.ValueOf(b)
	//这时的rVal返回的值不在是int类型，而是reflect.Value类型
	fmt.Printf("rVal = %v \n rVal 的类型为%T\n", rVal, rVal)

	//3.获取变量对应的kind，有两种方法
	//(1)rVal.Kind()
	//(2)rTyp.Kind()
	//两者之间的效果一样，本质也是一样的
	fmt.Printf("kind = %v  kind = %v\n", rVal.Kind(), rTyp.Kind())

	//下面我们将rVal转成interface{}
	iv := rVal.Interface()
	//结构体转换为接口类型后，其类型仍然为main.Student
	fmt.Printf("iv = %v  iv = %T\n", iv, iv)
	//将interface{}通过断言转成需要的类型,要是不适用断言则不能够使用结构体内的数据
	//这里使用简单的一带检测的类型断言，可以使用switch的断言形式来做得更加灵活
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("stu.Name = %v\n", stu.Name)
	}
}

type Student struct {
	Name string
	Age int
}

func main() {

	//延时对基本数据类型、interface{}、reflect.Value进行反射的基本操作
	var num int = 100
	reflectTest01(num)

	stu := Student{
		Name : "tom",
		Age : 18,
	}
	reflectTest02(stu)
}