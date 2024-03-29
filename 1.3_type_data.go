package main

import (
	"fmt"
	"strconv"
)

//Целочисленные типы:

//Ряд типов представляют целые числа:

//int8: представляет целое число от -128 до 127 и занимает в памяти 1 байт (8 бит)

//int16: представляет целое число от -32768 до 32767 и занимает в памяти 2 байта (16 бит)

//int32: представляет целое число от -2147483648 до 2147483647 и занимает 4 байта (32 бита)

//int64: представляет целое число от –9 223 372 036 854 775 808 до 9 223 372 036 854 775 807
//и занимает 8 байт (64 бита)

//uint8: представляет целое число от 0 до 255 и занимает 1 байт

//uint16: представляет целое число от 0 до 65535 и занимает 2 байта

//uint32: представляет целое число от 0 до 4294967295 и занимает 4 байта

//uint64: представляет целое число от 0 до 18 446 744 073 709 551 615 и занимает 8 байт

//byte: синоним типа uint8, представляет целое число от 0 до 255 и занимает 1 байт

//rune: синоним типа int32, представляет целое число от -2147483648 до 2147483647 и занимает 4 байта

//int: представляет целое число со знаком, которое в зависимости о платформы может занимать либо 4 байта,
//либо 8 байт. То есть соответствовать либо int32, либо int64.

//uint: представляет целое беззнаковое число только без знака, которое, аналогично типу int,
//в зависимости о платформы может занимать либо 4 байта, либо 8 байт.
//То есть соответствовать либо uint32, либо uint64.

//Стоит отметить типы int и uint.
//Они имеют наиболее эффективный размер для определенной платформы (32 или 64 бита).
//Это наиболее используемый тип для представления целых чисел в программе.
//Причем различные компиляторы могут предоставлять различный размер для этих типов даже
//для одной и той же платформы.
// Примеры:
//var a int8 = -1
//var b uint8 = 2
//var c byte = 3  // byte - синоним типа uint8
//var d int16 = -4
//var f uint16 = 5
//var g int32 = -6
//var h rune = -7     // rune - синоним типа int32
//var j uint32 = 8
//var k int64 = -9
//var l uint64 = 10
//var m int = 102
//var n uint = 105

//Числа с плавающей точкой:

//float32: представляет число с плавающей точкой от 1.4*10-45 до 3.4*1038(для положительных).
//Занимает в памяти 4 байта (32 бита)

//float64: представляет число с плавающей точкой от 4.9*10-324 до 1.8*10308 (для положительных)
//и занимает 8 байт.
// Примеры:
//var f float32 = 18
//var g float32 = 4.5
//var d float64 = 0.23
//var pi float64 = 3.14
//var e float64 = 2.7

// Комплексные числа:
//complex64: комплексное число, где вещественная и мнимая части представляют числа float32

//complex128: комплексное число, где вещественная и мнимая части представляют числа float64
//Примеры:
//var f complex64 = 1+2i
//var g complex128 = 4+3i

// Тип bool:
// Логический тип или тип bool может иметь одно из двух значений: true (истина) или false (ложь)
//var isAlive bool = true
//var isEnabled bool = false

// Строки:
//Строки представлены типом string.
//В Go строке соответствует строковый литерал - последовательность символов,
//заключенная в двойные кавычки:
// var name string = "Том Сойер"

// Кроме обычных символов строка может содержать специальные последовательности
//(управляющие последовательности), которые начинаются с обратного слеша \.
//Наиболее распространенные последовательности:
// \n: переход на новую строку
// \r: возврат каретки
// \t: табуляция
// \": двойная кавычка внутри строк
// \\: обратный слеш

//Значение по умолчанию
//Если переменной не присвоено значение, то она имеет значение по умолчанию,
//которое определено для ее типа. Для числовых типов это число 0, для логического типа - false,
//для строк - ""(пустая строка).

// Неявная типизация
// При определении переменной мы можем опускать тип в том случае,
// если мы явно инициализируем переменную каким-нибудь значением:
// var name = "Tom"

//В этом случае компилятор на основании значения неявно выводит тип переменной.
//Если присваивается строка, то то соответственно переменная будет представлять тип string,
//если присваивается целое число, то переменная представляет тип int и т.д.

//То же самое по сути происходит при кратком определении переменной,
//когда также явным образом не указывается тип данных:
//name := "Tom"

//При этом стоит учитывать, что если мы не указываем у переменной тип,
//то ей обязательно надо присвоить некоторое начальное значение.
//Объявление переменной одновременно без указания типа данных и начального значения будет ошибкой:
//var name    // ! Ошибка
//Надо либо указать тип данных (в этом случае переменная будет иметь значение по умолчанию):

//Неявная типизация нескольких переменных:
//var (
//	name = "Tom"
//	age = 27
//)
// Или так
//var name, age = "Tom", 27

func typeData() {
	var number int = 1500 // целое число
	fmt.Println(number)
}

func typeConversion() {
	fmt.Println("Преобразуем типы данных")
	// Строка в десятичное число (10) с точностью (64) разрядная система
	a, err := strconv.ParseInt("66", 10, 64)
	fmt.Println(a, err)

	// Строка в int по умолчанию в int64 без указанию доп параметров
	b, err := strconv.Atoi("99")
	fmt.Println(b, err)

	// Сканируем строку и записываем в переменную c int
	var c int
	_, err = fmt.Sscanf("77", "%d", &c)
	fmt.Println(c)

	// Преобразуем строку в float с указанием разрядности
	d, err := strconv.ParseFloat("33.1", 64)
	fmt.Println(d, err)

	// Числа в строки по умолчанию в int 64
	s1 := strconv.Itoa(42)
	fmt.Println(s1)

	// Числа в строки с параметров системы исчесления (10)
	s2 := strconv.FormatInt(123, 10)
	fmt.Println(s2)

	s3 := fmt.Sprintf("%d", 77)
	fmt.Println(s3)

	// Преобразование float в string с указанием разрадя (2 знака после запятой) и точности(64)
	s4 := strconv.FormatFloat(0.1, 'f', 2, 64)
	fmt.Println(s4)
}

/*
Преобразование строк
Для работы со строками часто используется стандартная библиотека strings.

// обрезает символы, переданные вторым аргументом, с обеих сторон строки
Trim(s, cutset string) string
// пример
strings.Trim(" hello ", " ") // "hello"

// преобразует все буквы в строке в нижний регистр
strings.ToLower(s string) string
// пример
strings.ToLower("пРиВеТ") // "привет"

// озаглавливает первую букву в каждом слове в строке
strings.Title(s string) string
// пример
strings.Title("привет, джон") // "Привет, Джон"

// функция проверяет, что строка name начинается с подстроки "Mr."
strings.HasPrefix(name, "Mr.")

// Для замены символов в строке существует функция
ReplaceAll(s, old, new string) string из пакета strings:
strings.ReplaceAll("hello world!", "world!", "buddy!") // hello buddy!

или пример:

// ModifySpaces modifies string s depending on mode.
func ModifySpaces(s, mode string) string {
	var replacement string

	switch mode {
	case "dash":
		replacement = "-"
	case "underscore":
		replacement = "_"
	default:
		replacement = "*"
	}

	return strings.ReplaceAll(s, " ", replacement)
}

// Наличие пробелов можно проверить с помощью функции
strings.Contains(firstName, " ").

// проверяет наличие подстроки в строке
strings.Contains("hello", "h") // true

// разбивает строку по Юникод символам или по переданному разделителю
strings.Split("hello", "") // ["h", "e", "l", "l", "o"]

// склеивает строки из слайса с разделителем
strings.Join([]string{"hello","world!"}, " ") // "hello world!"

// обрезает символы из второго аргумента в строке
strings.Trim(" hey !", " ") // "hey !"
*/

/*
Очень важная часть пакета strings — это Builder. Когда необходимо собрать большую строку по
каким-то правилам, использование конкатенации — не лучшее решение, потому что каждая операция
создает новую строку, что сильно влияет на производительность при большом количестве операций.
Такая задача решается с помощью билдера:

import "strings"

sb := &strings.Builder{}

sb.WriteString("hello")
sb.WriteString(" ")
sb.WriteString("world")

sb.String() // "hello world"
*/
