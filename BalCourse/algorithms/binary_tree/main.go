package main

import "fmt"

/*
time: O(n) - нам нужно обойти все узлы дерева обычно
memory: O(h) - где h это высота дерева, в худшем случае h = n,
поэтому может быть по памяти и O(n)

По памяти для обхода дерева нам нужно создать стек.
*/

/*
В бинарном дереве у одного узла не может быть больше двух детей.
Узел - это, то из чего состоит дерево (ещё называют нода или вершина).
Лист - это узел (нода) у которых нет детей.
Корень - это узел (нода) у которого нет родителя.
*/

// Node представляет один узел дерева
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// BinaryTree представляет само дерево, хранящее ссылку на корень
type BinaryTree struct {
	Root *Node
}

/*
1. Insert(value int) — метод дерева:
Проверяет, есть ли корень (t.Root):
Если пусто — создаёт новый узел и делает его корнем.
Если не пусто — вызывает рекурсивную вставку через n.insert(value).

2. insert(value int) — метод узла:
Сравнивает value с текущим значением узла (n.Value):
Если value <= n.Value — идёт влево.
Если value > n.Value — идёт вправо.

Затем:

Если соответствующий потомок (Left или Right) пустой, вставляет новый узел туда.
Если не пустой, вызывает insert рекурсивно, чтобы продолжить поиск позиции глубже.
*/

// Insert - Метод для вставки значения в дерево
func (t *BinaryTree) Insert(value int) {
	if t.Root == nil {
		t.Root = &Node{Value: value}
	} else {
		t.Root.insert(value)
	}
}

// Рекурсивная вставка в узел
func (n *Node) insert(value int) {
	if value <= n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.insert(value)
		}
	}
}

/*
Pre-order traversal (Прямой обход):
Посещаем текущий узел → сохраняем его значение.
Обходим левое поддерево рекурсивно.
Обходим правое поддерево рекурсивно.

По сути все обходы имеют действие, в зависимости от того, где будет действие и есть разные способы обхода дерева.
Действие - это вывод вершины в консоль например, или добавление в массив и т.д.

Действие
Влево
Вправо

Влево
Действие
Вправо

Влево
Вправо
Действие
*/

// PreOrder - Метод дерева: возвращает pre-order (корень -> лево -> право) обход в виде среза
func (t *BinaryTree) PreOrder() []int {
	return preOrderHelper(t.Root)
}

// Рекурсивная функция обхода
func preOrderHelper(n *Node) []int {
	if n == nil {
		return []int{}
	}

	result := []int{n.Value}                            // 1. Посетить текущий узел
	result = append(result, preOrderHelper(n.Left)...)  // 2. Обойти левое поддерево
	result = append(result, preOrderHelper(n.Right)...) // 3. Обойти правое поддерево
	return result
}

// InOrder возвращает значения в симметричном порядке (лево → корень → право)
func (t *BinaryTree) InOrder() []int {
	return inOrderHelper(t.Root)
}

func inOrderHelper(n *Node) []int {
	if n == nil {
		return []int{}
	}

	result := inOrderHelper(n.Left)                    // 1. Обход левого поддерева
	result = append(result, n.Value)                   // 2. Посещение текущего узла
	result = append(result, inOrderHelper(n.Right)...) // 3. Обход правого поддерева
	return result
}

// PostOrder возвращает значения в обратном порядке (лево → право → корень)
func (t *BinaryTree) PostOrder() []int {
	return postOrderHelper(t.Root)
}

func postOrderHelper(n *Node) []int {
	if n == nil {
		return []int{}
	}

	result := postOrderHelper(n.Left)                    // 1. Обход левого поддерева
	result = append(result, postOrderHelper(n.Right)...) // 2. Обход правого поддерева
	result = append(result, n.Value)                     // 3. Посещение текущего узла
	return result
}

/*
Задача 102. Binary Tree Level Order Traversal с LeetCode
Нужно вернуть уровневый обход бинарного дерева (от корня к листьям, слева направо).
time: O(n) - где n число вершин
memory: O(h) - храним все узлы

Как работает
dfs принимает текущий узел, уровень и ссылку на результат.
Если уровень ещё не встречался — добавляем пустой список.
Кладём значение узла в соответствующий уровень.
Рекурсивно вызываем dfs для Left и Right.
*/

func (t *BinaryTree) LevelOrder() [][]int {
	result := [][]int{}
	dfs(t.Root, 0, &result)
	return result
}

func dfs(n *Node, level int, result *[][]int) {
	if n == nil {
		return
	}

	// Если уровень ещё не создан — создаём новый срез
	if len(*result) == level {
		*result = append(*result, []int{})
	}

	// Добавляем значение в соответствующий уровень
	(*result)[level] = append((*result)[level], n.Value)

	// PreOrder: обходим левое поддерево, потом правое
	dfs(n.Left, level+1, result)
	dfs(n.Right, level+1, result)
}

/*
199. Binary Tree Right Side View:
нужно вернуть список значений узлов, которые видны при взгляде на бинарное дерево справа.
Решение похоже на Level Order
time: O(n) - где n число вершин
memory: O(h) - храним все узлы
*/

func (t *BinaryTree) RightSideViewDFS() []int {
	result := []int{}
	dfsRightSide(t.Root, 0, &result)
	return result
}

func dfsRightSide(n *Node, level int, result *[]int) {
	if n == nil {
		return
	}

	if len(*result) == level {
		*result = append(*result, n.Value) // первый узел на уровне (справа налево)
	}

	dfsRightSide(n.Right, level+1, result)
	dfsRightSide(n.Left, level+1, result)
}

/*
Задача 101. Symmetric Tree: нужно проверить, является ли бинарное дерево зеркально симметричным.
time: O(n) - где n число вершин
memory: O(h) - храним все узлы
*/

func (t *BinaryTree) IsSymmetric() bool {
	if t.Root == nil {
		return true
	}
	return isMirror(t.Root.Left, t.Root.Right)
}

func isMirror(t1, t2 *Node) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Value == t2.Value &&
		isMirror(t1.Left, t2.Right) &&
		isMirror(t1.Right, t2.Left)
}

/*
Задача 100. Same Tree: нужно проверить, идентичны ли два бинарных дерева (структура + значения).
time: O(n) - где n число вершин
memory: O(2h) - храним все узлы
*/

func IsSameTree(p *Node, q *Node) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Value != q.Value {
		return false
	}
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

/*
112. Path Sum: нужно проверить, существует ли в бинарном дереве путь от корня до листа,
сумма значений которого равна targetSum.
time: O(n) - где n число вершин
memory: O(h) - храним все узлы
*/

func HasPathSum(root *Node, targetSum int) bool {
	if root == nil {
		return false
	}

	// если лист → проверяем сумму
	if root.Left == nil && root.Right == nil {
		return root.Value == targetSum
	}

	// рекурсивно проверяем левое и правое поддерево
	remaining := targetSum - root.Value
	return HasPathSum(root.Left, remaining) || HasPathSum(root.Right, remaining)
}

func main() {
	tree := &BinaryTree{}

	// Создаём узлы
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n4 := &Node{Value: 4}
	n5 := &Node{Value: 5}
	n6 := &Node{Value: 6}
	n7 := &Node{Value: 7}

	// Строим дерево вручную
	n1.Left = n2
	n1.Right = n3

	n2.Left = n4
	n2.Right = n5

	n3.Left = n6
	n3.Right = n7

	tree.Root = n1

	/*

	       1
	     /   \
	    2     3
	   / \   / \
	  4  5  6   7

	*/

	fmt.Println("PreOrder:  ", tree.PreOrder())  // [1 2 3 4 5 6 7]
	fmt.Println("InOrder:   ", tree.InOrder())   // [4 2 5 1 6 3 7]
	fmt.Println("PostOrder: ", tree.PostOrder()) // [4 5 2 6 7 3 1]

	// Пример решение задачи 102. Binary Tree Level Order Traversal с LeetCode
	treeLevel := &BinaryTree{}

	root := &Node{Value: 3}
	root.Left = &Node{Value: 9}
	root.Right = &Node{Value: 20}
	root.Right.Left = &Node{Value: 15}
	root.Right.Right = &Node{Value: 7}

	treeLevel.Root = root

	/*
	    3
	   / \
	  9  20
	    /  \
	   15   7

	*/

	fmt.Println("LevelOrder: ", treeLevel.LevelOrder()) // [[3], [9,20], [15,7]]

	// Пример решения задачи 199. Binary Tree Right Side View
	treeRightSide := &BinaryTree{}

	rootRight := &Node{Value: 1}
	rootRight.Left = &Node{Value: 2}
	rootRight.Right = &Node{Value: 3}
	rootRight.Left.Right = &Node{Value: 5}
	rootRight.Right.Right = &Node{Value: 4}

	treeRightSide.Root = rootRight

	/*

	    1
	   / \
	  2   3
	   \    \
	    5    4

	*/

	fmt.Println("RightSideView: ", treeRightSide.RightSideViewDFS()) // [1 3 4]

	// Пример решения Задача 101. Symmetric Tree
	treeSymmetric := &BinaryTree{}

	rootSymmetric := &Node{Value: 1}
	rootSymmetric.Left = &Node{Value: 2}
	rootSymmetric.Right = &Node{Value: 2}
	rootSymmetric.Left.Left = &Node{Value: 3}
	rootSymmetric.Left.Right = &Node{Value: 4}
	rootSymmetric.Right.Left = &Node{Value: 4}
	rootSymmetric.Right.Right = &Node{Value: 3}

	treeSymmetric.Root = rootSymmetric

	/*
			    1
			   / \
			  2   2
			 / \ / \
			3  4 4  3

		out: true

		    1
		   / \
		  2   2
		   \   \
		   3    3

		out: false
	*/

	fmt.Println("SymmetricTree: ", treeSymmetric.IsSymmetric()) // [true]

	// Пример решение задачи 100. Same Tree
	p := &Node{Value: 1}
	p.Left = &Node{Value: 2}
	p.Right = &Node{Value: 3}

	q := &Node{Value: 1}
	q.Left = &Node{Value: 2}
	q.Right = &Node{Value: 3}

	/*
			p:    1        q:    1
			     / \            / \
			    2   3          2   3

			out: true

		p:    1        q:    1
		     /              \
		    2                2

		out: false
	*/

	fmt.Println("SameTree: ", IsSameTree(p, q)) // [true]

	// Пример решение задачи 112. Path Sum
	rootPathSum := &Node{Value: 5}
	rootPathSum.Left = &Node{Value: 4}
	rootPathSum.Right = &Node{Value: 8}
	rootPathSum.Left.Left = &Node{Value: 11}
	rootPathSum.Left.Left.Left = &Node{Value: 7}
	rootPathSum.Left.Left.Right = &Node{Value: 2}
	rootPathSum.Right.Left = &Node{Value: 13}
	rootPathSum.Right.Right = &Node{Value: 4}
	rootPathSum.Right.Right.Right = &Node{Value: 1}

	/*
				    5
				   / \
				  4   8
				 /   / \
				11  13  4
				/ \       \
				7   2       1

		targetSum = 22
		Выход: true (путь 5 → 4 → 11 → 2).
	*/

	fmt.Println("PathSum: ", HasPathSum(rootPathSum, 22))
}
