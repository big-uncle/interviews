package main

import (
	"fmt"
	"sync"
)

//顺序输出0-9
// 如果加锁的话，那么开销比较大而且输出的是 1-10，不符合预期结果
func main() {
	aa := make(chan int, 10)
	wg := sync.WaitGroup{}

	wg.Add(10)
	go func() {
		for {
			fmt.Println(<-aa)
			wg.Done()
		}
	}()
	for i := 0; i < 10; i++ {
		aa <- i
	}

	wg.Wait()
}
