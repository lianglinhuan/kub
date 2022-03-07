package main
import "fmt"

func main() {

	//演示 % 的使用
	//公式 a % b = a - a/b*b 
	fmt.Println("10%3=", 10 % 3)    //1
	fmt.Println("-10%3=", -10 % 3)  //-1
	fmt.Println("10%-3=", 10 % -3)  //1
	fmt.Println("-10%-3=", -10 % -3)//-1
}