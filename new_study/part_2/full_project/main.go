package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var coal atomic.Int64
	var mails []string

	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(3 * time.Second)
		minerCancel()
	}()

	go func() {
		time.Sleep(6 * time.Second)
		postmanCancel()
	}()

	coalTransferPoint := minerPool(minerContext, 2)
	mailTransferPoint := postmanPool(postmanContext, 2)

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range coalTransferPoint {
			coal.Add(int64(v))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range mailTransferPoint {
			mu.Lock()
			mails = append(mails, v)
			mu.Unlock()
		}
	}()

	wg.Wait()

	fmt.Println("coal:", coal.Load())

	mu.Lock()
	fmt.Println("mails:", len(mails))
	mu.Unlock()
}
