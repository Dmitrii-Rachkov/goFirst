package main

import "fmt"

const (
	jopa = "string"
	even = 2
)

func main() {
	// Переменные
	fmt.Println("This is Print")
	fmt.Println("New study code")

	score := 0
	fmt.Println(score)

	score = 50 + 10
	fmt.Println(score)

	score += 20
	fmt.Println(score)

	var text string
	text = "jopa"
	fmt.Println(text)

	otherText := "other"
	fmt.Println(otherText)

	drob := 5.0
	fmt.Println(drob)

	var boolean bool
	fmt.Println(boolean)

	fmt.Println(jopa)
	fmt.Println(even)

	ten := 10
	resDivide := ten / 3
	fmt.Println(resDivide)
	ostatok := ten % 3
	fmt.Println(ostatok)

	fmt.Println(10.0 / 3)

	var evenNum int64 = 55
	fmt.Println(evenNum)

	// Ветвления
	if score == 90 {
		fmt.Println("Score =", score)
	} else if drob == 5.0 {
		fmt.Println("Drob =", drob)
	}

	// Циклы
	for i := 0; i <= 10; i++ {
		fmt.Println("Цикл:", i)
		if i == 5 {
			break
		}
	}

	var gtFor = 0

BEGIN:
	if gtFor == 7 {
		goto END
	}
	fmt.Println("Goto num:", gtFor)
	gtFor++
	goto BEGIN
END:
	fmt.Println("Goto is END")

	hello()
	defer fmt.Println("Main defer")
	goodbye("main func")
	defer func() {
		fmt.Println("Finish main defer")
	}()

	number := 5
	pointNum := &number
	fmt.Println("Init Number: ", number)

	*pointNum = 10
	fmt.Println("Pointer num: ", *pointNum)
	fmt.Println("Number after: ", number)
}

func hello() {
	fmt.Println("I am hello func")
	defer fmt.Println("Defer in hello func")
	goodbye("hello func")
	defer fmt.Println("hello finish defer")
}

func goodbye(param string) {
	defer func() {
		fmt.Println("start goodbye defer")
	}()
	fmt.Println("I am goodbye in", param)
	defer fmt.Println("Defer goodbye in ", param)
}

/*
I am hello func
I am goodbye in
Defer goodbye in hello func
start goodbye defer
hello finish defer
Defer in hello func
I am goodbye in main func
Defer goodbye in main func
start goodbye defer
Finish main defer
Main defer
*/
