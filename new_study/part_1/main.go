package main

import (
	"fmt"
	"time"
)

const (
	jopa = "string"
	even = 2
)

func main() {
	// Переменные
	fmt.Println("This is Print")
	fmt.Println("New study code")

	score := 0
	fmt.Println(score)

	score = 50 + 10
	fmt.Println(score)

	score += 20
	fmt.Println(score)

	var text string
	text = "jopa"
	fmt.Println(text)

	otherText := "other"
	fmt.Println(otherText)

	drob := 5.0
	fmt.Println(drob)

	var boolean bool
	fmt.Println(boolean)

	fmt.Println(jopa)
	fmt.Println(even)

	ten := 10
	resDivide := ten / 3
	fmt.Println(resDivide)
	ostatok := ten % 3
	fmt.Println(ostatok)

	fmt.Println(10.0 / 3)

	var evenNum int64 = 55
	fmt.Println(evenNum)

	// Ветвления
	if score == 90 {
		fmt.Println("Score =", score)
	} else if drob == 5.0 {
		fmt.Println("Drob =", drob)
	}

	// Циклы
	for i := 0; i <= 10; i++ {
		fmt.Println("Цикл:", i)
		if i == 5 {
			break
		}
	}

	var gtFor = 0

BEGIN:
	if gtFor == 7 {
		goto END
	}
	fmt.Println("Goto num:", gtFor)
	gtFor++
	goto BEGIN
END:
	fmt.Println("Goto is END")

	hello()
	defer fmt.Println("Main defer")
	goodbye("main func")
	defer func() {
		fmt.Println("Finish main defer")
	}()

	number := 5
	pointNum := &number
	fmt.Println("Init Number: ", number)

	*pointNum = 10
	fmt.Println("Pointer num: ", pointNum)
	fmt.Println("Number after: ", number)

	// Создать экземпляр структуры
	newSome := new(Some)
	newSome.SetName("Вася")
	fmt.Println("Имя структуры: ", newSome.name)

	// Массивы
	arrString := [...]string{"Люся", "Петя"}
	fmt.Println(arrString[1:])

	// Slice
	slice := []int{2, 3, 4}

	// Копируем в начало slice новый элемент
	newValue := 1
	slice = append([]int{newValue}, slice...)
	fmt.Println(slice)

	// Копируем несколько элементов в начало slice
	slice2 := []int{2, 3, 4}
	newVal2 := []int{0, 1}
	slice2 = append(newVal2, slice2...)
	fmt.Println(slice2)

	// Копирование через сдвиг
	slice3 := []int{2, 3, 4}
	newValue3 := 1
	slice3 = append(slice3, 0)
	copy(slice3[1:], slice3)
	slice3[0] = newValue3
	fmt.Println(slice3)

	// Изменяем значение slice
	sliceUpdate := []int{1, 2, 3, 4, 5}
	fmt.Println("Source slice: ", sliceUpdate)
	UpdateSlice(sliceUpdate)
	fmt.Println("UpdateSlice: ", sliceUpdate)
	UpdateSlicePoint(&sliceUpdate)
	fmt.Println("UpdateSlicePoint: ", sliceUpdate)

	// Map
	// Если map не инициализирована var test map[string]int, то при попытке записи в неё будет panic
	var strIntMap = map[int]string{
		1: "One",
		2: "Two",
		3: "Three",
	}
	fmt.Println(strIntMap)
	strIntMap[4] = "Four"
	fmt.Println(strIntMap)
	strIntMap[2] = "undefined"
	fmt.Println(strIntMap)
	UpdateMap(strIntMap)
	fmt.Println(strIntMap)

	// Map или Slice
	// 1. добавить 100 элементов: map быстрей
	// 2. добавить 300: slice быстрей дальше
	// 3. поиск элемента: map быстрей
	// 4. получение элемента из slice по индексу быстрей, чем map
	// 5. Если нужна быстрая запись и поиск, используем гибрид из slice и map:
	/*
		// Гибрид из slice и map:
			type DataStore struct {
				items []Item
				index map[int]int // key -> index in slice
			}

			func (ds *DataStore) Add(id int, item Item) {
				ds.index[id] = len(ds.items)
				ds.items = append(ds.items, item) // Быстрая запись
			}

			func (ds *DataStore) Get(id int) Item {
				return ds.items[ds.index[id]] // Быстрый поиск
			}
	*/

	srez := make([]int64, 0, 1)
	mapa := make(map[int]int64)

	before := time.Now()
	for i := 0; i < 100_000; i++ {
		srez = append(srez, int64(i))
	}
	fmt.Println(time.Since(before))

	before = time.Now()
	for i := 0; i < 100_000; i++ {
		mapa[i] = int64(i)
	}
	fmt.Println(time.Since(before))

	before = time.Now()
	for _, value := range srez {
		if value == 100_000 {
			return
		}
	}
	fmt.Println(time.Since(before))

	before = time.Now()
	_, ok := mapa[100]
	if !ok {
		return
	}
	fmt.Println(time.Since(before))

}

func hello() {
	fmt.Println("I am hello func")
	defer fmt.Println("Defer in hello func")
	goodbye("hello func")
	defer fmt.Println("hello finish defer")
}

func goodbye(param string) {
	defer func() {
		fmt.Println("start goodbye defer")
	}()
	fmt.Println("I am goodbye in", param)
	defer fmt.Println("Defer goodbye in ", param)
}

/*
I am hello func
I am goodbye in
Defer goodbye in hello func
start goodbye defer
hello finish defer
Defer in hello func
I am goodbye in main func
Defer goodbye in main func
start goodbye defer
Finish main defer
Main defer
*/

// Some - структура
type Some struct {
	name string
	age  int64
}

func (s *Some) GetName() string {
	return s.name
}

func (s *Some) SetName(name string) {
	s.name = name
}

func UpdateSlice(s []int) {
	s[2] = 0           // Изменяется оригинал, т.к. дескриптор копируется, но указывается под капотом на исходный массив
	s = append(s, 100) // Изменяем КОПИЮ дескриптора! В main ничего не изменилось
}

func UpdateSlicePoint(s *[]int) {
	(*s)[4] = 0          // Меняет элемент в оригинальном массиве
	*s = append(*s, 100) // Меняет сам дескриптор среза в оригинале
}

func UpdateMap(source map[int]string) {
	source[1] = "NewValue"
	source[10] = "Ten"
}
