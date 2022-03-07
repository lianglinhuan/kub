package cal

import (
	"testing"
)

func TestSub(t *testing.T) {
	res := Sub(10, 3)
	if res != 7 {
		t.Fatalf("Sub(10, 3)执行错误，期望值=%v 实际值=%v\n", 7, res)
	}
	t.Logf("Sub(10, 3)执行正确。。。。~~")
}