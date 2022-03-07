package main
import (
	"fmt"
	"sort"
	"math/rand"
)

/*
实现对Hero结构体切片的排序：sort.Sort(data Interface)
*/


//1.定义一个结构体
type Hero struct {
	Name string
	Age int
}

//2定义一个结构体切片类型
type HeroSlice []Hero

//3.实现Interface接口
//实现接口需要定义Len() int, Less(i, j int) bool和Swap(i, j int)这三个方法
//实现上面所说的方法，就已经定义好了Interface接口，这个接口只要你满足实现的三个方法，它就会自动存在
func (hs HeroSlice) Len() int{
	return len(hs)
}

//Less方法就是决定你用什么标准进行排序
//关键点！！！
func (hs HeroSlice) Less(i, j int) bool{
	return hs[i].Age < hs[j].Age
}

//Swap实现交换
func (hs HeroSlice) Swap(i, j int){
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp

	//上面的三句话等价于下面的一句话，注意了，开发可能常用到
	//多重赋值
	//hs[i], hs[j] = hs[j], hs[i]
}

//1.定义一个结构体
type Student struct {
	Name string
	Age int
	Score int
}

//2.声明一个切片类型的结构体
type StuSlice []Student

//3.实现Interface接口需要方法的声明，实现
func (stu StuSlice)Len() int{
	return len(stu)
}

func (stu StuSlice)Less(i, j int) bool {
	return stu[i].Score < stu[j].Score
}

func (stu StuSlice)Swap(i,j int) {
	temp := stu[i]
	stu[i] = stu[j]
	stu[j] = temp
}



func main() {

	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name : fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age : rand.Intn(100),
		}
		//每次增加的信息。追加到heroes切片上面
		heroes = append(heroes, hero)
	}

	//排序前的顺序
	for _, val := range heroes{
		fmt.Println(val)
	}

	sort.Sort(heroes)
	fmt.Println(".....排序后......")

	//排序后的顺序
	for _, val := range heroes {
		fmt.Println(val)
	}

	fmt.Println()
	fmt.Println()


	//对学会成绩进行排序

	var student StuSlice
	for i := 0; i < 10; i++ {
		stu := Student{
			Name : fmt.Sprintf("学会%d", rand.Intn(100)),
			Age : rand.Intn(25),
			Score : rand.Intn(100),
		}

		student = append(student, stu)
	}

	fmt.Println("。。。排序前。。。")
	for _, val := range student{
		fmt.Println(val)
	}

	
	sort.Sort(student)
	fmt.Println("。。。排序后。。。")

	for _, val := range student{
		fmt.Println(val)
	}
}

