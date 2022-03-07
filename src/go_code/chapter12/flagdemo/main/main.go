package main 

import (
	"fmt"
	"flag"
)

type NumChar struct {
	user string
	pwd string
	host string
	port int
}

func main() {

	var numch NumChar

	//flag.StringVar(接收参数的变量, 传递参数, 传递参数相对应的值, 解释 )
	//传递参数写u，在命令传递参数是是需要使用-u的，多了一个-
	flag.StringVar(&numch.user, "u", "", "用户名，默认为空")
	flag.StringVar(&numch.pwd, "p", "", "密码，默认为空")
	flag.StringVar(&numch.host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&numch.port, "P", 3306, "端口号，默认为3306")

	//定义好flag后，必须使用该操作
	flag.Parse()

	fmt.Printf("user=%v\n pwd=%v\n host=%v\n port=%v\n",
	numch.user, numch.pwd, numch.host, numch.port)
}