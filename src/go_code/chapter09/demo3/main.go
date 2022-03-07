package main
import (
	"fmt"
	"sort"
)

func main() {

	//先定义一个map
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90

	fmt.Println(map1)
	//现在输出的map1是无序的

	//下面对map1进行排序
	//1.先将map的key放入到切片中
	//2.对切片排序
	//3.遍历切片。然后按照key来输出map的值

	var keys []int
	for k, _ := range map1{
		//通过使用切片append的函数，将得到的数进行添加，得到一组数据
		keys = append(keys, k)
	}

	//sort.Ints()函数，的作用就是对切片的内容进行递增排序
	sort.Ints(keys)
	fmt.Println(keys)

	for _, k := range keys{
		fmt.Printf("map1[%v]=%v \n",k, map1[k])
	}
}