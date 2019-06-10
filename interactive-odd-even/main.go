package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	n, _ := strconv.Atoi(os.Getenv("n"))
	even := make(chan int)
	finish := make(chan bool)
	evenOnceDown := true
	go func() {
		for i := 1; i <= n; i++ {
			if !evenOnceDown {
				time.Sleep(time.Second * 1)
			}
			if i%2 != 0 {
				fmt.Println("odd => ", i)
			} else {
				evenOnceDown = false
				even <- i
			}
		}

		close(even)
		finish <- true
	}()

	go func() {
		for i := range even {
			fmt.Println("even => ", i)
			evenOnceDown = true
		}
	}()
	<-finish
}
