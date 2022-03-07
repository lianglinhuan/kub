package main

import (
	"fmt"
	"reflect"
)
func reflectTest01(b interface{}) {
	//2.获取到reflect.ValueOf
	rVal := reflect.ValueOf(b)
	//3.获取变量对应的kind，为ptr
	fmt.Printf("kind = %v \n", rVal.Kind())
	//4.rVal.SetInt(num)方法是可以对rVal修改指定的num值
	//又因为SetInt()方法规定是value类型使用的，而rVal.Kind()得出，rVal是一个指针类型
	//所以不能直接使用rVal.SetInt(num)，而是需要将指针转换会value类型，
	//rval.Elem()刚好可以满足这种需求
	rVal.Elem().SetInt(20)
}

func main() {

	//延时对基本数据类型、interface{}、reflect.Value进行反射的基本操作
	var num int = 100
	fmt.Println("前num = ", num)
	//因为要修改num值，所以只能传入指针
	reflectTest01(&num)
	fmt.Println("后num = ", num)
}