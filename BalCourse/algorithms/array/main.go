package main

import "fmt"

/*
Задача 303. Range Sum Query - Immutable
нужно создать структуру NumArray, которая по массиву чисел умеет быстро отвечать
на запрос суммы подотрезка [left, right].

Чтобы не считать сумму заново каждый раз (O(n)), используем префиксные суммы:
prefix[i] = сумма элементов от nums[0] до nums[i-1]

Тогда:
sumRange(left, right) = prefix[right+1] - prefix[left]

Мы заранее считаем префиксные суммы — массив prefix, где prefix[i] — сумма всех элементов nums[0..i-1].
Тогда сумма подотрезка [left, right] получается за O(1)

time: O(n) - собираем массив из всех элементов
memory: O(n) - храним дополнительный префиксный массив
SumRange: O(1) - запрос выполняется за константное время
*/

// NumArray хранит только массив префиксных сумм prefix.
type NumArray struct {
	prefix []int
}

/*
make([]int, len(nums)+1) создаёт срез длины n+1. prefix[0] устанавливается по умолчанию в 0.
Цикл: на i-й итерации мы заполняем prefix[i+1] — сумму элементов nums[0]..nums[i] (т.е. добавляем nums[i] к предыдущей сумме prefix[i]).
В результате prefix[k] = сумма первых k элементов nums (если считать от 0: prefix[0]=0, prefix[1]=nums[0], ...).
*/

func Constructor(nums []int) NumArray {
	prefix := make([]int, len(nums)+1) // +1, чтобы удобно считать
	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}
	return NumArray{prefix}
}

/*
Возвращаем разницу между суммой первых right+1 элементов и суммой первых left элементов → получаем сумму nums[left..right].
Индексы безопасны, потому что prefix длины len(nums)+1, а right+1 <= len(nums).
*/

func (n *NumArray) SumRange(left int, right int) int {
	return n.prefix[right+1] - n.prefix[left]
}

/*
724. Find Pivot Index.
Нужно найти индекс поворота (pivot index) массива — то есть такой индекс i, что:
сумма элементов слева == сумма элементов справа
Если таких индексов несколько → вернуть левый самый первый.
Если ни одного → вернуть -1.

nums = [1,7,3,6,5,6]

слева от индекса 3 (6): сумма = 1+7+3=11
справа: сумма = 5+6=11
Ответ: 3

time: O(n) - собираем массив из всех элементов
memory: O(1) - один массив на входе
*/

func PivotIndex(nums []int) int {
	totalSum := 0
	for _, v := range nums {
		totalSum += v
	}

	leftSum := 0
	for i, v := range nums {
		// проверяем условие pivot
		if leftSum == totalSum-leftSum-v {
			return i
		}
		leftSum += v
	}

	return -1
}

/*
560. Subarray Sum Equals K
нужно найти количество подмассивов, сумма которых равна k.
Внутри массива подмассивы не прерываются, элементы идут друг за другом.

nums = [1,1,1], k = 2
Ответ = 2  (подмассивы [1,1] начиная с индекса 0 и с индекса 1)

nums = [1,2,3], k = 3
Ответ = 2  (подмассивы [1,2], [3])

Эффективное решение (O(n)) — через префиксные суммы + map
Идея:

Пусть prefix[i] = сумма nums[0..i].
Подмассив nums[l..r] имеет сумму k, если:
prefix[r] - prefix[l-1] = k

Значит, если мы идём по массиву и считаем prefixSum,
то надо проверить, сколько раз встречался prefixSum - k раньше.

time: O(n) - делаем один проход по массиву nums → O(n)
memory: O(1) - хранит все возможные уникальные prefixSum
*/

func SubarraySum(nums []int, k int) int {
	count := 0
	prefixSum := 0
	freq := map[int]int{0: 1} // базовый случай: prefixSum = 0 встречается 1 раз

	for _, num := range nums {
		prefixSum += num

		// проверяем, есть ли подходящий prefixSum - k
		if val, ok := freq[prefixSum-k]; ok {
			count += val
		}

		// обновляем частоту текущего prefixSum
		freq[prefixSum]++
	}

	return count
}

/*
304. Range Sum Query 2D - Immutable
у нас есть матрица m x n, и нужно быстро отвечать на запрос:
sumRegion(row1, col1, row2, col2)
= сумма элементов подматрицы от (row1,col1) до (row2,col2) включительно.

Идея решения:

Чтобы не считать сумму каждый раз с нуля, используем 2D-префиксные суммы.
Определим prefix[r][c] = сумма элементов прямоугольника от (0,0) до (r-1,c-1) (верхний левый угол включительно).
prefix[r][c] = matrix[r-1][c-1] + prefix[r-1][c] + prefix[r][c-1] - prefix[r-1][c-1]
*/

type NumMatrix struct {
	prefix [][]int
}

func PrefixMatrix(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{[][]int{}}
	}

	m, n := len(matrix), len(matrix[0])
	prefix := make([][]int, m+1)
	for i := range prefix {
		prefix[i] = make([]int, n+1)
	}

	for r := 1; r <= m; r++ {
		for c := 1; c <= n; c++ {
			prefix[r][c] = matrix[r-1][c-1] +
				prefix[r-1][c] +
				prefix[r][c-1] -
				prefix[r-1][c-1]
		}
	}

	return NumMatrix{prefix}
}

func (n *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return n.prefix[row2+1][col2+1] -
		n.prefix[row1][col2+1] -
		n.prefix[row2+1][col1] +
		n.prefix[row1][col1]
}

func main() {
	// Пример решение задачи 303. Range Sum Query - Immutable
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)

	fmt.Println("303. Range Sum Query - Immutable")
	fmt.Println(obj.SumRange(0, 2)) // 1  => (-2 + 0 + 3)
	fmt.Println(obj.SumRange(2, 5)) // -1 => (3 - 5 + 2 - 1)
	fmt.Println(obj.SumRange(0, 5)) // -3 => (-2 + 0 + 3 - 5 + 2 - 1)

	// Пример решение задачи 724. Find Pivot Index.
	fmt.Println("724. Find Pivot Index")
	fmt.Println(PivotIndex([]int{1, 7, 3, 6, 5, 6})) // 3
	fmt.Println(PivotIndex([]int{1, 2, 3}))          // -1
	fmt.Println(PivotIndex([]int{2, 1, -1}))         // 0

	// Пример решения задачи 560. Subarray Sum Equals K
	fmt.Println("560. Subarray Sum Equals K")
	fmt.Println(SubarraySum([]int{1, 1, 1}, 2)) // 2
	fmt.Println(SubarraySum([]int{1, 2, 3}, 3)) // 2
	fmt.Println(SubarraySum([]int{1}, 0))       // 0

	// Пример решения задачи 304. Range Sum Query 2D - Immutable
	fmt.Println("304. Range Sum Query 2D - Immutable")
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}

	prefixMatrix := PrefixMatrix(matrix)
	fmt.Println(prefixMatrix.SumRegion(2, 1, 4, 3)) // 8
	fmt.Println(prefixMatrix.SumRegion(1, 1, 2, 2)) // 11
	fmt.Println(prefixMatrix.SumRegion(1, 2, 2, 4)) // 12
}
