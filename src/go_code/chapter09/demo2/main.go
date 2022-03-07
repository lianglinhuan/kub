package main
import (
	"fmt"
)

func main() {
	//演示map切片的使用
	//1.先声明一个map切片
	var monsters []map[string]string
	//准备放入两个妖怪
	monsters = make([]map[string]string, 2)

	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"  
		monsters[0]["age"] = "500"  
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "玉兔"  
		monsters[1]["age"] = "300"  
	}
	
	// //因为原本切片定义的容量就是2，再加的话就会报错
	// //因此需要在原本定量好的基础上再添加就得使用切片的append
	// if monsters[2] == nil {
	// 	monsters[2] = make(map[string]string, 2)
	// 	monsters[2]["name"] = "狐狸"  
	// 	monsters[2]["age"] = "500"  
	// 

	//使用切片的append函数，动态增加monster
	newMonster := map[string]string{
		"name" : "新妖怪",
		"age" : "200",
	}
	monsters = append(monsters, newMonster)

	
	fmt.Println(monsters)
	
}