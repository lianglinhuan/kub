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
}

//反序列化结构体
func unmarshalstruct() {
	str := "{\"Name\":\"牛魔王\",\"Age\":189,\"Sal\":6669.8,\"Skill\":\"牛角冲击\"}"

	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err = %v\n", err)
	}
	fmt.Printf("unmarshal struct = %v\n", monster)
}

//对map进行序列化
func testMap() string {
	//定义一个map
	var a map[string]interface{}

	//使用map，需要make才行
	a = make(map[string]interface{})
	a["name"] = "红孩儿~~~~"
	a["age"] = 230
	a["address"] = "洪崖洞"

	data, err := json.Marshal(&a)
	if err != nil {
		fmt.Printf("序列化错误，err=%v\n", err)
	}

	//fmt.Printf("a map序列化后的结构为=%v\n", string(data))
	return string(data)
}

//反序列化map
func unmarshalMap() {
	//str := "{\"address\":\"洪崖洞\",\"age\":230,\"name\":\"红孩儿\"}"

	str := testMap()
	var a map[string]interface{}

	//在反序列过程中，map类型的使用时不需要使用make
	//因为make操作被封装在了,unmarshal()函数里面了
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err = %v\n", err)
	}
	fmt.Printf("unmarshal map = %v\n", a)
}

//反序列化slice
func unmarshalslice() {

	//要是字符串过长，则可以使用"+"解决，即字符串的拼接
	str := "[{\"address\":\"北京\",\"age\":17,\"name\":\"jack\"}," +
	"{\"address\":[\"上海\",\"成都\"],\"age\":18,\"name\":\"mary\"}]"

	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err = %v\n", err)
	}
	fmt.Printf("unmarshal map = %v\n", slice)
}

func main() {
	unmarshalstruct()
	unmarshalMap()
	unmarshalslice()
}