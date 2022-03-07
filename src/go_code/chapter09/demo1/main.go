package main

import (
	"fmt"
)

func main() {

	//第一种方法
	var a map[string]string
	//在使用map前，需要先make，make的作用就是给map分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no1"] = "武松"
	a["no3"] = "吴用"
	//在map中键是惟一的，新建会替换老键，而值却可以有相同的
	fmt.Println(a)

	
	//第二中方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	//在map中因为没有下标索引，所以只能通过for-range来遍历map
	for k, v := range cities{
		fmt.Printf("k=%v  v=%v\t", k, v)
	}
	fmt.Println()

	//第三种方式
	heroes := map[string]string{
		"hero1" : "宋江",
		"hero2" : "吴用",
		//在没定义一行后一定要加个逗号，分隔，不然编译会报错
	}
	//第三种方式定义完之后还是可以在后面添加
	heroes["hero3"] = "林冲"
	fmt.Println(heroes)


	//课堂练习，输出三个学生的信息，信息包含name,sex等信息
	studentMap := make(map[string]map[string]string)
	//因为上面已经定义过了，所以直接用make赋值就行了而不需要使用:=
	studentMap["no1"] = make(map[string]string)
	studentMap["no1"]["name"] = "宋江"
	studentMap["no1"]["sex"] = "男"

	studentMap["no2"] = make(map[string]string)
	studentMap["no2"]["name"] = "武媚娘"
	studentMap["no2"]["sex"] = "女"

	//在map中因为没有下标索引，所以只能通过for-range来遍历map
	//因为使用两个map所以for也要使用两次遍历
	for k1, v1 := range studentMap{
		fmt.Println(k1)
		for k2, v2 := range v1{
			fmt.Printf("\tk2=%v  v2=%v", k2, v2)
		}
	}

	fmt.Println()
	fmt.Println(studentMap["no1"])
	fmt.Println(studentMap["no2"])
	//还可以更加灵活
	fmt.Println(studentMap["no2"]["name"])
}