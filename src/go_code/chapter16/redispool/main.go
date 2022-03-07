package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool {
		MaxIdle : 8, //最大空闲数链接数
		MaxActive : 0, //表示和数据库的最大链接数，0表示没有限制
		IdleTimeout : 300,//最大空闲时间
		Dial : func() (redis.Conn, error) {//初始化链接的代码，链接哪个ip的redis
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	//先从pool取出一个链接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("set", "name", "Tom")
	if err != nil {
		fmt.Println("conn.Do set err = ", err)
		return
	}

	r, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("conn.Do get err = ", err)
		return
	}
	fmt.Println("r = ", r)

	//如果我们要从pool取出链接，一定要保证连接池是没有关闭的
	//pool.Close()//是用来关闭连接池的
	conn2 := pool.Get()
	//要是关闭了链接池，以下的操作就会报错，
	//注意！！！conn2 := pool.Get()这个操作还是可以进行的，但是没有了意义
	_, err = conn2.Do("set", "name2", "Tom~~~~")
	if err != nil {
		fmt.Println("conn2.Do set err = ", err)
		return
	}

	r2, err := redis.String(conn2.Do("get", "name2"))
	if err != nil {
		fmt.Println("con2n.Do get err = ", err)
		return
	}
	fmt.Println("r2 = ", r2)
}