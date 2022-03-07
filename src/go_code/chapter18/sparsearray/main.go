package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"io"

)

type ValNode struct {
	row int
	col int
	val int
}

func writetofile(str string){
	filename := "f:/chessmap.data"
	//打开文件
	file, err := os.OpenFile(filename, os.O_TRUNC | os.O_WRONLY | os.O_CREATE , 0666 )
	if err != nil {
		fmt.Println(err)
	}

	//将关闭文件推入defer栈中
	defer file.Close()

	//还是以缓冲区的方式来对文件内容进行添加
	//func NewWriter(w io.Writer) *Writer，该函数的作用是，通过传入一个接口(接口可以接收所有的数据类型)
	//从而返回具有默认缓存的一个Reader结构体的指针类型（与NewReader（）函数十分相似，几乎一模一样）
	writer := bufio.NewWriter(file)

	//发送
	writer.WriteString(str)

	//因为以上的操作都是在缓存中进行的，要是不使用func (b *Writer) Flush() error，
	//Flush方法将缓冲中的数据写入下层的io.Writer接口。
	//即该函数对缓存的数据刷新，而使得缓冲区的数据传入到文件中去。
	writer.Flush()
}

func readonfile() {
	filename := "f:/chessmap.data"
	//打开文件
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("open file error = %v", err)
		return
	}
	//关闭文件
	defer file.Close()

	//文件以缓存的方式导出
	reader := bufio.NewReader(file)

	num := 0
	
	var rowlen int
	var collen int
	var row1 int
	var col1 int
	var val1 int

	var row2 int
	var col2 int
	var val2 int

	var data1 string
	var data2 string
	var data3 string
	var allnumb [6]int
	

	for {
		//func (b *Reader) ReadString(delim byte) (line string, err error)
		//ReadString读取直到第一次遇到delim字节，返回一个包含已读取的数据和delim字节的字符串
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if num == 0 {
			for id, val := range str{
				if id <= 1 {
					data1 = data1 + string(val)
				}else if id > 2 && id <= 4 {
					data2 = data2 + string(val)
				}
				// if err != nil {
				// 	// //rowlen = int(val)
				// 	// rowlen, _ = strconv.Atoi(string(val))
				// 	// //collen = int(val)
				// 	// collen, _ = strconv.Atoi(string(val))
				// }
			}
			rowlen, _ = strconv.Atoi(data1)
			collen, _ = strconv.Atoi(data2)
			fmt.Println("字符串转数字")
			fmt.Println(rowlen, collen)
		}else if num == 1 {
			data1 = ""
			data2 = ""
			data3 = ""
			for id, val := range str{
				if id == 0 {
					data1 = data1 + string(val)
				}else if id == 2 {
					data2 = data2 + string(val)
				}else if id == 4 {
					data3 = data3 + string(val)
				}
			}
			row1, _ = strconv.Atoi(data1)
			allnumb[0] = row1
			col1, _ = strconv.Atoi(data2)
			allnumb[1] = col1
			val1, _ = strconv.Atoi(data3)
			allnumb[2] = val1
			fmt.Println(row1, col1, val1)
			
		}else if num == 2 {
			data1 = ""
			data2 = ""
			data3 = ""
			for id, val := range str{
				if id == 0 {
					data1 = data1 + string(val)
				}else if id == 2 {
					data2 = data2 + string(val)
				}else if id == 4 {
					data3 = data3 + string(val)
				}
			}
			row2, _ = strconv.Atoi(data1)
			allnumb[3] = row2
			col2, _ = strconv.Atoi(data2)
			allnumb[4] = col2
			val2, _ = strconv.Atoi(data3)
			allnumb[5] = val2
			fmt.Println(row2, col2, val2)
		
	}
	num++
	// for id1, v1 := range chessmap {
	// 	for id2, v2 := range v1 {
	// 		fmt.Printf("%d\t", v2)
	// 		if id2 > collen {
	// 			continue
	// 		}
	// 	}
	// 	fmt.Println()
	// 	if id1 > rowlen {
	// 		break
	// 	}
	// }
}
	// collen1 := collen
	// rowlen1 := rowlen

	//数组范围不能用变量来确定，该如何是好！！！！！！
	
	var chessmap [11][11]int
	chessmap[allnumb[0]][allnumb[1]] = allnumb[2]
	chessmap[allnumb[3]][allnumb[4]] = allnumb[5]
	//2.输出看看原始数组
	for _, v1 := range chessmap {
		for _, v2 := range v1 {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

}

func main() {
	//1.创建一个原始数组
	var chessmap [11][11]int
	chessmap[1][2] = 1
	chessmap[2][3] = 2

	//2.输出看看原始数组
	for _, v1 := range chessmap {
		for _, v2 := range v1 {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	//3.转成稀疏数组
	//因为程序是不知道原先数组的有效数组事多少的，得用切片进行存储
	//思路
	//(1).遍历chessmap,如果我们发现有一个元素的值不为0，就创建一个node结构体
	//(2).将其放入到对应的切片即可
	var sparseArr []ValNode

	//标准的一个稀疏数组应该还有一个记录元素的二维数组的规模(行和列，默认值)
	//传入一个valnode值节点
	valNode := ValNode{
		row : 11,
		col : 11,
		val : 0,
	}
	sparseArr = append(sparseArr, valNode)

	for i, v1 := range chessmap {
		for j, v2 := range v1 {
			if v2 != 0 {
				//创建一个值节点
				valNode := ValNode{
					row : i,
					col : j,
					val : v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	//输出稀疏数组
	var str string
	fmt.Println("当前的稀疏数组是。。。。。")
	for i, valNode := range sparseArr{
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
		str = str + fmt.Sprintf("%d %d %d\r\n", valNode.row, valNode.col, valNode.val)
	}

	fmt.Println("字符串化", str)

	//将这个稀疏数组，存盘f:/chessmap.data
	writetofile(str)
	//如何恢复原始的数组

	//1.打开这个f:/chessmap.data， ==》 恢复原始数组
	readonfile()
	//2.这里使用稀疏数组恢复

	//先创建一个原始数组


}