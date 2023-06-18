package computation

import "fmt"

func Fact(n int) int {

	fmt.Println("Функция из пакета computation")
	var result = 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

/*
Код файла factorial.go принадлежит пакету computation.
Важно отметить, что название функции начинается с заглавной буквы.
Благодаря этому данная функция будет видна в других пакетах.

И чтобы использовать функцию factorial, надо импортировать этот пакет в файле main.go:
*/
