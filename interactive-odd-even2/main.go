package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	n, _ := strconv.Atoi(os.Getenv("n"))
	oddCond := sync.NewCond(&sync.Mutex{})
	evenCond := sync.NewCond(&sync.Mutex{})
	finish := make(chan bool)
	i := 1
	go func() {
		for i <= n {
			oddCond.L.Lock()
			oddCond.Wait()
			if i%2 != 0 {
				fmt.Println("odd =>", i)
			} else {
				evenCond.Signal()
			}
			i++
			oddCond.L.Unlock()
		}
		finish <- true
	}()
	go func() {
		for i <= n {
			evenCond.L.Lock()
			evenCond.Wait()
			fmt.Println("even =>", i)
			evenCond.L.Unlock()
		}
		finish <- true
	}()

	oddCond.Signal()
	<-finish
}
