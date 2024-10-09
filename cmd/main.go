package main

import (
	"fmt"
	"time"
)

// Замыкание
func main() {

	myTimer := getTimer()
	time.Sleep(1 * time.Second)

	f := func() {
		myTimer()
	}

	f()
}

func getTimer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Time from start %v\n", time.Since(start))
	}
}
