package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// Generics позволяют делать обобщение типов, не нужно плодить разные функции, достаточно сделать одну,
// которая работает не зависимо от типов данных.
// Generics можно использования для обобщения типов функции/типов.
// В функциях используется в качестве аргументов и возвращаемых параметрах.
// Обобщенные типы - это наш

// Number - возможный тип который содержит Numbers
type Number interface {
	~int64 | float64
}

// CustomInt - Кастомный тип на основе базового типа, чтобы его использовать
// нужно объявить приближение типов, например в Number ~int64.
// Приближение типов работает только с примитивными типами данных.
type CustomInt int64

func (ci CustomInt) IsPositive() bool {
	return ci > 0
}

// Numbers - обобщенный тип - это slice Number
type Numbers[T Number] []T

func main() {
	// showSum()
	// showContains()
	// showAny()
	// unionInterfaceAndType()
	// typeApproximation()
	// user()
	useCustomMap()
}

// Эта функция показывает результат сложения для целых чисел и чисел с плавающей точкой.
// Сам расчёт сложения осуществляется в функции sum, с помощью generics.
func showSum() {
	floats := []float64{1.0, 2.0, 3.0}
	ints := []int64{1, 2, 3}

	// Нельзя передать данный slice в функцию sum
	//wrongFloats := []interface{}{"string", struct{}{}, true}

	// Передаём slice без указания типа данных
	fmt.Println(sum(floats))

	// Передаём slice с указанием типа данных
	fmt.Println(sum[int64](ints))
	//fmt.Println(sum(wrongFloats))
}

func showContains() {
	type Person struct {
		name     string
		age      int64
		jobTitle string
	}

	ints := []int64{1, 2, 3, 4, 5}
	fmt.Println("int:", contains(ints, 4))

	strings := []string{"Vasya", "Dima", "Katya"}
	fmt.Println("strings:", contains(strings, "Katya"))
	fmt.Println("strings:", contains(strings, "Sasha"))

	people := []Person{
		{
			name:     "Vasya",
			age:      20,
			jobTitle: "Programmer",
		},
		{
			name:     "Dasha",
			age:      23,
			jobTitle: "Designer",
		},
		{
			name:     "Pasha",
			age:      30,
			jobTitle: "Admin",
		},
	}

	fmt.Println("structs:", contains(people, Person{
		name:     "Vasya",
		age:      21,
		jobTitle: "Programmer",
	}))

	fmt.Println("structs:", contains(people, Person{
		name:     "Vasya",
		age:      20,
		jobTitle: "Programmer",
	}))
}

func showAny() {
	show(1, 2, 3)
	show("test1", "test2", "test3")
	show([]int64{1, 2, 3}, []int64{4, 5, 6})
	show(map[string]int64{
		"first":  1,
		"second": 2,
	})
	show(interface{}(1), interface{}("string"), any(struct{ name string }{name: "Vasya"}))
}

// Использование обобщенного типа
func unionInterfaceAndType() {
	// Обязательно указываем какой из обобщенных типов нужен
	var ints Numbers[int64]
	ints = append(ints, []int64{1, 2, 3, 4, 5}...)

	floats := Numbers[float64]{1.0, 2, 5, 3, 5}

	fmt.Println(sumUnionInterface(ints))
	fmt.Println(sumUnionInterface(floats))
}

func typeApproximation() {
	customInts := []CustomInt{1, 2, 3, 5, 6}
	castedInts := make([]int64, len(customInts))

	for idx, val := range customInts {
		castedInts[idx] = int64(val)
	}

	fmt.Println(sumUnionInterface(customInts))
	fmt.Println(sumUnionInterface(castedInts))
}

// T comparable - это интерфейс, который говорит, что типом Т может быть всё кроме map, slice
// и структур у которых в качестве полей используются map и slice.
// comparable - значит мы можем сравнивать значения этих типов.
func contains[T comparable](elements []T, searchEl T) bool {
	for _, el := range elements {
		if searchEl == el {
			return true
		}
	}

	return false
}

// V - это тип который доступен только внутри этой функции
// Далее в квадратных скобках перечисляем примитивные доступные на вход в эту функцию
// Также тип V мы можем использовать в качестве аргументов функции и типе возвращаемых значений
// Если пришёл int64, то и везде далее мы работаем с int64
// int64 | float64 - это constrain - ограничение типов
func sum[V int64 | float64](numbers []V) V {
	var sum V
	for _, num := range numbers {
		sum += num
	}

	return sum
}

func sumUnionInterface[V Number](numbers []V) V {
	var sum V
	for _, num := range numbers {
		sum += num
	}

	return sum
}

// any - это alias пустого интерфейса
// Функция принимает slice элементов и выводит на печать slice
func show[T any](entities ...T) {
	fmt.Println(entities)
}

// Add - Также в пакете golang X есть готовые обобщенные типы
func Add[T constraints.Ordered](a, b T) T {
	return a + b
}

// Ещё один пример для структур

// CustomData - Можем создавать свои кастомные обобщенные типы и использовать их в полях структур
type CustomData interface {
	constraints.Ordered | []byte | []rune
}

type User[T CustomData] struct {
	ID   int64
	Name string
	Data T // можно использовать interface{}, но каждый раз нужно приводить к нужному типу
}

func user() {
	u := User[int64]{
		ID:   1,
		Name: "Man",
		Data: 5,
	}

	fmt.Println(u)
}

// Ещё пример с map

// SimpleMap - что если мы хотим использовать разные типы для ключей map?
// Приходится возиться с interface{} и приводить к нужному типу
type SimpleMap map[interface{}]interface{}

// CustomMap Это можно исправить с помощью generics
type CustomMap[K comparable, V comparable] map[K]V

func useCustomMap() {
	m := make(CustomMap[int, string])
	m[3] = "ok"
	fmt.Println(m)
}
