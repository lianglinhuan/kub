package model

import (
	"fmt"
)

type account struct {
	zhanghao string
	monery float64
	pwd string
}

func NewAccount (zhanghao string, pwd string, monery float64) *account{
	if len(zhanghao) < 6 || len(zhanghao) > 10{
		fmt.Println("账号的长度不对。。。")
		return nil
	}
	if len(pwd) != 6 {
		fmt.Println("密码的长度不对。。。")
		return nil
	}
	if monery < 20 {
		fmt.Println("输入的金额不对。。。")
		return nil
	}
	return &account {
	    zhanghao : zhanghao,
		pwd : pwd,
		monery : monery,
	}
}
