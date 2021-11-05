package main

import (
	"fmt"
	"runtime"
	"runtime/trace"
)

func Producer(ch chan int) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}

func Consumer(id int, ch chan int) {

	go Consumer(id*id,ch)
	for {
		value, ok := <-ch
		if ok {
			fmt.Printf("id: %d, recv: %d,%d\n", id, value,runtime.NumGoroutine())
		} else {
			fmt.Printf("id: %d, closed,%d \n", id,runtime.NumGoroutine())
			break
		}
	}
}

func main() {

	ch := make(chan int, 3)
	coNum := 10
	for i := 1; i <= coNum; i++ {
		go Consumer(i, ch)
	}
	go Producer(ch)
	trace.Stop()
	for i:=0;i<10000000000;i++{
		if i%1000000000==0{
			fmt.Println("runtime.NumGoroutine:",runtime.NumGoroutine())
		}
	}

}
