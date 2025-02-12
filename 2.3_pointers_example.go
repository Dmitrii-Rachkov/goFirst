//package main
//
//import "fmt"
//
//type User struct {
//	name string
//	age  int
//}
//
//func createUser(user User) User {
//	user.name = "Anna"
//	user.age = 10
//	return user
//}
//
//func main() {
//	// Создаём пользователя и печатаем значение полей
//	// указатель user получает адрес в памяти на структуру User
//	user := &User{name: "Petya", age: 20}
//	fmt.Println(user)
//
//	// Пытаемся изменить значение полей
//	// внутри функции создается копия исходной структуры, и любые изменения, сделанные внутри функции,
//	// затрагивают только эту копию, а не оригинальный объект
//	// Чтобы изменения были внесены в оригинальный объект, необходимо передавать указатель на структуру
//	// func createUser(user *User) *User
//	createUser(*user)
//	// В итоге поля не изменились
//	fmt.Println(user)
//}

package main

import "fmt"

type Friend struct {
	name string
	age  int
}

func createFriend(friend **Friend) **Friend {
	(*friend).name = "Anna"
	(*friend).age = 10
	return friend
}

func main() {
	// Создаём пользователя и печатаем значение полей
	// указатель user получает адрес в памяти на структуру User
	friend := &Friend{name: "Petya", age: 20}
	fmt.Println(friend)

	// Пытаемся изменить значение полей
	// внутри функции создается копия исходной структуры, и любые изменения, сделанные внутри функции,
	// затрагивают только эту копию, а не оригинальный объект
	// Чтобы изменения были внесены в оригинальный объект, необходимо передавать указатель на структуру
	// func createUser(user *User) *User
	createFriend(&friend)
	// В итоге поля не изменились
	fmt.Println(friend)
}
