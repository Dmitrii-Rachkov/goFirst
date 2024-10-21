package main

import "fmt"

/*
Отображения

Отображение или map представляет ссылку на хеш-таблицу - структуру данных, где каждый элемент
представляет пару "ключ-значение". При этом каждый элемент имеет уникальный ключ, по которому
можно получить значение элемента. Отображение определяется как объект типа map[K]V, где К представляет
тип ключа, а V - тип значения. Причем тип ключа K должен поддерживать операцию сравнения ==, чтобы
отображение могло сопоставить значение с одним из ключей и хеш-таблицы.

Например, определение отображения, которое в качестве ключей имеет тип string,
а в качестве значений - тип int:
*/
func mapOne() {
	fmt.Println("Отображения или map (ключ: значение")
	// var people map[string]int // Ключи представляют тип string, значения - тип int
	var people = map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}
	fmt.Println(people) // map[Tom:1 Bob:2 Sam:4 Alice:8]

	/*
		Как и в массиве или в срезе элементы помещаютя в фигурные скобки. Каждый элемент представляет
		пару ключ -значение. Вначале идет ключ и через двоеточие значение. Определение элемента
		завершается запятой.

		Для обращения к элементам нужно применять ключи:
	*/
	fmt.Println(people["Alice"]) // 8
	fmt.Println(people["Bob"])   // 2
	people["Bob"] = 32
	fmt.Println(people["Bob"]) // 32

	/*
		Для проверки наличия элемента по определенному ключу можно применять выражение if:
	*/
	if val, ok := people["Tom"]; ok {
		fmt.Println(val)
	}
	/*
		Если значение по заданному ключу имеется в отображении, то переменная ok будет равна true,
		а переменная val будет хранить полученное значение. Если переменная ok равна false,
		то значения по ключу в отображении нет.
	*/

	/*
		Для перебора элементов применяется цикл for:
		Важно, при каждом запуске элементы будут находится в рандомном порядке
	*/
	for key, value := range people {
		fmt.Println(key, value)
	}

	/*
			Функция make представляет альтернативный вариант создания отображения.
			Она создает пустую хеш-таблицу:
		people := make(map[string]int)
	*/

	/*
		Добавление и удаление элементов
		Для добавления элементов достаточно просто выполнить установку значения по новому ключу
		и элемент с этим ключом будет добавлен в коллекцию:
	*/
	people["Oleg"] = 128
	fmt.Println(people["Oleg"])

	/*
		Для удаления применяется встроенная функция delete(map, key), первым параметром которой
		является отображение, а вторым - ключ, по которому надо удалить элемент.
	*/
	delete(people, "Bob")
	fmt.Println(people)
}

/*
Ещё способы создания map

// default value
	var defaultMap map[int64]string

// slice by make
	mapByMake := make(map[string]string)
	fmt.Printf("Type: %T Value: %#v\n\n", mapByMake, mapByMake)

// slice by make with cap (указываем количество элементов для экономии памяти)
	mapByMakeWithCap := make(map[string]string, 3)
	fmt.Printf("Type: %T Value: %#v\n\n", mapByMakeWithCap, mapByMakeWithCap)

// slice by literal
	mapByLiteral := map[string]int{"Vasya": 18, "Dima": 20}
	fmt.Printf("Type: %T Value: %#v\n", mapByLiteral, mapByLiteral)
	fmt.Printf("Len: %d\n\n", len(mapByLiteral))

// slice by new
	mapWithNew := *new(map[string]string)
	fmt.Printf("Type: %T Value: %#v\n\n", mapWithNew, mapWithNew)
*/

/*
Проверка наличия элментов в map:

При чтении элемента по несуществующему ключу возвращается нулевое значение данного типа.
Это приводит к ошибкам логики, когда используется bool как значение. Для решения данной проблемы при
чтении используется вторая переменная, в которую записывается наличие элемента в мапе:

existedIDs := map[int64]bool{1: true, 2: true}
idExists, elementExists := existedIDs[2] // true, true
idExists, elementExists := existedIDs[225] // false, false
*/

/*
Мапы в Go всегда передаются по ссылке:

package main
import (
    "fmt"
)

func main() {
    m := map[int]string{1: "hello", 2: "world"}
    modifyMap(m)
    fmt.Println(m) // вывод: map[1:changed 2:world 200:added]
}

func modifyMap(m map[int]string) {
    m[200] = "added"
    m[1] = "changed"
}
*/

/*
Задание
Реализуйте функцию UniqueUserIDs(userIDs []int64) []int64, которая возвращает слайс,
состоящий из уникальных идентификаторов userIDs. Порядок слайса должен сохраниться.

package solution

// UniqueUserIDs removes duplicates from the userIDs slice saving the IDs order.
func UniqueUserIDs(userIDs []int64) []int64 {
	// пустая структура struct{} — это тип данных, который занимает 0 байт
	// используется, когда нужно проверять в мапе только наличие ключа
	processed := make(map[int64]struct{})

	uniqUserIDs := make([]int64, 0)
	for _, uid := range userIDs {
		if _, ok := processed[uid]; ok {
			continue
		}

		uniqUserIDs = append(uniqUserIDs, uid)
		processed[uid] = struct{}{}
	}
	return uniqUserIDs
}
*/

/*
Задание
Реализуйте функцию MostPopularWord(words []string) string, которая возвращает самое часто встречаемое
слово в слайсе. Если таких слов несколько, то возвращается первое из них.

package solution

// MostPopularWord returns most popular word from the words slice.
// If there are multiple popular words it returns the first one depending on the words slice order.
func MostPopularWord(words []string) string {
	wordsCount := make(map[string]int, 0)
	mostPopWord := ""
	max := 0

	for _, word := range words {
		wordsCount[word]++
		if wordsCount[word] > max {
			max = wordsCount[word]
			mostPopWord = word
		}
	}
	return mostPopWord
}
*/

/*
// unique values (достаём из slice только уникальные значения)

// Предположим есть slice в котором id = 45 повторяется
	users := []User{
		{
			Id:   1,
			Name: "Vasya",
		},
		{
			Id:   45,
			Name: "Petya",
		},
		{
			Id:   57,
			Name: "John",
		},
		{
			Id:   45,
			Name: "Petya",
		},
	}


// Создаём map в которую будем складывать ключи id(int64) и значения пустые struct{},
// Кладём struct{} так как она не занимает места в памяти
	uniqueUsers := make(map[int64]struct{}, len(users))

	for _, user := range users {
		// Проверяем есть ли в мапе ключ из slice и если его нет - добавляем ключ, а в качестве значения пустую struct{}
		if _, ok := uniqueUsers[user.Id]; !ok {
			uniqueUsers[user.Id] = struct{}{}
		}
	}
	fmt.Printf("Type: %T Value: %#v\n", uniqueUsers, uniqueUsers)
*/

/*
Быстрый поиск значения с помощью map

Например нам нужно найти какого-то пользователя по его id.
Это быстрей, чем функция findInSlice.
Мы на основе полученного slice создаём map в которую складываем уникальных пользователей
В итоге есть map из которой по id в функции findInMap мы достаём нужные данные всего за одну операцию, алгоритм O(1)

// find by value
	usersMap := make(map[int64]User, len(users))
	for _, user := range users {
		if _, ok := usersMap[user.Id]; !ok {
			usersMap[user.Id] = user
		}
	}

	fmt.Println(findInSlice(57, users))
	fmt.Println(findInMap(57, usersMap))

// Поиск в slice
// Эта функция низко производительная, требуется много итераций
// Алгоритм O(n) - зависит от количества элементов в slice
func findInSlice(id int64, users []User) *User {
	for _, user := range users {
		if user.Id == id {
			return &user
		}
	}

	return nil
}

// Поиск в map
func findInMap(id int64, usersMap map[int64]User) *User {
	if user, ok := usersMap[id]; ok {
		return &user
	}

	return nil
}
*/
