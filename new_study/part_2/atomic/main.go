package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var number atomic.Int64

func increase(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		number.Add(1)
	}
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(10)

	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)

	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)

	wg.Wait()

	fmt.Println("Total: ", number.Load())
}
