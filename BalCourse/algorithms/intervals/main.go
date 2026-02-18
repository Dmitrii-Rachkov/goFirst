package main

import (
	"fmt"
	"sort"
)

/*
Задачи на интервалы
*/

/*
Есть два отрезка:

a1 ------------- b1

       a2 --------------- b2

a1 <= b1
a2 <= b2

- Как понять, пересекаются эти два отрезка или нет?
Это можно понять по такой формуле:

max(a1,a2) <= min(b1,b2)

1. То есть мы ищем максимальное значение из a1 и a2
2. Ищем минимальное значение из b1 и b2
3. Сравниваем максимальное значение a и минимальное значение b
4. Если a меньше или равно b, то отрезки пересекаются
5. Иначе если b больше a отрезки не пересекаются.

- Как найти длину пересечения, подотрезок пересечения, из примера выше (a2,b1)
Вот формула:
[max(a1,a2), min(b1,b2)]

- Как объединять два интервала (отрезка), из примера выше чтобы на выходе получился (a1,b2)?
Думаю тут формула будет [min(a1,a2), max(b1,b2).
То есть ищем начало интервала - это минимальный a,
и конец второго интервала - это максимальный b
*/

/*
56. Merge intervals
Дан массив интервалов intervals, где intervals[i] = [start, end].
Нужно объединить все пересекающиеся интервалы и вернуть массив непересекающихся интервалов,
который покрывает все входные.

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]

Алгоритм

1. Отсортировать интервалы по началу.
2. Завести результат res.
3. Пройти по всем интервалам:
Если res пуст или текущий интервал не пересекается с последним в res → добавить в res.
Если пересекается → обновить end последнего интервала как max(end1, end2).

Также надо на собесе спрашивать, что делать с интервалами если их точки пересекаются?
Объединять или нет? Вот такие например: [1,4][4,6]. Точка 4 пересекается.

time:
Мы сортируем массив из n интервалов.
Стоимость: O(n log n).
Обход всех интервалов:

Мы один раз проходим массив после сортировки.
Стоимость: O(n).

Итого:
O(n log n + n) = O(n log n)

memory:
Мы создаём результат res, который хранит объединённые интервалы.
В худшем случае O(n)
*/

func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// Сортировка по началу интервала
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// В res кладём первый интервал [[1,3]]
	res := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1]
		current := intervals[i]

		// Проверяем пересечение
		if current[0] <= last[1] {
			// объединяем
			if current[1] > last[1] {
				last[1] = current[1]
			}
			res[len(res)-1] = last
		} else {
			// нет пересечения
			res = append(res, current)
		}
	}

	return res
}

/*
1094. Car pooling

Условие

У нас есть машина с capacity мест.
Дан массив trips, где trips[i] = [numPassengers, from, to] означает:

numPassengers человек садятся в машине на остановке from, и выходят на остановке to.
Нужно вернуть true, если в машине никогда не окажется пассажиров больше чем capacity.

Идея (Difference Array / Sweep Line):
Вместо симуляции каждой остановки, используем разницу (difference array):
При посадке +numPassengers на from.
При высадке -numPassengers на to.

Потом проходим по массиву (в порядке остановок) и считаем текущее количество пассажиров.
Если в какой-то момент > capacity → возвращаем false.

time:
Проход по trips → O(n)
Проход по массиву остановок (до 1000) → O(1), т.к. фиксированная константа.
Итог: O(n).

memory:
Храним массив diff длиной 1001.
O(1) (константа).
*/

func СarPooling(trips [][]int, capacity int) bool {
	// Максимальное значение остановки <= 1000 по условию задачи
	const maxStop = 1000
	diff := make([]int, maxStop+1)

	// Заполняем difference array
	for _, trip := range trips {
		num, from, to := trip[0], trip[1], trip[2]
		diff[from] += num
		diff[to] -= num
	}

	// Считаем пассажиров на каждой остановке
	passengers := 0
	for i := 0; i <= maxStop; i++ {
		passengers += diff[i]
		if passengers > capacity {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("56. Merge intervals")
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(MergeIntervals(intervals)) // [[1 6] [8 10] [15 18]]

	fmt.Println("1094. Car pooling")
	trips := [][]int{{2, 1, 5}, {3, 3, 7}}
	capacity := 4
	fmt.Println(СarPooling(trips, capacity)) // false

	trips2 := [][]int{{2, 1, 5}, {3, 3, 7}}
	capacity2 := 5
	fmt.Println(СarPooling(trips2, capacity2)) // true
}
