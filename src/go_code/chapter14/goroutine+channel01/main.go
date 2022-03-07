package main
import (
	"fmt"
	"time"
)

func WriteChan(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan<- i
		fmt.Println("WriteChan is ", i)
		//添加一个个延时是为了更加直观得观察WriteChan线程和ReadChan线程
		//之间是交替运行的
		time.Sleep(time.Second)
	}
	//等到写入操作完成后，关闭管道
	close(intChan)
}

func ReadChan(intChan chan int, exitChan chan bool) {
	for {
		//监控是否已经读取完成管道的信息
		//若是读取到管道最后一个值，ok接收到一个false
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("ReadChan 读到的数据=%v\n",v)
		//延时，作用与WriteChan线程的效果一样
		time.Sleep(time.Second)
	}
	//当管道读取完成后，将一个true传递到exitChan管道中
	//exitChan也是需要关闭的
	exitChan<- true
	close(exitChan)
}

func main() {
	//设定两个管道
	//intChan管道是用来写入数据和读取数据逇
	//exitChan管道是用来标定退出主线程的
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go WriteChan(intChan)
	//若是注销掉读协程操作，即管道没有进行读取操作，
	//当管道写入的数据超过了管道的容量后,管道会出现阻塞现象，报deadlock错误
	//  intChan := make(chan int, 10)
	//  //go ReadChan(intChan, exitChan)
	go ReadChan(intChan, exitChan)

	//在主线程中使用死循环检测协程是否结束
	//因为exitChan管道只有在读完intChan管道数后才向其传递以个值
	//所以在没有关闭exitChan管道的时候，一致读取管道都是有效的
	//即ok一直为true
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

}