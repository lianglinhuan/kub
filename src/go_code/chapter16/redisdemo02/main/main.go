package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	//1.连接redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis Dial err = ", err)
		return
	}
	defer conn.Close()

	//2.通过狗向redis写入数据hash
	_, err = conn.Do("hmset", "user02", "name", "john", "age", 18)
	if err != nil {
		fmt.Println("hmset err = ", err)
		return
	}

	//3.通过go向redis读取数据
	//注意：！！！因为返回有多条数据，所以用Strings()带有s
	r, err := redis.Strings(conn.Do("hmget", "user02", "name", "age"))
	if err != nil {
		fmt.Println("hmget err = ", err)
		return
	}
	//返回多数据，从这里可以看出来，返回的是一个切片类型
	fmt.Printf("r = %v  r type = %T\n", r, r)

	for i, v := range r {
		fmt.Printf("r[%v] = %s\n", i, v)
	}

	fmt.Println("操作ok····")

}