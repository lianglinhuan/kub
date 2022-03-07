package main
import "fmt"

func main() {

	//还有97天放假，问：还有xx个星期零xx天
	var days int 
	days = 97
	fmt.Printf("还有97天放假也就是%v个星期%v天\n", days / 7, days % 7)

	//将华氏温度转化为摄氏温度
	var temp int
	var wendu fl
	//一定要使用5.0不然要是直接用5,5/9恒等于0.导致输出的结果也为0
	//公式
	wendu = 5.0/9*(float32(temp) - 100)  
	fmt.Printf("华氏温度为%v的设施温度是%v", temp, wendu)
}