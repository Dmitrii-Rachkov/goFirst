package main

/*
Рекурсивная функция представляет такую функцию, которая вызывает саму себя. Рекурсивные функции представляют мощный инструмент для обработки рекурсивных структур данных, например, различных деревьев.

Например, определим функцию вычисления факториала числа, которая получает результат рекурсивным способом:
*/
func factorial(n uint) uint {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

/*
Здесь функция factorial получает некоторое положительное число, для которого надо вычислить факториал.
Полученный результат возвращается из функции. Вначале идет условие, что если число равно 0,
то функция возвращает 1. Иначе функция возвращает произведение числа n на результат этой же функции
для числа n-1

При создании рекурсивной функции в ней обязательно должен быть некоторый базовый вариант,
который использует оператор return и помещается в начале функции.
В случае с факториалом это if n == 0 {return 1}

И, кроме того, все рекурсивные вызовы должны обращаться к подфункциям, которые в конце концов сходятся
к базовому варианту. Так, при передаче в функцию положительного числа при дальнейших рекурсивных
вызовах подфункций в них будет передаваться каждый раз число, меньшее на единицу.
И в конце концов мы дойдем до ситуации, когда число будет равно 0, и будет использован базовый вариант.

Например, вызов factorial(4) фактически можно расписать следующим образом:
factorial(4)
4 * factorial(3)
4 * 3 * factorial(2)
4 * 3 * 2 * factorial(1)
4 * 3 * 2 * 1 * factorial(0)
4 * 3 * 2 * 1 * 1
*/

/*
Другим распространенным показательным примером рекурсивной функции служит функция, вычисляющая числа Фибоначчи.
n-й член последовательности Фибоначчи определяется по формуле: f(n)=f(n-1) + f(n-2),
причем f(0)=0, а f(1)=1.
*/
func fibonachi(n uint) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonachi(n-1) + fibonachi(n-2)
	}
}
