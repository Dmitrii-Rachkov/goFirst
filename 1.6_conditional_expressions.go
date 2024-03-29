package main

import "fmt"

/* Условные выражения

Условные выражения представляют логические операции и операции отношения.
Они представляют некоторое условие и возвращают значение типа bool:
true (если условие истинно) или false (если условие ложно).

Операции отношения.
Операции отношения позволяют сравнить два значения. В языке Go есть следующие операции отношения:
*/

func relation() {
	var a int = 8
	var b int = 3
	var c bool = a == b // Операция "равно". Возвращает true, если оба операнда равны, и false, если они не равны:
	fmt.Println(c)      // false
	var d bool = a > b  // Операция "больше чем". Возвращает true, если первый операнд больше второго, и false, если первый операнд меньше второго:
	fmt.Println(d)      // true
	var e bool = a < b  // Операция "меньше чем". Возвращает true, если первый операнд меньше второго, и false, если первый операнд больше второго:
	fmt.Println(e)      // false
	var f bool = a <= b // Операция "меньше или равно". Возвращает true, если первый операнд меньше или равен второму, и false, если первый операнд больше второго:
	fmt.Println(f)      // false
	var g bool = a >= b // Операция "больше или равно". Возвращает true, если первый операнд больше или равен второму, и false, если первый операнд меньше второго:
	fmt.Println(g)      // true
	var h bool = a != b // Операция "не равно". Возвращает true, если первый операнд не равен второму, и false, если оба операнда равны:
	var i bool = a != 8
	fmt.Println(h) // true
	fmt.Println(i) // false
}

/* Логические операции

Логические операции сравнивают два условия.
Как правило, они применяются к отношениям и объединяют несколько операций отношения.
К логическим операциям относят следующие:
*/

func logicOperations() {
	// ! (операция отрицания)
	// Инвертирует значение. Если операнд равен true, то возвращает false, иначе возвращает true.
	var a bool = true
	var b bool = !a // false
	var c bool = !b // true
	fmt.Println(a, b, c)
	// && (конъюнкция, логическое умножение)
	// Возвращает true, если оба операнда не равны false. Возвращает false, если хотя бы один операнд равен false.
	fmt.Println("Логическое умножение")
	var d bool = 4 > 5 && 6 > 8   // false
	var e bool = 3 <= 5 && 10 > 8 // true
	fmt.Println(d)
	fmt.Println(e)
	var f bool = 4 > 5 && 10 > 8 // false если одно из условий false
	fmt.Println(f)
	var g bool = 3 <= 5 && 6 > 8 // false если одно из условий false
	fmt.Println(g)
	// || (дизъюнкция, логическое сложение)
	// Возвращает true, если хотя бы один операнд не равен false. Возвращает false, если оба операнда равны false.
	var h bool = 4 > 5 || 6 > 8   // false
	var i bool = 3 == 5 || 10 > 8 // true
	fmt.Println("Логическое сложение")
	fmt.Println(h)
	fmt.Println(i)
}
