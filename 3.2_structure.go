package main

import "fmt"

/*
Структуры

Структуры представляют тип данных, определяемый разработчиком и служащий для представления каких-либо
объектов. Структуры содержат набор полей, которые представляют различные атрибуты объекта.
Для определения структуры применяются ключевые слова type и struct:

type имя_структуры struct{
    поля_структуры
}

Каждое поле имеет название и тип данных, как переменная.
Например, определим структуру, которая представляет человека:
*/

func structurePeople() {
	type person struct {
		name string
		age  int
	}
	/*
		Структура называется person. Она имеет два поля: name (имя человека, представляет тип string)
		и age (возраст человека, представляет тип int).
	*/

	/*
		Создание и инициализация структуры
		Структура представляет новый тип данных, и мы можем определить переменную данного типа:
		var tom person
		С помощью инициализатора можно передать структуре начальные значения:
	*/
	fmt.Println("Создание человека Tom с помощью структуры")
	var tom person = person{"Tom", 23}
	fmt.Println(tom)
	/*
			Инициализатор представляет набор значений в фигурных скобках.
			Причем эти значения передаются полям структуры в том порядке, в котором поля определены в структуре.
			Например, в данном случае строка "Tom" передается первому полю - name,
			а второе значение - 23 передается второму полю - age.

		Также мы можем явным образом указать какие значения передаются свойствам:
	*/
	var alice person = person{age: 29, name: "Alice"}
	fmt.Println(alice)
	jora := person{name: "Jora", age: 31}
	fmt.Println(jora)

	// Можно даже не указывать никаких значений, в этом случае свойства структуры
	//получат значения по умолчанию:
	undefined := person{}
	fmt.Println(undefined)

	/*
			Обращение к полям структуры

		Для обращения к полям структуры после переменной ставится точка и указывается поле структуры:
	*/
	fmt.Println(alice.name)
	fmt.Println(jora.age)

	tom.age = 38 // изменяем значение
	fmt.Println(tom.age)

	/*
			Указатели на структуры

		Как и в случае с обычными переменнами, можно создавать указатели на структуры.
	*/
	fmt.Println("Указатели в структурах")
	anton := person{name: "Anton", age: 77}
	var antonPointer *person = &anton
	antonPointer.age = 29
	fmt.Println(anton.age) // 77
	(*antonPointer).age = 32
	fmt.Println(anton.age) // 32

	/*
		Для инициализации указателя на структуру необязательно присваивать ему адрес переменной.
		Можно присвоить адрес безымянного объекта следующим образом:
	*/
	var kola *person = &person{name: "Kola", age: 15}
	fmt.Println(*kola)
	var loki *person = new(person)
	fmt.Println(*loki)
	/*
		Для обращения к полям структуры через указатель можно использовать операцию разыменования
		((*tomPointer).age), либо напрямую обращаться по указателю (tomPointer.age).
	*/

	/*
		Также можно определять указатели на отдельные поля структуры:
	*/
	fmt.Println("Указатели на отдельные поля структуры")
	serg := person{name: "Serg", age: 18}
	var agePointer *int = &serg.age // указатель на поле tom.age
	*agePointer = 35                // изменяем значение поля
	fmt.Println(serg.age)           //  35
}

/*
Если имя и поля структуры названы с маленькой буквы, то они доступны только внутри пакета.
Если имя и поля структуры названы с большой буквы, то они доступны в других пакетах.

type Foo struct {        // сама структура с большой буквы значит доступна в других пакетах
	X := 4              // доступно в других пакетах
}
*/

/*
У любого поля структуры можно указать теги. Они используются для метаинформации о поле для сериализации,
валидации, маппинга данных из БД и тд. Тег указывается после типа данных через бектики:
type User struct {
    ID int64 `json:"id" validate:"required"`
    Email string `json:"email" validate:"required,email"`
    FirstName string `json:"first_name" validate:"required"`
}

Тег json используется для названий полей при сериализации/десериализации структуры в json и обратно:

package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    ID        int64  `json:"id"`
    Email     string `json:"email"`
    FirstName string `json:"first_name"`
}

func main() {
    u := User{}
    u.ID = 22
    u.Email = "test@test.com"
    u.FirstName = "John"

    bs, _ := json.Marshal(u)

    fmt.Println(string(bs)) // {"id":22,"email":"test@test.com","first_name":"John"}
}

Тег validate используется Go-валидатором. В следующем примере присутствует вызов функции у структуры
v.Struct(u). Функции структур — методы — мы разберем подробно в соответствующем уроке, а пока
просто посмотрите, как происходит вызов:

package main

import (
    "fmt"
    "github.com/go-playground/validator/v10"
)

type User struct {
    ID        int64  `validate:"required"`
    Email     string `validate:"required,email"`
    FirstName string `validate:"required"`
}

func main() {
    // создали пустую структуру, чтобы проверить валидацию
    u := User{}

    // создаем валидатор
    v := validator.New()

    // метод Struct валидирует переданную структуру и возвращает ошибку `error`, если какое-то поле некорректно
    fmt.Println(v.Struct(u))
}
Вывод программы:

Key: 'User.ID' Error:Field validation for 'ID' failed on the 'required' tag
Key: 'User.Email' Error:Field validation for 'Email' failed on the 'required' tag
Key: 'User.FirstName' Error:Field validation for 'FirstName' failed on the 'required' tag
*/
