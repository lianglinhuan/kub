package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {
	start := time.Now().Unix()
	for num := 1; num <= 80000; num++ {
		flag := true
		for i := 2; i < num; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}
		if flag {
			
		}
	}
	end := time.Now().Unix()
	fmt.Println("普通的方法耗时=", end - start)
	cpunum := runtime.NumCPU()
    fmt.Printf("cpunum is %v\n", cpunum)

}