package main
import (
	"fmt"
	"os"
	"bufio"
	"io"
)

type CharCount struct {
	ChCount int
	NumCount int
	SpaceCount int
	OtherCount int
}

func main() {

	filename := "f:/abctest.txt"

	//打开文件
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("open file error = %v", err)
		return
	}

	//关闭文件
	defer file.Close()

	var count CharCount

	//文件以缓存的方式导出
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		
		for _, val := range str {
			switch {
			case val >= 'a' && val <= 'z' :
				count.ChCount++
			case val >= 'A' && val <= 'Z' :
				count.ChCount++
			case val == ' ' || val == '\t' :
				count.SpaceCount++
			case val >= '0' && val <= '9' :
				count.NumCount++
			default :
			    count.OtherCount++

			}
			fmt.Printf(string(val))
		}
	}
	fmt.Printf("文件的字符数为%v，数字数为%v，空格或tab数为%v，其他字符数为%v",
	count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)


}