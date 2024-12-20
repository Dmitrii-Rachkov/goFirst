package main

import (
	"fmt"
	"runtime"
)

func main() {

	// defer функции (отложенные) складываются в специальный stack и выполняются в конце программы (функции main)
	// после return
	defer fmt.Println(1)
	defer fmt.Println(2)

	// С помощью defer можно изменять именованные возвращаемые значения
	fmt.Println(sum(2, 3)) // увидим здесь 10 вместо 5

	// Пример с defer в циклах
	deferValues()

	// --#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--

	// Так создаём простую горутину с функцией showNumbers()
	// Если запустить функцию без горутины (go), то по порядку выведутся числа от 0 до 99
	// Если с горутиной, то результат может быть:
	/*
		0
		exit
		1
		2
		3
	*/
	// Для наглядности добавьте time.Sleep
	go showNumbers(100)
	// time.Sleep(20 * time.Microsecond)

	// Такой вывод мы видим, потому что цель планировщика go как можно быстрей завершить программу, в данном случае
	// функцию main.
	// Вот здесь go showNumbers(100), при вызове функции создаётся новая горутина, при этом горутина main
	// продолжает выполняться. Т.к. мы не синхронизируем работу наших горутин, то как только завершается работа
	// нашей основной горутины main, наша программа завершается.
	// Поэтому горутина нужно как-то синхронизировать.
	// Если мы запускаем одну горутину внутри другой и не синхронизируем их, то как только выполнится основная горутина
	// мы не дождемся результата других функции.

	// --#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--

	// Так мы можем посмотреть количество логических ядер процессора
	fmt.Println(runtime.NumCPU())

	// Так можем задать количество логических ядер для работы программы
	//runtime.GOMAXPROCS(1)

	// Позволяет в ручную переключаться между горутинами
	//runtime.Gosched()

	// Если поставим задержку 1 секунду, то планировщик Go переключиться с одной горутины на другую
	//time.Sleep(time.Second)

	fmt.Println("exit")

	// --#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--#--

	// Функция для паники
	makePanic()
}

func showNumbers(num int) {
	for i := 0; i < num; i++ {
		fmt.Println(i)
	}
}

func sum(x, y int) (sum int) {
	defer func() {
		sum *= 2
	}()

	sum = x + y
	return
}

func deferValues() {
	// Здесь распечатается first и числа с 9 до 0
	// т.е. в обратном порядке
	// Аргументы вычисляются сразу и складываются в defer stack
	for i := 0; i < 10; i++ {
		defer fmt.Println("first", i)
	}

	// Выводятся одинаковые значения second 10
	// На момент вызова defer func() значение i = 10
	// Так делать не надо
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println("second", i)
		}()
	}

	// Предыдущее неправильное поведение можно исправить
	// путём создания локальной переменной k
	for i := 0; i < 10; i++ {
		k := i
		defer func() {
			fmt.Println("third", k)
		}()
	}

	// Или добавить параметр i в качестве аргумента
	for i := 0; i < 10; i++ {
		defer func(k int) {
			fmt.Println("fourth", k)
		}(i)
	}
}

// Panic - это особая ситуация в Go, когда наша программа говорит о том, произошло что-то неожиданное
// и программа работать не сможет и аварийно завершит свою работу
func makePanic() {
	// Конструкция с recover, чтобы ловить панику
	// результат функции recover это значение паники
	defer func() {
		panicValue := recover()
		fmt.Println(panicValue)
	}()

	// Можем положить сюда любое значение
	panic("some panic")

	// После паники следующий код не работает
	fmt.Println("Unreachable code")
}
