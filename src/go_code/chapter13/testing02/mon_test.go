package main 

import (
	"testing"
)

func TestStore(t *testing.T) {

	//先创建以个Monster实例
	monster := &Monster{
		Name : "红孩儿",
		Age : 108,
		Skill : "三昧真火",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("Store()执行错误，期望值=%v 实际值=%v\n", true, res)
	}
	t.Logf("monster.Store()测试成功！")
}

func TestReStore(t *testing.T) {
	var monster = &Monster{}
	res := monster.ReStore()
	if !res {
		t.Fatalf("ReStore()执行错误，期望值=%v 实际值=%v\n", true, res)
	}
	//多重判断
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.ReStore()错误，期望值是=%v 实际值为=%v" , "红孩儿", monster.Name)
	}
	t.Logf("monster.ReStore()测试成功！")
}

