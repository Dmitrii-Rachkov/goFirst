package principles

import "fmt"

/*
3. Liskov Substitution Principle (LSP)
Подтипы должны быть заменяемыми на свои базовые типы.
*/

type Bird2 interface {
	Fly() string
}

type Sparrow struct{}

func (s Sparrow) Fly() string {
	return "Flying"
}

type Ostrich struct{}

func (o Ostrich) Fly() string {
	return "Cannot fly"
}

// Проблема с LSP, если мы ожидаем, что все птицы будут летать
func MakeBirdFly(b Bird2) {
	fmt.Println(b.Fly())
}
