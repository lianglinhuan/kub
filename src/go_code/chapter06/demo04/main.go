package main
import (
	"fmt"
	"time"
	"strconv"
)

func test() {

	//对字符串拼接100000次空字符
	str := ""
	for i := 0; i <= 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {

	start := time.Now().Unix()
	test()
	end := time.Now().Unix()
	fmt.Printf("执行test函数耗费时间为%v秒n\n", end - start)
}