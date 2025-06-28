package main

import "fmt"

/*
Вложенные структуры

Поля одних структур могут представлять другие структуры. Например:
*/
func nestedStruct() {
	fmt.Println("Вложенные структуры")
	type contact struct {
		email string
		phone string
	}

	type person struct {
		name        string
		age         int
		contactInfo contact
	}

	var luda = person{
		name: "Luda",
		age:  45,
		contactInfo: contact{
			email: "luda45@gmail.com",
			phone: "+79884570102",
		},
	}

	luda.contactInfo.phone = "0102030405"
	fmt.Println(luda.contactInfo.email)
	fmt.Println(luda.contactInfo.phone)
	/*
			В данном случае структура person имеет поле contactInfo, которое представляет
		другую структуру contact.
	*/

	/*
				Можно сократить определение поля следующим образом:

			package main
			import "fmt"

			type contact struct{
			    email string
			    phone string
			}

			type person struct{
			    name string
			    age int
			    contact
			}

			func main() {

			    var tom = person {
			        name: "Tom",
			        age: 24,
			        contact: contact{
			            email: "tom@gmail.com",
			            phone: "+1234567899",
			        },
			    }
			    tom.email = "supertom@gmail.com"

			    fmt.Println(tom.email)      // supertom@gmail.com
			    fmt.Println(tom.phone)      // +1234567899
			}

		Поле contact в структуре person фактические эквивалентно свойству contact contact,
		то есть свойство называется contact и представляет тип contact.
		Это позволяет нам сократить путь к полям вложенной структуры.
		Например, мы можем написать tom.email, а не tom.contact.email.
		Хотя можно использовать и второй вариант.
	*/

	/*
			Хранения ссылки на структуру того же типа

			При этом надо учитывать, что структура не может иметь поле, которое представляет тип этой
			же структуры. Например:
			type node struct{
			value int
			next node
			}
		Подобное определение будет неправильным. Вместо этого поле должно представлять указатель на структуру:
	*/
	fmt.Println("Вывод односвязанного списка")
	type node struct {
		value int
		next  *node
	}
	first := node{value: 4}
	second := node{value: 5}
	third := node{value: 6}

	first.next = &second // связанная структура
	second.next = &third // связанная структура

	var current *node = &first
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
	/*
		Здесь определена структура node, которая представляет типичный узел односвязного списка.
		Она хранит значение в поле value и ссылку на следующий узел через указатель next.

		В функции создаются три связанных структуры, и с помощью цикла for
		и вспомогательного указателя current выводятся их значения.
	*/
}
