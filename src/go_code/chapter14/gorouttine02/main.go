package main
import (
	"fmt"
	"time"
	"sync"
)

//myMap := make(map[int]int) 在函数体外不能这样定义一个变量，需要使用var模式，例如下面的例子
var (
	myMap = make(map[int]int)

	//重新定义一个sync包下面的Mutex结构体的数据类型变量
	lock sync.Mutex
)

// func testMap(n int) {
// 	res := 1
// 	for i := 0; i <= n; i++ {
// 		res *= 1
// 	}
// 	myMap[n] = res
// }

//出现的问题，因为协程的并发关系，多个数据同时对map空间里面写入数据，造成资源竞争的问题
//而导致程序无法正常运行。

//1.问题解决方法一：在写入操作加上互斥锁
func testMap(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//加锁
	lock.Lock()
	myMap[n] = res
	//解锁
	lock.Unlock()
}

func main() {
	for i := 0; i <= 20; i++{
		//使用协程，go
		go testMap(i)
	}
	time.Sleep(time.Second * 5)

	//按理说10秒协程都应该执行完了，后面就不应该出现资源竞争的问题，但是实际运行中，还是可能在
	//读取数据这块出现（运行是增加 -race参数确实会发现有资源竞争问题），这是因为从程序设计上可以知道
	//5秒就执行完所有协程，但是主线程并不知道，因此底层仍然会出现资源争夺，所以在读取数据的时候也是
	//需要加上互斥锁，防止资源竞争。
	lock.Lock()
	for i, v := range myMap{
		fmt.Printf("myMap[%v] = %v\n", i, v)
	}
	lock.Unlock()
}