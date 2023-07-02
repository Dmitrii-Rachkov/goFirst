package main

import "fmt"

/* Константы

Константы, как и переменные, хранят некоторые данные, но в отличие от переменных
значения констант нельзя изменить, они устанавливаются один раз.
Вычисление констант производится во время компиляции. Благодаря этому уменьшается количество работы,
которую необходимо произвести во время выполнения, упрощается поиск ошибок, связанных с константами
(так как некоторые из них можно обнаружить на момент компиляции).

Для определения констант применяется ключевое слово const:

const pi float64 = 3.1415

И в отличие от переменной мы не можем изменить значение константы.
А если и попробуем это сделать, то при компиляции мы получим ошибку:

const pi float64 = 3.1415
pi = 2.7182             // ! Ошибка

В одном определении можно объявить разу несколько констант:

const (
    pi float64 = 3.1415
    e float64 = 2.7182
)

Или так:
const pi, e = 3.1415, 2.7182

Если у константы не указан тип, то он выводится неявно на основании того значения,
которым инициализируется константа:

const n = 5     //  тип int

В то же время необходимо обязательно инициализировать константу начальным значением при ее объявлении.
Например, следующие определения констант являются недопустимыми, так как они не инициализируются:

const d
const n int

Если определяется последовательность констант, то инициализацию значением можно опустить для
всех констант, кроме первой. В этом случае константа без значения полчит значение предыдущей константы:

const (
    a = 1
    b
    c
    d = 3
    f
)
fmt.Println(a, b, c, d, f)      // 1, 1, 1, 3, 3

Константы можно инициализировать только константными значениями, например, литералами типа чисел
или строк, или значениями других констант. Но инициализировать константу значением переменной мы не можем:
var m int = 7
// const k = m      // ! Ошибка: m - переменная
const s = 5     // Норм: 5 - числовая константа
const n = s     // Норм: s - константа
*/

func constantEx() {
	const num float64 = 5.086
	fmt.Println(num)
	const (
		a = 1
		b
		c
		d = 3
		f
	)
	fmt.Println(a, b, c, d, f)
}

/*
Для последовательных числовых констант следует использовать идентификатор iota, который присвоит
для списка чисел значения от 0:

package main

import "fmt"

const (
    zero = iota
    one
    two
    three
)

func main() {
    fmt.Println(zero, one, two, three) // 0 1 2 3
}
*/
