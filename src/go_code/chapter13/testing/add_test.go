package cal

import (
	"testing"
)

func TestAdd(t *testing.T) {
	res := Add(10)
	if res != 55 {
		//fmt.Printf("Add(10)执行错误，期望值=%v 实际值=%v\n", 55, res)

		//t.Fatlf("内容")，该函数的作用是输出“内容”作为日志，并停止该程序
		t.Fatalf("Add(10)执行错误，期望值=%v 实际值=%v\n", 55, res)
	}

	//如果正确，则输出日志
	t.Logf("Add(10)执行正确。。。。")
}

// func TestSub(t *testing.T) {
// 	res := Sub(10, 3)
// 	if res != 7 {
// 		t.Fatalf("Sub(10, 3)执行错误，期望值=%v 实际值=%v\n", 7, res)
// 	}
// 	t.Logf("Sub(10, 3)执行正确。。。。")
// }