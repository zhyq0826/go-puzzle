package main

import (
	"fmt"
	"sync"
)

func main() {
	// n, _ := strconv.Atoi(os.Getenv("n"))
	n := 10
	even := make(chan int)
	finish := make(chan bool)
	oddCond := sync.NewCond(&sync.Mutex{})
	evenIsFinished := true
	go func() {
		for i := 1; i <= n; i++ {
			if !evenIsFinished {
				oddCond.L.Lock()
				oddCond.Wait()
				oddCond.L.Unlock()
			}
			if i%2 != 0 {
				fmt.Println("odd => ", i)
			} else {
				evenIsFinished = false
				even <- i
			}
		}

		close(even)
		finish <- true
	}()

	go func() {
		for i := range even {
			fmt.Println("even => ", i)
			evenIsFinished = true
			oddCond.Signal()
		}
	}()
	<-finish
}
