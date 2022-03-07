package main
import "fmt"

type MethodUtils struct {

}

//输出9*9乘法表
func (met MethodUtils)Print(n int) {
	for i := 1; i <= n; i++ {
		for j :=1; j <= i; j++ {
			fmt.Printf("%vX%v=%v  ", i, j, i*j)
		}
		fmt.Println()
	}
}

type arr struct {
	arr[3][3] int
}

// func (a arr)Print1(arr[3][3] int) {
// 	for i := 0; i < len(arr); i++ {
// 		for j := 0; j < len(arr[i]); j++{
// 			if i == 1 && j >= 1 {
// 				var n int = 1
// 				arr[i][j] = arr[i][j] + arr[1][n]
// 				arr[1][n] = arr[i][j] - arr[1][n]
// 				arr[i][j] = arr[i][j] - arr[1][n] 
// 				n++
// 				}
// 			if i == 1 && j == 2{
// 				var temp int 
// 				temp = arr[i][j]
// 				arr[i][j] = arr[i+1][j-1]
// 				arr[i+1][j-1] = temp
// 			}
// 			}
// 		}
// 	} 
// }

func (a *arr)Print1() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++{
			if j > i {
				var temp int
				temp = a.arr[i][j]
				a.arr[i][j] = a.arr[j][i]
				a.arr[j][i] = temp
			}
		}
	}
}

func main() {
	// var n int
	// var met MethodUtils
	// fmt.Println("请输入一个数。。。。")
	// fmt.Scanln(&n)
	// met.Print(n)

	var a1 arr
	a1.arr[0][0] = 10
	a1.arr[1][0] = 100

	//给矩阵赋初值
	var n int = 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++{
			a1.arr[i][j] = n + j + 1
		}
		n = n + 3
	}

	//输出初始矩阵
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++{
			fmt.Printf("%v  ", a1.arr[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()

	//调用转置函数
	a1.Print1()
	//（*a1）.Print1()与上面的效果是一样的

	//输出转置后的矩阵
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++{
			fmt.Printf("%v  ", a1.arr[i][j])
		}
		fmt.Println()
	}
	
	// fmt.Println(len(a1.arr))
	// fmt.Println(len(a1.arr[1]))
	// fmt.Printf("[0][0]的地址为%v  [1][0]的地址为%v", &a1.arr[0][0], &a1.arr[1][0])

}