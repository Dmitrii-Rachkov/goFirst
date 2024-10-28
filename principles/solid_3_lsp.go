package principles

import "fmt"

/*
L) Liskov substitution principle / Принцип подстановки Лисков

Этот принцип гласит, что объекты старших классов должны быть заменимы объектами подклассов, и приложение
при такой замене должно работать так, как ожидается.
*/

// Давайте рассмотрим структуру Animal:

type Animal struct {
	Name string
}

func (a Animal) MakeSound() {
	fmt.Println("Animal sound")
}

// Теперь предположим, что мы хотим создать новую структуру Bird, представляющую определенный тип животного:

type Bird struct {
	Animal
}

func (b Bird) MakeSound() {
	fmt.Println("Chirp chirp")
}

/*
Этот принцип гласит, что объекты суперкласса должны быть заменяемы объектами подкласса без нарушения корректности программы.
Это помогает гарантировать, что отношения между классами четко определены и поддерживаются.
*/

type AnimalBehavior interface {
	MakeSound()
}

// MakeSound represent a program that works with animals and is expected
// to work with base class (Animal) or any subclass (Bird in this case)
func MakeSound(ab AnimalBehavior) {
	ab.MakeSound()
}

// a := Animal{}
// b := Bird{}
// MakeSound(a)
// MakeSound(b)

/*
Это демонстрирует наследование в Go, а также принцип подстановки Лисков, поскольку объекты подтипа Bird могут
использоваться везде, где ожидаются объекты базового типа Animal, не влияя на правильность программы.
*/
