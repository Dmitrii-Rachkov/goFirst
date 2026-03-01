package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Контекст позволяет нам расслаивать работу нашего приложения, мы можем создать родительский контекст и от него дочерние.
// Отменять отдельно дочерние контексты, чтобы завершить часть работы приложения.
// Отменять родительский контекст который отменяет все дочерние.
// Можем передавать через контекст пользовательские данные, JWT токены и др.

// foo - работает с родительским контекстом
func foo(ctx context.Context, n int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Foo завершилась :(", n)
			return
		default:
			fmt.Println("Foo", n)
		}

		time.Sleep(100 * time.Millisecond)
	}
}

// boo - работает с дочерним контекстом
func boo(ctx context.Context, n int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Boo завершилась :(", n)
			return
		default:
			fmt.Println("Boo", n)
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// Graceful shutdown
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)

	// Родительский и дочерний контекст
	ctxParent, cancelParent := context.WithCancel(context.Background())
	ctxChild, cancelChild := context.WithCancel(ctxParent)

	go foo(ctxParent, 1)
	go foo(ctxParent, 2)
	go foo(ctxParent, 3)

	go boo(ctxChild, 1)
	go boo(ctxChild, 2)
	go boo(ctxChild, 3)

	// Если пришел сигнал завершить программу
	<-chSignal

	time.Sleep(1 * time.Second)
	cancelChild()

	time.Sleep(1 * time.Second)
	cancelParent()

	time.Sleep(3 * time.Second)
	fmt.Println("Main завершился")
}
