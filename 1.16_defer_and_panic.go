package main

import "fmt"

/*
defer
Оператор defer позволяет выполнить определенную функцию в конце программы, при этом не важно,
где в реальности вызывается эта функция. Например:
*/
func finish() {
	fmt.Println("Program has been finished")
}

func deferFunc() {
	fmt.Println("Функция defer")
	defer finish()
	fmt.Println("Program has been started")
	fmt.Println("Program is working")
}

/*
Здесь функция finish вызывается с оператором defer, поэтому данная функция в реальности будет
вызываться в самом конце выполнения программы, несмотря на то, что ее вызов определен
в начале функции deferFunc
*/

/*
Если несколько функций вызываются с оператором defer, то те функции, которые вызываются раньше,
будут выполняться позже всех. Например:
*/
func deferMany() {
	fmt.Println("Много функции defer")
	defer fmt.Println("Program has been finished")
	fmt.Println("Program is working")
	defer fmt.Println("Program has been started")
}

/*
panic
Оператор panic позволяет сгенерировать ошибку и выйти из программы:
*/
func divide(x, y float64) float64 {
	fmt.Println("Работа оператора panic")
	if y == 0 {
		panic("Division by zero")
	}
	return x / y
}

/*
Оператору panic мы можем передать любое сообщение, которое будет выводиться на консоль.
Например, в данном случае в функции divide, если второй параметр равен 0, то осуществляется
вызов panic("Division by zero!").

В функции main в вызове fmt.Println(divide(4, 0)) будет выполняться оператор panic,
поскольку второй параметр функции divide равен 0.
И в этом случае все последующие операции,которые идут после этого вызова,
не будут выполняться.
И в конце вывода будет идти диагностическая информация о том, где возникла ошибка.
*/

/*
Обработка ошибок

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)
// При объявлении функции мы указываем, что она возвращает значение типа string
// и может вернуть error - ошибку:
// Если переменная a == 43 то возвращаем 'ok' и пустоту
// Иначе возвращаем пустую строку и новую ошибку с текстом 'some error
func foo(a int) (string, error) {
	if a == 42 {
		return "ok", nil
	}
	return "", errors.New("some error")
}

func main() {
	s, err := foo(42)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(s)
}
*/
