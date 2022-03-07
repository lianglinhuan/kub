package main

// import "fmt"
// import "unsafe"
import(
	"fmt"
	"unsafe"
)

func main() {
	var n1 = 100 //n1是什么类型
	//使用fmt.Printf()函数，格式化输出变量的类型
	fmt.Printf("n1 的 类型为 %T ", n1)
	fmt.Printf("\n")

	//如何在程序查看某个变量占用字节大小和数据类型（使用较多）
	var n2 int64 = 10
	//unsafe.Sizeof(n2)是unsafe包的一个函数，可以返回n2变量占用的字节大小
	fmt.Printf("n2 的类型为 %T   n2占用的字节数是 %d ", n2, unsafe.Sizeof(n2))
}