package main

import (
	"fmt"
	"encoding/json"
)
//定义一个结构体
type Monster struct {
	Name string
	Age int
	Sal float64 
	Skill string 
	// //使用反引号，做标签，使得序列化后的字段改为标签的内容
	// Sal float64 `json:"monster_sal"` //通过反射机制实现
	// Skill string `json:"monster_skill"`
}

//定义一个方法，对结构体数据进行json序列化
func testStruct() {
	monster := Monster{
		Name : "牛魔王",
		Age : 189,
		Sal : 6669.8,
		Skill : "牛角冲击",
	}

	//json.Marshal()函数返回的数据的[]byte的切片类型
	//要是向正确输出json字符串的内容还需要，对json数据进行转换
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}

	fmt.Printf("monster序列化后的结果=%v\n", string(data))
}

//对map进行序列化
func testMap() {
	//定义一个map
	var a map[string]interface{}

	//使用map，需要make才行
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 230
	a["address"] = "洪崖洞"

	data, err := json.Marshal(&a)
	if err != nil {
		fmt.Printf("序列化错误，err=%v\n", err)
	}

	fmt.Printf("a map序列化后的结构为=%v\n", string(data))
}

//对切片进行序列化,切片的类型为map[string]interface{}
func testSilce() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}

	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 17
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "mary"
	m2["age"] = 18
	//使用数组定力类型，可以添加多个值
	m2["address"] = [2]string{"上海", "成都"}
	slice = append(slice, m2)

	data, err :=json.Marshal(slice)
	if err != nil {
		fmt.Printf("slice序列化错误，err=%v\n", err)
	}

	fmt.Printf("slice序列化后的结果=%v\n", string(data))
}

func main(){
	testStruct()
	testMap()
	testSilce()

}