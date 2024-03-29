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
