package main
import (
	"fmt"
	"io/ioutil"
)

func main() {
	//将文件导入到另外一个文件
	//1.首先将f:/test01.txt文件内容导入到内存
	//2.将读取到的内容写入到f:\kkktest.txt

	file1 := "f:/test01.txt"
	file2 := "f:/kkktest.txt"

	//data使用个[]byte的切片类型
	data, err := ioutil.ReadFile(file1)
	if err != nil {
		fmt.Printf("read file err= %v\n", err)
		return
	}

	err = ioutil.WriteFile(file2, data, 0666)
	if err != nil {
		fmt.Printf("write file error= %v\n", err)
	}

}
