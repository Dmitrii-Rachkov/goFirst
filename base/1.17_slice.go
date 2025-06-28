package main

import "fmt"

/*
Срезы (slice) представляют последовательность элементов одного типа переменной длины.
В отличие от массивов длина в срезах не фиксирована и динамически может меняться, то есть можно
добавлять новые элементы или удалять уже существующие.

Срез определяется также, как и массив, за тем исключением, что у него не указывается длина:
*/
func oneSlice() {
	fmt.Println("Работа среза")
	var users = []string{"Tom", "Alice", "Kate"}
	fmt.Println(users[2])
	users[2] = "Katherine"
	fmt.Println(users[2])
	// К элементам среза обращение происходит также, как и к элементам массива - по индексу
	//и также мы можем перебирать все элементы с помощью цикла for:
	for _, value := range users {
		fmt.Println(value)

	}
}

/*
С помощью функции make() можно создать срез из нескольких элементов,
которые будут иметь значения по умолчанию:
*/
func makeSlice() {
	fmt.Println("Функция make в срезах")
	var users = make([]int, 3)
	for index, value := range users {
		fmt.Println(index, value)
	}
}

/*
Добавление в срез
Для добавления в срез применяется встроенная функция append(slice, value).
Первый параметр функции - срез, в который надо добавить, а второй параметр - значение, которое
нужно добавить. Результатом функции является увеличенный срез.
*/
func addSlice() {
	fmt.Println("Добавление в срез")
	users := []string{"Tom", "Alice", "Kate"}
	users = append(users, "Bob")
	for _, value := range users {
		fmt.Println(value)
	}
}

/*
Оператор среза

Оператор среза s[i:j] создает из последовательности s новый срез, который содержит элементы
последовательности s с i по j-1. При этом должно соблюдаться условие 0 <= i <= j <= cap(s).
В качестве исходной последовательности, из которой берутся элементы, может использоваться массив,
указатель на массив или другой срез. В итоге в полученном срезе будет j-i элементов.

Если значение i не указано, то применяется по умолчанию значение 0.
Если значение j не указано, то вместо него используется длина исходной последовательности s.
*/
func operatorSlice() {
	fmt.Println("Использование оператора i:j для вывода определенных элементов среза")
	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	users1 := initialUsers[2:6] // с 3-го по 6-й
	users2 := initialUsers[:4]  // с 1-го по 4-й
	users3 := initialUsers[3:]  // с 4-го до конца

	fmt.Println(users1) // ["Kate", "Sam", "Tom", "Paul"]
	fmt.Println(users2) // ["Bob", "Alice", "Kate", "Sam"]
	fmt.Println(users3) // ["Sam", "Tom", "Paul", "Mike", "Robert"]
}

/*
Удаление элемента

Что делать, если необходимо удалить какой-то определенный элемент?
В этом случае мы можем комбинировать функцию append и оператор среза:
*/
func deleteSlice() {
	fmt.Println("Удаляем четвертый элемент среза")
	users := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	//удаляем 4-й элемент
	var n = 3
	users = append(users[:n], users[n+1:]...) // ломает исходный слайс
	fmt.Println(users)                        //["Bob", "Alice", "Kate", "Tom", "Paul", "Mike", "Robert"]

	// Ещё один способ удаления
	slice := []int{1, 2, 3, 4, 5}
	i := 2

	withCopy := slice[:i+copy(slice[i:], slice[i+1:])] // ломает исходный слайс
	fmt.Printf("Type: %T Value: %#v\n", withCopy, withCopy)
	fmt.Println(slice)
}

/*
Так как слайс имеет нефиксированную длину, "под капотом" лежит более сложная структура, чем у массива.
Помимо самих значений слайс хранит 2 дополнительных свойства: длину массива len (длина)
и cap (вместимость). Благодаря этому возможно инициализировать слайс нужной длины с помощью встроенной
функции func make(t Type, len, cap IntegerType) Type. Понимание, где лучше использовать какой
способ инициализации, приходит с опытом, но для старта рекомендуется использовать make везде, где можно:

// len = 5. Массив сразу будет заполнен 5-ю нулевыми значениями
nums := make([]int, 5, 5) // [0, 0, 0, 0, 0]

// len = 0, но cap = 5. Массив будет пустым, однако заполнение слайса через append будет эффективным,
потому что в памяти уже выделен массив нужной длины
nums := make([]int, 0, 5) // []
*/

/*
Копирование слайсов

Допустим, в вашей функции происходят изменения элементов, но вы не хотите затронуть входной слайс.
В языке есть встроенная функция func copy(dst, src []Type) int, которая копирует слайс src в слайс dst
и возвращает кол-во скопированных элементов:

nums := []int{1,2,3,4,5}
// важно инициализировать слайс той же длины
numsCp := make([]int, len(nums))
copy(numsCp, nums)
fmt.Println(numsCp) // [1,2,3,4,5]

Почему мы не можем просто перезаписать слайс в другую переменную и изменять ее?
Как и с функциями, при присваивании слайса к переменной, копируется только длина и вместимость,
но массив передается по ссылке:

nums := []int{1,2,3,4,5}
numsCp := nums
// исходный слайс nums тоже будет изменен
numsCp[0] = 10
fmt.Println(nums) // [10,2,3,4,5]

Существует распространенная ошибка, когда пытаются скопировать слайсы различной длины.
В этом случае элементы, выходящие за рамки слайса dst, не будут скопированы:

nums := []int{1, 2, 3, 4, 5}
// создали слайс с длиной 0
numsCp := make([]int, 0)
// при копировании в пустой слайс ничего не произойдет
copy(numsCp, nums)
fmt.Println(numsCp) // []
*/

/*
Сортировка слайсов

Сортировка массива — распространненая задача в программировании. Во всех языках существуют готовые
решения для этой задачи, и Go — не исключение. Стандартный пакет sort предоставляет функции
для сортировки:

nums := []int{2,1,6,5,3,4}
sort.Slice(nums, func(i, j int) bool {
    return nums[i] < nums[j]
})
fmt.Println(nums) // [1 2 3 4 5 6]

Рассмотрим функцию Slice(x interface{}, less func(i, j int) bool). В описании функции присутствует
неизвестный тип данных interface{}. Понятие интерфейса будет рассмотренно в следующих модулях курса.
Следует запомнить, что пустой интерфейс interface{} в Go означает тип данных, под который подходит
любой другой тип. Например:

func Print(arg interface{}) {
    fmt.Println(arg)
}

func main() {
    Print("hello!")
    Print(123)
    Print([]int{1,5,10})
}
Вывод:
hello!
123
[1 5 10]

То есть в функцию Slice(x interface{}, less func(i, j int) bool) передается слайс любого типа данных,
как первый аргумент. Вторым аргументом передается функция, которая берет элементы по индексу и
определяет должен ли элемент по индексу i находиться перед элементом по индексу j.

"Под капотом" в функции sort.Slice используется быстрая сортировка. В пакете также присутствует
сортировка вставками sort.SliceStable:

nums := []int{2,1,6,5,3,4}
sort.SliceStable(nums, func(i, j int) bool {
    return nums[i] < nums[j]
})
fmt.Println(nums) // [1 2 3 4 5 6]

Выбор алгоритма зависит от набора и размера данных, архитектуры процессора, скорости доступа к памяти,
то есть от многих факторов. Для большинства стандартных случаев используется sort.Slice, пока
производительность или нестабильность алгоритма не станет "узким горлышком".

Задание
Реализуйте функцию UniqueSortedUserIDs(userIDs []int64) []int64, которая возвращает отсортированный
слайс, состоящий из уникальных идентификаторов userIDs. Обработка должна происходить in-place,
то есть без выделения доп. памяти.

package solution
import "sort"

// UniqueSortedUserIDs sorts and removes duplicates from the userIDs slice.
func UniqueSortedUserIDs(userIDs []int64) []int64 {
	if len(userIDs) < 2 {
		return userIDs
	}

	sort.SliceStable(userIDs, func(i, j int) bool { return userIDs[i] < userIDs[j] })
	uniqPointer := 0
	for i := 1; i < len(userIDs); i++ {
		if userIDs[uniqPointer] != userIDs[i] {
			uniqPointer++
			userIDs[uniqPointer] = userIDs[i]
		}
	}
	return userIDs[:uniqPointer+1]
}
*/

// Конвертация slice в указатель на массив, при этом длина массива должна быть равна длине slice
func convertToArrayPointer() {
	initialSlice := []int{1, 2}
	fmt.Printf("Type: %T Value: %#v\n", initialSlice, initialSlice)
	fmt.Printf("Length: %d Capacity: %d\n\n", len(initialSlice), cap(initialSlice))

	intArray := (*[2]int)(initialSlice)
	fmt.Printf("Type: %T Value: %#v\n", intArray, intArray)
	fmt.Printf("Length: %d Capacity: %d\n\n", len(intArray), cap(intArray))
}

// Создание slice через функцию new
// Так мы получаем указатель на slice
func sliceWithNew() {
	slicePointer := new([]int)

	fmt.Printf("Type: %T Value: %#v\n", slicePointer, *slicePointer)
	fmt.Printf("Length: %d Capacity: %d\n\n", len(*slicePointer), cap(*slicePointer))

	newSlice2 := append(*slicePointer, 1)
	fmt.Printf("Type: %T Value: %#v\n", newSlice2, newSlice2)
	fmt.Printf("Length: %d Capacity: %d\n\n", len(newSlice2), cap(newSlice2))
}
