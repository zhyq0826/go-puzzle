```go
package main

import (
	"fmt"
	"time"
)

func call() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}

		go call()
	}()

	time.Sleep(time.Second * 1)
	proc()
}

func main() {
	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出
		call()
	}()

	select {}
}

func proc() {
	panic("ok")
}

```