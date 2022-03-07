package main

import (
    "fmt"
    "strconv"
)
func main() {
	number := "what9"

    number_int, error := strconv.Atoi(number)
    if error == nil {
        fmt.Println("转换成功",number_int)
    }else {
        fmt.Println("转换错误,",error)
	}
	k, _ := strconv.Atoi("135")
    fmt.Println(k)
}
