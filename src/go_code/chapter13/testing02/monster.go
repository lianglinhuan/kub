package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Monster struct{
	Name string
	Age int
	Skill string
}


//将结构体数据的内容序列化
//然后将序列化的数据传送到文件中
func (this *Monster)Store() bool{
	
	file := "f:/testMonster.txt"
	//1.先序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
		return false
	}

	//fmt.Printf("monster序列化后的结果=%v\n", string(data))

	//2.ioutil.WriteFile()一次性写入文件
	err = ioutil.WriteFile(file, data, 0666)
	if err != nil {
		fmt.Printf("write file error= %v\n", err)
		return false
	}
	return true
}


//将序列化的数据从文件中读取出来
//然后在将序列化的数据进行反序列化
func (this *Monster)ReStore() bool {

	file := "f:/testMonster.txt"

	//1.先从文件中，读取序列化的字符串
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file error= %v\n", err)
		return false
	}

	//2.将数据进行反序列化
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Printf("反序列化错误 err=%v\n", err)
		return false
	}
	return true
}