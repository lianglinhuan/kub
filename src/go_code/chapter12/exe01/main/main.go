package main

import (
	"fmt"
	"io"
	"bufio"
	"os"
)

func main() {
	filename := "f:/test01.txt"

	file, err := os.OpenFile(filename, os.O_RDWR | os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("open file err =%v\n", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		// if err != nil {
		// 	fmt.Printf("read file err =%v", err)
		// 	break
		// }
		if err == io.EOF {
			fmt.Println("文件读取完成")
			break
		}
		fmt.Printf(str)
	}

	str := "你好，北京！\r\n"

	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	writer.Flush()
}