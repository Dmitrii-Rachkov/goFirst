// Пример 1

package main

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

// Пример 2

//package main
//
//import "fmt"
//
//type Friend struct {
//	name string
//	age  int
//}
//
//func createFriend(friend **Friend) **Friend {
//	(*friend).name = "Anna"
//	(*friend).age = 10
//	return friend
//}
//
//func main() {
//	// Создаём пользователя и печатаем значение полей
//	// указатель user получает адрес в памяти на структуру User
//	friend := &Friend{name: "Petya", age: 20}
//	fmt.Println(friend)
//
//	// Пытаемся изменить значение полей
//	// внутри функции создается копия исходной структуры, и любые изменения, сделанные внутри функции,
//	// затрагивают только эту копию, а не оригинальный объект
//	// Чтобы изменения были внесены в оригинальный объект, необходимо передавать указатель на структуру
//	// func createUser(user *User) *User
//	createFriend(&friend)
//	// В итоге поля не изменились
//	fmt.Println(friend)
//}

// Пример 3

//package main
//
//import "fmt"
//
//type User struct {
//	name string
//	age  int
//}
//
//func createUser(user *User) {
//	user = &User{
//		name: "Anna",
//		age:  10,
//	}
//}
//
//func main() {
//	// Создаём пользователя и печатаем значение полей
//	// указатель user получает адрес в памяти на структуру User
//	user := &User{name: "Petya", age: 20}
//	fmt.Println(user)
//
//	// Пытаемся изменить значение полей
//	// внутри функции createUser создается копия исходной структуры, и любые изменения, сделанные внутри функции,
//	// затрагивают только эту копию, а не оригинальный объект.
//	// Чтобы изменения были внесены в оригинальный объект, необходимо передавать указатель на указатель на структуру
//	// func createUser(user **User) и внутри разыменовать
//	// *user = &User{
//	//		name: "Anna",
//	//		age:  10,
//	//	}
//	// а функцию createUser(&user) вызывать передав указатель на user
//	createUser(user)
//	fmt.Println(user)
//}

// Или в функции createUser работать напрямую с полями структуры user.name = "Anna"
