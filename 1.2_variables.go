/*Переменная представляет именованный участок в памяти, который может хранить некотоое значение.
Для определения переменной применяется ключевое слово var, после которого идет имя переменной,
а затем указывается ее тип.
Имя переменной представляет произвольный идентификатор, который состоит из алфавитных и цифровых
символов и символа подчеркивания. При этом первым символом должен быть либо алфавитный символ,
 либо символ подчеркивания.*/

package main

import (
	"fmt"
	"unicode/utf8"
)

func variables1() {
	// Данная переменная называется hello и она представляет тип string, то есть некоторую строку.
	var hello string
	hello = "Variables"
	fmt.Println(hello)

	// Можно одновременно объявить сразу несколько переменных через запятую:
	var a, b, c string
	a, b, c = "Start", "Stop", "Step"
	fmt.Println(a, b, c)

	// Также можно сразу при объявлении переменной присвоить ей начальное значение.
	var x string = "X"
	fmt.Println(x)

	// Если мы хотим сразу определить несколько переменных и присвоить им начальные значения,
	//то можно обернуть их в скобки:
	var (
		name string = "Alex"
		age  int    = 25
	)
	fmt.Println(name, age)

	// Отличительной особенностью переменных является то, что их значение можно многократно изменять:
	var example string = "First var"
	fmt.Println(example)
	example = "Second var"
	fmt.Println(example)
	example = "Third var"
	fmt.Println(example)

	// Также допустимо краткое определение переменной в формате:
	// В этом случае тип данных явным образом не указывается, он выводится автоматически
	//из присваиваемого значения.
	surname := "Salomon"
	fmt.Println(surname)

	// Строки представляют набор байтов
	var s string = "qы"

	// функцией len() мы не можем считать кол-во символов, т.к. есть русский символ
	// русский символ занимает 2 байта, соответственно len(s) = 3 а не 2
	// чтобы посчитать длину в таком случае нужно преобразовать строку в utf-8 и считать символы
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	// выводим на печать индекс и значение символа в строке
	for i, r := range s {
		fmt.Println(i, r, string(r))
	}

	// можно преобразовать строку в руну и итерироваться по руне
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		fmt.Println(i, runes[i])
	}

	// также можем изменять руну
	runes[0] = 't'
	s2 := string(runes)
	fmt.Println(s2)
}
