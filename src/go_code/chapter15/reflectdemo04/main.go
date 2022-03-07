package main
import (
	"fmt"
	"reflect"
)
//定义了一个方法
type Monster struct {
	Name string `json:"name"`
	Age int `json:"monster_age"`
	Sorce float32
	Sex string
}
//方法，显示s的值
func (s Monster)Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}
//方法，返回两个数的和
func (s Monster)GetSum(n1, n2 int) int {
	return n1 + n2
}
//方法，接收4个值赋给Monster
func (s Monster)Set(name string, age int, sorce float32, sex string) {
	s.Name = name
	s.Age = age
	s.Sorce = sorce
	s.Sex = sex
}

func TestStruct(a interface{}) {
	//获取reflect.Type类型
	typ := reflect.TypeOf(a)
	//获取reflect.Value类型
	val := reflect.ValueOf(a)
	//获取a对应的类别
	kd := val.Kind()
	//如果传入的不是struct，就退出函数
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	//获取到该结构体有几个字段
	num := val.NumField()
	fmt.Printf("struct has %v fields\n", num)

	//遍历结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %v: 值为=%v\n", i, val.Field(i))
		//获取struct标签，注意需要通过reflect.Type来回去tag标签的值
		//reflect.Type下面的Field()返回的就是一个StructField结构体类型
		//结构体类型下面就偶遇个Tag字段，该字段也是一个StructTag结构体类型
		//在StructTag结构体类型的方法中就有一个Get()方法同来获取结构体中的tag
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %v: tag为=%v\n", i, tagVal)
		}
	}

	//获取到结构体有多少个方法,总共三个
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methosds\n", numOfMethod)

	//获取到第二个方法，而且给方法传入一个空值
	//方法是按照方法名的字母（ASSIC码）进行排序的
	val.Method(1).Call(nil)

	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)
	fmt.Println("res=", res[0].Int())
}

func main() {
	///创建一个Monster实例
	var a Monster = Monster{
		Name : "黄鼠狼精",
		Age : 400,
		Sorce : 30.8,
	}
	TestStruct(a)
}