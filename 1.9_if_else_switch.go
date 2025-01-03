package main

import "fmt"

/*
Условные конструкции

Условные конструкции проверяют истинность некоторого условия и в зависимости от результатов
проверки позволяют направить ход программы по одному из путей.

Конструкция if принимает условие - выражение, которое возвращает значение типа bool.
И если это условие истинно, то выполняется последующий блок инструкций.
*/
func ifElse() {
	fmt.Println("Конструкция if")
	a := 6
	b := 7
	if a < b {
		fmt.Println("a is less than b")
	}

	// Если необходимо задать альтернативную логику, которая выполняется,
	//в случае если условие неверно, то добавляется выражение else:

	fmt.Println("Конструкция if else")
	c := 6
	d := 5
	if c < d {
		fmt.Println("c is less than d")
	} else {
		fmt.Println("d is less than c")
	}

	// Если необходимо проверить несколько альтернативных вариантов, то можно добавить выражения else if:
	fmt.Println("Конструкция else if")
	f := 8
	g := 8
	if f < g {
		fmt.Println("f less g")
	} else if f > g {
		fmt.Println("g less f")
	} else {
		fmt.Println("f = g")
	}
}

/*
Switch

Конструкция switch проверяет значение некоторого выражения.
С помощью операторов case определяются значения для сравнения.
Если значение после оператора case совпадает со значением выражения из switch,
то выполняется код данного блока case.
*/

func switchSimple() {
	fmt.Println("Switch")
	a := 8
	switch a {
	case 9:
		fmt.Println("a = 9")
	case 8:
		fmt.Println("a = 8")
	case 7:
		fmt.Println("a = 7")
	}
	/*
		В качестве выражения конструкция switch использует переменную a.
		Ее значение последовательно сравнивается со значениями после операторов case.
		Поскольку переменная a равна 8, то будет выполняться блок case 8: fmt.Println("a = 8").
		Остальные блоки case не выполняются.
	*/
}

//При этом после оператора switch мы можем указывать любое выражение, которое возвращает значение.
//Например, операцию сложения:

func switchSum() {
	b := 7
	switch b + 2 {
	case 9:
		fmt.Println("b = 9")
	case 8:
		fmt.Println("b = 8")
	case 7:
		fmt.Println("b = 7")
	}
}

/*
Также конструкция switch может содержать необязательных блок default, который выполняется,
если ни один из операторов case не содержит нужного значения:
*/

func switchDefault() {
	c := 87
	switch c {
	case 9:
		fmt.Println("c = 9")
	case 8:
		fmt.Println("c = 8")
	default:
		fmt.Println("Значение не определено")
	}
}

/*
Также конструкция switch может содержать необязательных блок default, который выполняется,
если ни один из операторов case не содержит нужного значения:
*/

// Switch с инициализацией переменной
func switchMany() {
	switch d := 5; d {
	case 9:
		fmt.Println("d = 9")
	case 6, 5, 4:
		fmt.Println("Тройной кейс")
	}
}

// Switch с fallthrough, который говорит, что нужно сделать следующий case
func switchFallthrough() {
	switch d := 5; d {
	case 9:
		fmt.Println("d = 9")
		fallthrough
	case 6, 5, 4:
		fmt.Println("fallthrough")
	}
}

func switchConditionAfter(x int) {
	switch {
	case x > 9:
		fmt.Println("x > 9")
		fallthrough
	case x < 6:
		fmt.Println("x < 6")
	}
}

/*
Таблица истинности для &&:
true && true = true
true && false = false
false && true = false
false && false = false
*/

/*
Таблица истинности для ||:
true || true = true
true || false = true
false || true = true
false || false = false
*/

/*
Отрицание !
!true = false
!false = true
*/
