package main

import (
	"fmt"
)

/*
time: O(n) - обычно по времени нам нужно обойти все узлы
memory: O(1) - обычно по памяти мы не создаём новых структур
*/

/*
Это структура Node, которая представляет один узел в односвязном списке:
Value int — значение, которое хранится в узле. В данном случае это целое число (int).
Next *Node — указатель на следующий узел в списке. Это и делает список «связанным».

Например, если у нас три узла:
[1] -> [2] -> [3] -> nil,
то Node{Value: 1, Next: ptrToNode2} и так далее.
*/

// Node — элемент односвязного списка
type Node struct {
	Value int
	Next  *Node
}

/*
Это структура LinkedList, представляющая весь связный список:
Head *Node — указатель на первый узел (голову) списка.
Если список пуст — Head == nil.
*/

// LinkedList — структура списка
type LinkedList struct {
	Head *Node
}

/*
Это метод Push структуры LinkedList, который добавляет новый элемент в начало списка. Разберем его по шагам:
Входной параметр:
value int — значение, которое нужно добавить.
newNode - Создает новый узел:
Новый узел будет содержать переданное значение.
Его Next указывает на текущую Head — то есть он будет стоять перед текущим первым элементом.
ll.head - Обновляет Head, чтобы она указывала на новый узел. Теперь новый узел становится первым в списке.
*/

// Push — вставка в начало
func (ll *LinkedList) Push(value int) {
	newNode := &Node{Value: value, Next: ll.Head}
	ll.Head = newNode
}

/*
Append Добавляет элемент в конец списка.
Если Head == nil (список пуст), просто ставим новый узел первым.
Иначе:
Проходим по списку до последнего узла (Next == nil).
Делаем current.Next = newNode.
*/

// Append — вставка в конец
func (ll *LinkedList) Append(value int) {
	newNode := &Node{Value: value}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

/*
Print
Идёт от головы до конца (nil), печатая значения каждого узла.
*/

// Print — вывод списка
func (ll *LinkedList) Print() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

/*
Find Ищет узел по значению.
Ищет, сравнивая Value в каждом узле.
Возвращает указатель на найденный узел или nil, если не найден.
*/

// Find — поиск по значению
func (ll *LinkedList) Find(value int) *Node {
	current := ll.Head
	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}
	return nil
}

/*
Delete Удаляет первое вхождение элемента по значению.
Как работает:
Проверяет, не голова ли это (ll.Head.Value == value).
Иначе — идёт по списку, отслеживая предыдущий узел (prev).
Когда находит нужный current, делает prev.Next = current.Next, пропуская его.
*/

// Delete — удаление по значению (первое вхождение)
func (ll *LinkedList) Delete(value int) {
	if ll.Head == nil {
		return
	}
	if ll.Head.Value == value {
		ll.Head = ll.Head.Next
		return
	}
	prev := ll.Head
	current := ll.Head.Next
	for current != nil {
		if current.Value == value {
			prev.Next = current.Next
			return
		}
		prev = current
		current = current.Next
	}
}

/*
Length Возвращает количество узлов в списке.
Просто проходит по всем узлам и считает.
*/

// Length — возвращает количество элементов в списке
func (ll *LinkedList) Length() int {
	count := 0
	current := ll.Head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

/*
Reverse Разворачивает список наоборот, меняя направления связей.
1. prev = nil — будет указывать на предыдущий узел (изначально nil, потому что "после" последнего элемента ничего нет).
2. current = ll.Head — текущий обрабатываемый узел (начинаем с головы).
3. next := current.Next - Сохраняем ссылку на следующий узел
4. current.Next = prev - Переворачиваем ссылку
5. prev = current - Сдвигаем prev вперёд
6. current = next - Сдвигаем current вперёд
7. Переходим к следующему узлу, который мы сохранили ранее
8. ll.Head = prev После завершения цикла prev указывает на новый первый элемент списка (ранее последний)
*/

// Reverse — разворачивает список (in-place)
func (ll *LinkedList) Reverse() {
	var prev *Node = nil
	current := ll.Head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	ll.Head = prev
}

/*
Sort Сортирует список по возрастанию.

Классический алгоритм пузырьковой сортировки.
Пока находятся пары, которые нужно поменять местами — проход повторяется.
Обмениваются значения, а не сами узлы.
*/

// Sort — сортировка списка (с использованием пузырька)
func (ll *LinkedList) Sort() {
	if ll.Head == nil || ll.Head.Next == nil {
		return
	}
	swapped := true
	for swapped {
		swapped = false
		current := ll.Head
		for current.Next != nil {
			if current.Value > current.Next.Value {
				current.Value, current.Next.Value = current.Next.Value, current.Value
				swapped = true
			}
			current = current.Next
		}
	}
}

/*
Создаём "фиктивный" узел dummy, который указывает на голову списка.
Заводим два указателя: fast и slow, оба изначально указывают на dummy.
Двигаем fast на n шагов вперёд.
Потом двигаем оба указателя, пока fast.Next != nil.
Теперь slow.Next — это узел, который нужно удалить.
Меняем slow.Next = slow.Next.Next, тем самым удаляя нужный элемент.
*/

// DeleteNthFromEnd - удалить элемент находящийся на позиции n с конца списка
func (ll *LinkedList) DeleteNthFromEnd(n int) {
	// Со
	dummy := &Node{Next: ll.Head}
	fast := dummy
	slow := dummy

	// Сдвигаем fast на n шагов вперёд
	for i := 0; i < n; i++ {
		if fast.Next == nil {
			return // n больше длины списка
		}
		fast = fast.Next
	}

	// Двигаем оба указателя, пока fast не дойдёт до конца
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// Удаляем slow.Next (n-й с конца)
	slow.Next = slow.Next.Next

	// Обновляем голову (если был удалён первый элемент)
	ll.Head = dummy.Next
}

/*
Заводим два указателя: slow и fast.
Оба начинаем с головы (ll.Head).
В цикле:
slow двигается по одному узлу.
fast двигается по два узла.
Когда fast достигнет конца (nil), slow будет в середине.
*/

// FindSecondMiddle - поиск второй середины списка
func (ll *LinkedList) FindSecondMiddle() *Node {
	slow := ll.Head
	fast := ll.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// FindFirstMiddle - поиск первой середины
func (ll *LinkedList) FindFirstMiddle() *Node {
	slow := ll.Head
	fast := ll.Head

	// цикл останавливается раньше
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// ReverseFromNode - разворачивает часть листа начиная с определенной Node и возвращает первую перевернутую
func reverseFromNode(head *Node) *Node {
	var prev *Node
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func (ll *LinkedList) IsPalindrome() bool {
	if ll.Head == nil || ll.Head.Next == nil {
		return true
	}

	// 1. Найдём середину (slow окажется в середине)
	middle := ll.FindSecondMiddle()

	// 2. Разворачиваем вторую половину
	secondHalf := reverseFromNode(middle)

	// 3. Сравниваем первую и вторую половины
	p1 := ll.Head
	p2 := secondHalf
	isPalindrome := true

	for p2 != nil {
		if p1.Value != p2.Value {
			isPalindrome = false
			break
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// 4. (Необязательно) восстановить список
	reverseFromNode(secondHalf)

	return isPalindrome
}

/*
Найти середину списка.
Развернуть вторую половину списка.
Слить две половины, чередуя узлы из первой и второй.
*/

// Reorder - перестроить список
func (ll *LinkedList) Reorder() {
	if ll.Head == nil || ll.Head.Next == nil {
		return
	}

	// 1. Найдём середину
	slow, fast := ll.Head, ll.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2. Разворачиваем вторую половину
	second := reverseFromNode(slow.Next)
	slow.Next = nil // обрубаем первую половину

	// 3. Слияние двух половин
	first := ll.Head
	for second != nil {
		// сохраняем ссылки
		tmp1 := first.Next
		tmp2 := second.Next

		// чередуем
		first.Next = second
		second.Next = tmp1

		// сдвигаемся дальше
		first = tmp1
		second = tmp2
	}
}

// MergeSorted - слияние двух списков
func MergeSorted(l1, l2 *LinkedList) *LinkedList {
	dummy := &Node{}
	tail := dummy

	p1 := l1.Head
	p2 := l2.Head

	for p1 != nil && p2 != nil {
		if p1.Value < p2.Value {
			tail.Next = p1
			p1 = p1.Next
		} else {
			tail.Next = p2
			p2 = p2.Next
		}
		tail = tail.Next
	}

	// Добавляем оставшийся хвост
	if p1 != nil {
		tail.Next = p1
	}
	if p2 != nil {
		tail.Next = p2
	}

	return &LinkedList{Head: dummy.Next}
}

// MergeWith - слияние двух списков как метод
func (ll *LinkedList) MergeWith(other *LinkedList) {
	dummy := &Node{}
	tail := dummy

	p1 := ll.Head
	p2 := other.Head

	for p1 != nil && p2 != nil {
		if p1.Value < p2.Value {
			tail.Next = p1
			p1 = p1.Next
		} else {
			tail.Next = p2
			p2 = p2.Next
		}
		tail = tail.Next
	}

	if p1 != nil {
		tail.Next = p1
	}
	if p2 != nil {
		tail.Next = p2
	}

	ll.Head = dummy.Next
}

// MergeWithMany сливает текущий список с несколькими отсортированными списками
func (ll *LinkedList) MergeWithMany(others []*LinkedList) {
	for _, other := range others {
		ll.MergeWith(other)
	}
}

func main() {
	list := LinkedList{}

	// Добавим элементы
	list.Append(30)
	list.Append(10)
	list.Append(20)
	list.Push(40) // в начало

	fmt.Println("Исходный список:")
	list.Print()

	// Длина
	fmt.Println("Длина:", list.Length())

	// Сортировка
	list.Sort()
	fmt.Println("Отсортированный список:")
	list.Print()

	// Разворот
	list.Reverse()
	fmt.Println("Развёрнутый список:")
	list.Print()

	// Удаление
	list.Delete(20)
	fmt.Println("После удаления 20:")
	list.Print()

	// Пример удаление n элемента с конца списка
	listDeleteNthFromEnd := LinkedList{}
	listDeleteNthFromEnd.Append(10)
	listDeleteNthFromEnd.Append(20)
	listDeleteNthFromEnd.Append(30)
	listDeleteNthFromEnd.Append(40)
	listDeleteNthFromEnd.Append(50)

	fmt.Print("Исходный список: ")
	listDeleteNthFromEnd.Print() // 10 -> 20 -> 30 -> 40 -> 50 -> nil

	listDeleteNthFromEnd.DeleteNthFromEnd(2)

	fmt.Print("После удаления 2-го с конца: ")
	listDeleteNthFromEnd.Print() // 10 -> 20 -> 30 -> 50 -> nil

	// Пример вывода списка с второй середины
	listMiddle := LinkedList{}
	listMiddle.Append(10)
	listMiddle.Append(20)
	listMiddle.Append(30)
	listMiddle.Append(40)
	listMiddle.Append(50)

	mid := listMiddle.FindSecondMiddle()
	fmt.Print("Список, начиная со второй середины: ")
	current := mid
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")

	// Пример проверки на палиндром
	listIsPalindrome := LinkedList{}
	listIsPalindrome.Append(1)
	listIsPalindrome.Append(2)
	listIsPalindrome.Append(3)
	listIsPalindrome.Append(2)
	listIsPalindrome.Append(1)

	fmt.Println("Это палиндром?", listIsPalindrome.IsPalindrome()) // true

	// Пример Reorder list
	listReorder := LinkedList{}
	listReorder.Append(1)
	listReorder.Append(2)
	listReorder.Append(3)
	listReorder.Append(4)
	listReorder.Append(5)

	fmt.Print("Исходный список: ")
	listReorder.Print()

	listReorder.Reorder()

	fmt.Print("После reorder: ")
	listReorder.Print()

	// Слияние двух списков MergeSorted
	list1 := &LinkedList{}
	list1.Append(1)
	list1.Append(3)
	list1.Append(5)

	list2 := &LinkedList{}
	list2.Append(1)
	list2.Append(4)
	list2.Append(6)

	merged := MergeSorted(list1, list2)

	fmt.Print("Merged list: ")
	merged.Print()

	// Слияние нескольких списков MergeWithMany
	l1 := &LinkedList{}
	l1.Append(1)
	l1.Append(5)

	l2 := &LinkedList{}
	l2.Append(2)
	l2.Append(6)

	l3 := &LinkedList{}
	l3.Append(3)
	l3.Append(4)

	l1.MergeWithMany([]*LinkedList{l2, l3})

	fmt.Print("Merged list: ")
	l1.Print()
}
