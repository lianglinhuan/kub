package main
import "fmt"

//定义全局变量
var n6 = 100
var n7 = 300
var name2 = "jack"
//上面的声明方式，也可以改成一次性声明
var (
	n8 = 400
	n9 = 500
	name3 = "mary"
)

func main() {
	//golang的变量使用方法1
	//第一种：指定变量类型。声明后若不赋值。则使用的是默认值
	//int的默认值是0,
	var i int
	fmt.Println("i=", i)

	//第二种：根据值自行判断变量的类型（类型推导）
	var num = 10.11
	fmt.Println("num=", num)

	//第三种：省略var，注意 :=左侧的变量不应该是声明过的，否则导致编译错误
	//下面的方式等价 var name string name = "tom"
	// := 的 :不能省略。否则错误
	name := "tom"
	fmt.Println("name=", name)

	//golang如何一次行声明多个变量
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	//一次性声明多个不同类型的变量
	//变量定义不能同名，否则报错
	n4, name1, n5 := 100, "tom", 888
	fmt.Println("n4=", n4, "name1=", name1, "n5=", n5)
	fmt.Println("n6=", n6, "name2=", name2, "n7=", n7)
	fmt.Println("n8=", n8, "name3=", name3, "n9=", n9)
}