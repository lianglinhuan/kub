package main
import (
	"fmt"
	"go_code/chapter11/encapsulate01/model"
)

func main() {
	account := model.NewAccount("liang", "66666", 50)
	if account != nil {
		fmt.Println("创建成功")
	}else {
		fmt.Println("创建失败")
	}
}