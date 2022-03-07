package main
import (
	"fmt"
)

func fbn(n int) ([]uint64){
	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i - 1] + fbnSlice[i - 2]
	}
	return fbnSlice
}

func main() {

	//延时切片的使用 make
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[3] = 20
	//对于切片，必须make使用
	fmt.Println(slice)
	fmt.Println("slice的size=", len(slice))
	fmt.Println("slice的size=", cap(slice))

	//方式3
	fmt.Println()
	//第三种方式：定义一个切片，直接就指定具体数组，使用原理类似make
	var strslice []string = []string{"tom", "jack", "mary"}
	fmt.Println("strslice=", strslice)
	fmt.Println("strslice size=", len(strslice))
	fmt.Println("strslice size=", cap(strslice))


	//使用append内置函数。可以对切片进行动态追加
	var slice3 []int = []int{100, 200, 300}
	//通过append直接给slice3追加具体的元素
	slice3 = append(slice3, 400, 500, 600)
	fmt.Println("slice3", slice3) 

	//通过append将切片slice3追加给slice3,后面追加的只能是切片
	//不能是数组，其中也要注意切片名后面必须加上(...)
	slice3 = append(slice3, slice3...)
	fmt.Println("slice3", slice3) 


	//string底层是一个byte数组，因此string也可以进行切片处理
	str := "hello@atguigu"
	//使用切片获取atguigu
	slice4 := str[6:]
	fmt.Println("slice4=", slice4)


	//string是不可变的，也就是说不能通过str[0] = 'z'方式来修改字符串
	
	//如果需要修改字符串，可以先将string -> []byte 或者 []rune -> 修改 -> 重新转成string
	//"hell0@atguigu" => "zello@atguigu"
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("str=", str)

	//细节，转成[]byte后。可以处理英文和数字，但是不能处理中文
	//原因是[]byte字节来处理，而一个汉子，是三个字节，因此就会出现乱码
	//解决方法是将string转成[]rune即可，因为[]rune是按字节来处理，兼容汉字

	arr2 := []rune(str)
	arr2[0] = '北'
	str = string(arr2)
	fmt.Println("str=", str)

	fbn := fbn(10)
	fmt.Println("fbn=", fbn)
}