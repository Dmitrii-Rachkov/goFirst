package main

import "fmt"

/*
Методы указателей

При вызове метода, объект структуры, для которого определен метод, передается в него по значению.
Что это значит? Рассмотрим следующий пример:
*/
type person struct {
	name string
	age  int
}

func (p person) updateAge(newAge int) {
	p.age = newAge
}

func tomPerson() {
	fmt.Println("Метод структуры")
	var tom = person{
		name: "Tom",
		age:  24,
	}
	fmt.Println("before", tom.age)
	tom.updateAge(33)
	fmt.Println("after", tom.age)
}

/*
Для структуры person определен метод updateAge, который принимает параметр newAge и изменяет
значение поля age у структуры. То есть при вызове этого метода мы ожидаем, что возраст человека изменится.
Однако консольный вывод нам показывает, что значение поля age не изменяется:
before 24
after 24

Потому что при вызове tom.updateAge(33) метод updateAge получает копию структуры tom.
То есть структура tom копируется в другой участок памяти, и далее метод updateAge работает с копией,
никак не затрагивая оригинальную структуру tom.

Однако такое поведение может быть нежелальтельным.
Что если мы всетаки хотим таким образом изменять состояние структуры?
В этом случае необходимо использовать указатели на структуры:
*/
func (p *person) updateAgePoint(newAge int) {
	(*p).age = newAge
}

func tomPersonPointer() {
	fmt.Println("Метод с указателем на структуры")
	var tom = person{
		name: "Tom",
		age:  24,
	}
	var tomPointer *person = &tom
	fmt.Println("before", tom.age)
	tomPointer.updateAgePoint(33)
	fmt.Println("after", tom.age)
}

/*
Теперь метод updateAge принимает указатель на структуру person: p *person, то есть фактически
адрес структуры в памяти. С помощью операции разыменования получаем значение по этому адресу
в памяти и меняем поле age:
(*p).age = newAge

В функции tomPersonPointer определяем указатель на структуру person и передаем ему адрес структуры tom:
var tomPointer *person = &tom

Затем вызываем метод updateAgePoint:
tomPointer.updateAge(33)

Таким образом, метод updateAge получит адрес, который хранится в tomPointer и по этому адресу
обратиться к структуре tom, изменив значение ее свойства age.
before 24
after 33

Стоит отметить, что несмотря на то, что метод updateAgePointer определен для указателя на структуру
person, мы по прежнему можем вызывать этот метод и для объекта person:
var tom = person { name: "Tom", age: 24 }
fmt.Println("before", tom.age)      // before 24
tom.updateAgePoint(33)
fmt.Println("after", tom.age)       // after 33
*/
