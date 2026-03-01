package main

import (
	"fmt"
	"time"
)

// ПРИМЕР 1
/*
func main() {
	intCh := make(chan int)
	strCh := make(chan string)

	go func() {
		i := 1
		for {
			intCh <- 1
			i++

			time.Sleep(1000 * time.Millisecond)
		}
	}()

	go func() {
		i := 1
		for {
			strCh <- fmt.Sprintf("hi %d", i)
			i++

			time.Sleep(100 * time.Millisecond)
		}
	}()

	for {
		select {
		case num := <-intCh:
			fmt.Println("intCh: ", num)
		case str := <-strCh:
			fmt.Println("strCh: ", str)
		}
	}
}
*/

// ПРИМЕР 2

type Message struct {
	Author string
	Text   string
}

func main() {
	messageCh1 := make(chan Message)
	messageCh2 := make(chan Message)

	go func() {
		for {
			messageCh1 <- Message{
				Author: "John Doe",
				Text:   "Hello world!",
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			messageCh2 <- Message{
				Author: "Scott Ridley",
				Text:   "Alien!",
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	for {
		select {
		case msg1 := <-messageCh1:
			fmt.Println(msg1.Author, msg1.Text)
		case msg2 := <-messageCh2:
			fmt.Println(msg2.Author, msg2.Text)
		}
	}
}
