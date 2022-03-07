package main
 
import "fmt"

//演示goland中基本数据联系转成string使用
func main () {

	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var mychar byte = 'h'
	var str string

	//使用第一中方方法转换fmt.Sprintf()
	//其中输出%q可以换成%q，%q会在输出的结果上加上双引号
	//输出的格式%q，%T，%v等可以在官方文档哪里查看使用

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str type %T str=%q\n", str, str)

	//第二种方式，使用strconv函数	
	
}