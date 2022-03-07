package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	//通过go向redis写入数据和读取数据
	//1.连接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	//若连接redis错误则退出
	if err != nil {
		fmt.Println("redis.Dial err = ", err)
		return
	}
	//注意！！！记得及时关闭服务器接口，
	defer conn.Close()

	//2.通过go向redis写入数据string
	_, err = conn.Do("set", "name", "tomjerry")
	if err != nil {
		fmt.Println("set err = ", err)
		return
	}

	//3.通过go向redis读取数据string
	//因为conn.Do()函数返回的值(r)是一个空接口类型的数据,
	//所以要向在终端打印出获取到key的值的话，就需要进行类型转换
	//注意：使用普通的类型断言会出现错误，如下
	//nameString := r.(String)
	//而是要使用redis包自带的类型转换函数进行转换，redis.String()
	r, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("set err = ", err)
		return
	}
	fmt.Println(r)

	fmt.Println("操作成功")
}