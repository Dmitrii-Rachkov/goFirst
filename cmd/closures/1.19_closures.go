package main

import (
	"fmt"
)

// Замыкание
// Каждый раз при вызове внешней функции totalPrice, создаётся новый экземпляр внутренней функции замыкания

// Простой пример
func main() {
	// Есть переменная и её значение 30
	dollar := 30

	// Функция замыкания, которая замыкает в себе переменную dollar
	// при каждом вызове она идёт и смотрит актуальное значение переменной dollar
	getDollarValue := func() int {
		return dollar
	}

	// Печатаем значение переменной первый раз
	fmt.Println(getDollarValue())

	// Изменяем значение переменной и снова её печатаем
	dollar = 70
	fmt.Println(getDollarValue())

	// Сложный пример

	// Вызываем функцию первый раз и присваиваем переменной sum = 1
	orderPrice := totalPrice(1)

	// Все последующие вызовы запоминают предыдущее значение переменной sum
	fmt.Println(orderPrice(1))
	fmt.Println(orderPrice(1))
	fmt.Println(orderPrice(1))
	fmt.Println(orderPrice(1))
	fmt.Println(orderPrice(1))
}

// Сложный пример функции c замыканием
func totalPrice(intPrice int) func(int) int {
	sum := intPrice // при первом вызове присваиваем значение

	// Само замыкание
	return func(x int) int {
		sum += x
		return sum
	}
}
