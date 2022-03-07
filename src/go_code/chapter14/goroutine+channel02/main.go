package main

import (
	"fmt"
	"time"
)

//要求统计1~2000的素数，使用goroutine+channel来完成

//该方法向管道输入数据
func PushNum(intChan chan int) {
	for i := 1; i <= 80000; i++ {
		intChan<- i
	}
	//关闭intChan
	close(intChan)
}

//PrimeNum()方法，将管道intChan中的数据输出找出素数，然后将素数输入到primeChan管道中
//等到intChan管道中的数据输出完，则向exitChan管道输入true
func PrimeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for {
		//延时10毫秒，是为了更加方便观察程序的运行输出
		//不然，fmt.Printf("一个协程结束\n")因为Printf()函数的打印时间与处理时间不一致导致的
		//time.Sleep(time.Millisecond *10)
		flag := true
		num, ok := <-intChan
		if !ok{
			break
		}
		//判断num是否为素数
		for i := 2; i < num; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}
		//若num是素数，则将num输入primeChan管道
		if flag {
			primeChan<- num
		}
	}
	//每一个协程跑完，都需要向exitChan管道传入一个true
	exitChan<- true
	fmt.Printf("一个协程结束\n")
}


func main() {
	intChan := make(chan int, 200)
	primeChan := make(chan int, 20000)
	exitChan := make(chan bool, 4)

	//获取当前时间
	start := time.Now().Unix()

	go PushNum(intChan)

	//循环，生成4个协程，因为我的cpu核数为4个，所以就让它跑4个
	for i := 1; i <= 4; i++ {
		go PrimeNum(intChan, primeChan, exitChan)
	}

	//此处也是一个协程，该协程的作用是等待所有的PrimeNum协程结束
	//然后关闭primeNum管道，
	//要是不关闭管道的话，在读取primeNum管道时就会报deadlock错误
	go func() {
		for i := 1; i <= 4; i++ {
			<-exitChan
		}
		end := time.Now().Unix()
		fmt.Println("使用协程耗时=", end - start)
		close(primeChan)
	}()

	//主线程循环输出primeChan管道数据
	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
		//fmt.Printf("素数 = %v\n", num)
	}
	fmt.Printf("主线程退出")

}