package main
import (
	"runtime"
	"fmt"
)

func main() {
	cpunum := runtime.NumCPU()
	fmt.Printf("cpunum is %v\n", cpunum)

	//设置可以使用多少个cpu，默认全部使用
	runtime.GOMAXPROCS(cpunum - 1)
	fmt.Println("ok")
}