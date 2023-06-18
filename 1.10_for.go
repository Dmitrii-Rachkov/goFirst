package main

import (
	"fmt"
)

/*
Цикл for

Циклы позволяют в зависимости от определенного условия выполнять некоторые действия множество раз.
Фактически в Go есть только один цикл - цикл for, который может принимать разные формы.
Этот цикл имеет следующее формальное определение:

for [инициализация счетчика]; [условие]; [изменение счетчика]{
    // действия
}
*/

// Например, выведем с помощью цикла квадраты чисел:

func for1() {
	fmt.Println("Простой цикл for")
	for i := 1; i < 3; i++ {
		fmt.Println(i * i)
	}
}

/*
Объявление цикла for разбивается на три части. Вначале идет инициализация счетчика: i := 1.
Фактически она представляет объявление переменной, которая будет использоваться внутри цикла.
В данном случае это счетчик i, начальное значение которого равно 1.

Вторая часть представляет условие: i < 10.
Пока это условие истинно, то есть возвращает true, будет продолжаться цикл.

Третья часть представляет изменение (увеличение) счетчика на единицу.

В теле цикла на консоль выводится квадрат числа i.
Таким образом, цикл сработает 9 раз, пока значение i не станет равным 10.
И каждый раз это значение будет увеличиваться на 1. Каждый отдельный проход цикла называется итерацией.
То есть в данном случае будет 9 итераций.
*/

// Нам необязательно указывать все условия при объявлении цикла.
//Например, можно вынести объявление переменной вовне:

func for2() {
	fmt.Println("Цикл for без обЪявления переменной")
	var i = 1
	for ; i < 3; i++ {
		fmt.Println(i * i)
	}
}

// Можно убрать изменение счетчика в само тело цикла и оставить только условие:
func for3() {
	fmt.Println("Изменение счётчика в теле цикла")
	var i = 1
	for i < 3 {
		fmt.Println(i * i)
		i++
	}
}

// Если цикл использует только условие, то его можно сократить следующим образом:
func for4() {
	fmt.Println("Цикл который использует только условие")
	var i = 1
	for i < 3 {
		fmt.Println(i * i)
		i++
	}
}

/*
Вложенные циклы

Циклы могут быть вложенными, то есть располагаться внутри других циклов.
Например, выведем на консоль таблицу умножения:
*/

func forFor() {
	fmt.Println("Цикл вложенный в другой цикл")
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}
}

/*
Перебор массивов

Для перебора массивов можно использовать следующую форму цикла for:

for индекс, значение := range массив{
    // действия
}

При переборе мы можем по отдельности получить индекс элемента в массиве и значение этого элемента.
Например, перебирем массив строк:
*/

func forArray() {
	fmt.Println("перебор массива")
	var users = [3]string{"Tom", "Alice", "Jerry"}
	for index, value := range users {
		fmt.Println(index, value)
	}
}

/*
Если мы не планируем использовать значения или индексы элементов, то мы можем вместо них указать прочерк.
Например, нам не нужны индексы:
*/
func forUnderscore() {
	fmt.Println("Перебор массива без индекса")
	var users = [3]string{"Tom", "Alice", "Jerry"}
	for _, value := range users {
		fmt.Println(value)
	}
}

/*
Но также для перебора массива можно использовать и стандартную версию цикла for:
В данном случае счетчик i играет роль индекса.Цикл выполняется, пока счетчик i не станет
равным длине массива, которую можно получить с помощью функции len()
*/
func forStandard() {
	fmt.Println("Стандартный перебор массива")
	var users = [3]string{"Tom", "Alice", "Jerry"}
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}
}

/*
Операторы break и continue
Может возникнуть ситуация, когда нам надо при определенных условиях завершить текущую итерацию цикла,
не выполнять все инструкции цикла, а сразу перейти к следующей итерации.
В этом случае можно использовать оператор continue. Например, в массиве могу быть, как положительные,
так и отрицательные числа. Допустим, нам нужна сумма только положительных чисел, поэтому,
если нам встретится отрицательное число, мы можем просто перейти к следующей итерации с помощью continue:
*/

func forContinue() {
	fmt.Println("Оператор Continue")
	var numbers = [10]int{1, -2, 3, -4, 5, -6, -7, 8, -9, 10}
	var sum = 0
	for _, value := range numbers {
		if value < 0 {
			continue // переходим к следующей итерации
		}
		sum += value
	}
	fmt.Println("Sum:", sum)
}

// Оператор break полностью осуществляет выход из цикла:
func forBreak() {
	fmt.Println("Оператор break")
	var numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var sum = 0
	for _, value := range numbers {
		if value > 4 {
			break // если число больше 4 выходим из цикла
		}
		sum += value
	}
	fmt.Println("Sum:", sum)
}
