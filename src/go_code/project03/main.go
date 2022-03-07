package main
import (
	"fmt"
	"container/list"
)
func main() {
	//New函数初始化链表，返回的是指针类型
	//即tmplist是指针类型
	tmplist := list.New()
	
	for i := 1; i <= 10; i++ {
		tmplist.PushBack(i)
	}
	//链表的每一次插入操作都会返回一个*list.Element结构
	//即first就是该类型的数
	//其实，如果对链表中的成员进行删除、移动、或者指定插入操作的话都是通过插入返回的*list.Element该类型实现的
     first := tmplist.PushFront(0)
	 tmplist.Remove(first)

	 for list1 := tmplist.Front(); list1 != nil; list1 = list1.Next() {
		 //Front()函数的作用就是返回该链表的头元素
		 //Next()函数的作用就是返回当前链表的下一个元素
		 fmt.Print(list1.Value, " ")
	 }
}
