package main
import "fmt"
import "strings"

func getSun(n1 int, n2 int) int{
	return n1 + n2
}

//函数也是一个数据类型
func myFun(funvar func(int, int) int, num1 int, num2 int) int{
	return funvar(num1 , num2)
}

//该函数可以定义多个参数，方便函数的调用
//而且可变参数一定要放在确定参数的后面
func sum(n1 int, args... int) int{
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i] //args[0]表示取出agrs切片的第一个元素值，以此类推
	}
	return sum
}
//这个函数可以交换n1,与n2的值，这个不是题目要求的吧？
// func swap(n1 *int, n2 *int) (int, int){
// 	var n3 *int
// 	n3 = n1
// 	fmt.Printf("n3=%v\n", *n3)
// 	n1 = n2
// 	fmt.Printf("n1=%v\n", *n1)
// 	n2 = n3
// 	fmt.Printf("n2=%v\n", *n2)

// 	return *n1, *n2
// }

//下面这个函数才是真正实现题目要求的功能
//上面的那个函数只是进行简单的数值交换，然后返回，并没有体现出指针的特性
func swap(n1 *int, n2 *int) {
	//定义一个临时变量
	t := *n1
	*n1 = *n2
	*n2 = t
}

//编写一个函数makeSuffix(suffix string) 可任意接受一个文件后缀名（比如.jpg）,并返回一个闭包
//调用壁报，可以传染源一个文件名，如果该文件没有只当的后缀（比如.jpg），则在其后面加上后缀
//要求使用闭包的方式完成
//strings.HasSuffix，该函数可以判断某个字符串是否有只当的后缀
//strings.HasSuffix函数在strings包内，需要引用该包
func makeSuffix(suffix string) func (string) string {

	return func (name string) string {
		//如果name没有置地刚后缀，则加上，否则就返回原来的名字
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {

	res := myFun(getSun, 50, 60)
	fmt.Printf("res=%v\n", res)

	res2 := sum(10, 0, 90, 10, -1)
	fmt.Printf("res=%v\n", res2)
	
	var n1 int = 10
	n2 := 20
	fmt.Printf("交换前的n1=%v,n2=%v\n", n1, n2)
	// n1, n2 = swap(&n1, &n2)
	swap(&n1, &n2)
	fmt.Printf("交换后的n1=%v,n2=%v\n", n1, n2)

	//测试makeSuffix的使用
	//返回一个闭包
	f2 := makeSuffix(".jpg")
	fmt.Printf("文件名处理后=%v\n", f2("winter"))
	fmt.Printf("文件名处理后=%v\n", f2("bird.jpg"))
}