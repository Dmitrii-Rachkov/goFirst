package main

import "fmt"

/*
Массивы представляют последовательность элементов определенного типа.
Массив определяется следующим способом:

var numbers [число_элементов]тип_элементов
*/

func arrays() {
	fmt.Println("Массивы")
	// При таком определении все элементы массива инициализируются значениями по умолчанию.
	var numbers [5]int
	fmt.Println(numbers)

	// также можно инициализировать элементы массива другими значениями:
	var num [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(num)

	/*
		Значения передаются в фигурных скобках через запятую.
		При этом значений не может быть больше длины массива.
		В данном случае длина массива равна 5, поэтому нельзя в фигурных скобках определить
		больше пяти элементов. Но можно определить меньше элементов:
		В этом случае элементы, для которых не указано значение, будут иметь значение по умолчанию.
	*/
	var small [5]int = [5]int{1, 2}
	fmt.Println(small)

	// Также можно применять сокращенное определение переменной массива:
	exNumbers := [5]int{5, 4, 3, 2, 1}
	fmt.Println(exNumbers)

	// Если в квадратных скобках вместо длины указано троеточие, то длина массива определяется,
	// исходя из количества переданных ему элементов:
	var bigSize = [...]int{10, 9, 8, 7, 6} // длина 5
	fmt.Println(bigSize)
	size := [...]int{1, 2, 3} // длина 3
	fmt.Println(size)

	// При этом длина массива является частью его типа.
	// И, к примеру, следующие два массива представляют разные типы данных,
	// хотя они и хранят дванные одного типа:
	// И в данном случае при присвоении мы получим ошибку, так как данные одного типа
	// пытаемся передать переменной другого типа.
	// var ex1 [3]int = [3]int{1, 2, 3}
	// var ex2 [4]int = [4]int{1, 2, 3, 4}
	// ex1 = ex2  // ! Ошибка

	/*
		Индексы

		Для обращения к элементам массива применяются индексы - номера элементов.
		При этом нумерация начинается с нуля, то есть первый элемент будет иметь индекс 0. Индекс указывается в квадратных скобках. По индексу можно получить значение элемента, либо изменить его:
	*/

	var index [5]int = [5]int{7, 5, 6, 2, 8}
	fmt.Println(index[2]) // получение элемента по индексу
	index[0] = 4          // изменение элемента по индексу
	fmt.Println(index[0])

	/*
		Индексы в массиве фактически выступают в качестве ключей, по которым можно обратиться
		к соответствующему значению. И в прицнипе мы можем явным образом указать, какому ключу
		какое значение будет соответствовать. При этому числовые ключи необязательно располагать
		в порядке возрастания:
	*/
	colors := [3]string{2: "blue", 0: "red", 1: "green"}
	fmt.Println(colors[2])
}

/*
Двумерные массивы

var m = [3][3]int{
		{3, 4, 1},
		{5, 1},
	}

	for _, row := range m {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
*/
