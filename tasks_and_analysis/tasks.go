package tasks_and_analysis

import (
	"fmt"
	"sync"
)

// Задача 1
func task1() {
	test := "test"
	for _, symbol := range test {
		fmt.Printf("%s", symbol)
	}
}

/*
Переменная test содержит тип данных string, а строки в Go это последовательность байтов в кодировке UTF-8. Оператор range не возвращает индекс строки "_" и возвращает символ в строке "symbol".

Символ представляет собой тип данных rune, который соответствует Unicode-символу. Мы пытаемся вывести "symbol" как строку fmt.Printf("%s", symbol), но это не работает потому, что другой тип данных.

Тип данных rune в Go это alias int32, поэтому чтобы исправить код, нужно "symbol" преобразовать к строке fmt.Printf("%s", string(symbol)) или воспользоваться специальным символом для вывода fmt.Printf("%c", symbol).
*/

// Задача 2
func task2() {
	var test map[string]int
	test["test1"] = 1
}

/*
Map в Go представляет собой ссылку на хэш-таблицу, и в данном примере переменная test имеет тип данных map. Однако переменная test не инициализирована (не выделена память под неё) и по умолчанию равна nil.

При присвоении значения по ключу test["test1"] = 1 в неинициализированную map возникает panica. Поэтому сначала необходимо инициализировать map и выделить под неё память, например с помощью test := make(map[string]int) или с помощью пустого литерала test := map[string]int{}.
*/

// Задача 3
func task3() {
	test := map[string]int{"test1": 1, "test2": 2, "test3": 3}
	i := 0
	for key, _ := range test {
		if i == 0 && key == "test1" {
			fmt.Println("test1 on the first place")
		}
		i++
	}
}

/*
Порядок элементов в map неопределен, это значит, что код не гарантирует, что "test1" будет первым элементом. В данном случае если "test1" будет первым элементов, то выведется сообщение "test1 on the first place", а иначе ничего не выведется.
*/

// Задача 4
func task4() {
	test := map[string]int{"test1": 1, "test2": 2, "test3": 3}

	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		defer wg.Done()
		test["test1"] = 111
	}()

	go func() {
		defer wg.Done()
		test["test1"] = 112
	}()

	go func() {
		defer wg.Done()
		test["test1"] = 113
	}()

	go func() {
		defer wg.Done()
		test["test1"] = 114
	}()

	go func() {
		defer wg.Done()
		test["test1"] = 115
	}()

	wg.Wait()
	fmt.Println(test)
}

/*
Так и так, в данном примере будет возникать гонка данных (data race), что приведёт к неопределенному значению ключа "test1".

С помощью var wg sync.WaitGroup в примере объединены 5 горутин, соответственно пока вся группа не выполнится программа будет заблокирована.

В примере все 5 горутин меняют значение ключа "test1", соответственно горутины перезаписывают информацию друг друга, это плохо.

В Go map не является потокобезопасной (thread-safe), это значит что несколько горутин работающих с одной мапой приводят к гонке данных.

Гонка данных — это ситуация, когда несколько потоков пытаются читать или писать в одни и те же данные без надлежащей синхронизации. Данные могут перезаписываться или работа программы приведёт к панике.

Для предотвращения гонки данных можно воспользоваться синхронизацией данных с помощью sync.Mutex, которая будет блокировать доступ к мапе другим горутинам, пока идёт работа в одной горутине, а затем когда работа завершена - разблокировать доступ к map.

func main() {
    test := map[string]int{"test1": 1, "test2": 2, "test3": 3}

    var wg sync.WaitGroup
    wg.Add(5)

    // Мьютекс для защиты карты
    var mu sync.Mutex

    go func() {
       defer wg.Done()

       // Блокируем доступ к map
       mu.Lock()

       test["test1"] = 111

       // Разблокируем доступ к map
       mu.Unlock()
    }()

    go func() {
       defer wg.Done()
       mu.Lock()
       test["test1"] = 112
       mu.Unlock()
    }()

    go func() {
       defer wg.Done()
       mu.Lock()
       test["test1"] = 113
       mu.Unlock()
    }()

    go func() {
       defer wg.Done()
       mu.Lock()
       test["test1"] = 114
       mu.Unlock()
    }()

    go func() {
       defer wg.Done()
       mu.Lock()
       test["test1"] = 115
       mu.Unlock()
    }()

    wg.Wait()
    fmt.Println(test)
}
*/

// Задача 5
func task5() {
	test := map[string]int{"test1": 1, "test2": 2, "test3": 3}
	testFunc(test)
}

func testFunc(test map[string]int) {
	for key, value := range test {
		fmt.Printf("%s -> %d", key, value)
	}
}

/*
В языке Go передача map в другие функции осуществляется по указателю. Если бы мы в функции testFunc изменили бы map, то изменилась и исходная map test. Но в данном примере у нас в случайном порядке будет выводится ключ и значение в консоль, т.к. порядок итерации range не указан и не гарантируется что он будет одинаковым от запуска к запуску.
*/

// Задача 6
func task6() {
	var c chan int
	c <- 1
	fmt.Println(<-c)
}

/*
Каналы в языке Go необходимы для обмена данными между горутинами. Также как и map, каналы являются указателями на структуру, поэтому по умолчанию канал равен nil.

В нашем кейсе переменная "c" является каналом, однако она не инициализирована и под неё не выделена память. Соответственно при отправке/получении данных канал должен быть проинициализирован, а в нашем случае c <- 1 мы получим ошибку fatal error.

Чтобы это исправить нужно проинициализировать канал c := make(chan int).

Следующая проблема в том, что у нас есть функция main и по сути это единственная горутина в программе, а каналы нужны для обмена данными между минимум двумя горутинами.

В нашем случае когда мы пытаемся отправить данные в канал c <- 1 операция блокирует выполнение программы. Если канал пустой, то горутина-получатель блокируется, пока в канале не окажутся данные.

Когда мы пытаемся получить данные из канала <-c программа блокируется, потому что канал пуст и нет других горутин, которые могут отправить в него данные.

Исправленный код:
func main() {
    // Инициализируем канал
    c := make(chan int)

    // Горутина для отправки значения в канал
    go func() {
       c <- 1
    }()

    // Получаем данные из канала
    fmt.Println(<-c)
}

Таким образом на экране увидем 1.


*/

// Задача 7
func tasks7() {
	test := make(chan int)
	go func(test chan int) {
		<-test
	}(test)
	test <- 123
	test <- 456
	fmt.Println("OK")
}

/*
В примере создан небуферизированный канал test := make(chan int), это значит что для отправки значения в канал нужно чтобы кто-то его получил, иначе запись в канал будет блокировать выполнение программы.

Строка go func(test chan int) { <-test }(test) запускает горутину которая пытается принять значение из канала test. Но канал небуферизированный горутина заблокируется на операции <-test, ожидая поступления данных.

Основная горутина main пытается отправить данные в канал test <- 123, но так как горутина анонимной функции уже заблокирована на чтение из канала мы не можем отправить в канал значение и основная горутина main блокируется.

Строки test <- 456 и fmt.Println("OK") не выполнятся так как основная горутина main заблокирована.

Можно исправить это с помощью буферизированного канала:
func main() {
    test := make(chan int, 2)
    go func(test chan int) {
       <-test
    }(test)
    test <- 123
    test <- 456
    fmt.Println("OK")
}
*/

// Задача 8
/*
func task8() {
	test := make(chan int)
	go func(test chan<- int) {
		<-test
	}(test)
	test <- 123
	fmt.Println("OK")
}

В этом примере будет ошибка компиляции, т.к. мы создали канал test с типом на отправку/получение данных, а в горутине в анонимной функции мы ожидаем на вход канал в который можно только отправлять данные - это неверно. И также мы пытаемся получить данные из канала test внутри самой <-test горутины, хотя на вход анонимной функции мы ожидаем получить только канал на отправку данных.

Можно исправить например так:
func main() {
    test := make(chan int)
    go func(test chan int) {
       <-test
    }(test)
    test <- 123
    fmt.Println("OK")
}
*/

// Задача 9
func task9() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}
	for r := 0; r < 5; r++ {
		go func(r int) {
			mutex.Lock()
			state[r] = r + 1
			mutex.Unlock()

		}(r)
	}
	fmt.Println(state)
}

/*
В данном примере инициализируется map - var state = make(map[int]int) и создаётся mutex который блокирует доступ к map state в горутинах.

С помощью цикла создаётся 5 горутин, которые добавляют в map state пару ключ-значения, при этом благодаря mutex нет data race.

В конце выводится map state, но проблема в том, что вывод непредсказуем. Наша главная горутина main может завершить программу не дождавшись, пока выполнятся все 5 горутин.

Нужно синхронизировать работу 5 горутин с главной горутиной main. Например можно так:
func main() {
    var state = make(map[int]int)
    var mutex = &sync.Mutex{}

    // Добавляем wg для ожидания завершения всех горутин
    wg := &sync.WaitGroup{}
    for r := 0; r < 5; r++ {
       // Устанавливаем счётчик горутин 1
       wg.Add(1)
       go func(r int) {
          // В конце выполнения горутины даём сигнал о завершении
          defer wg.Done()

          mutex.Lock()
          state[r] = r + 1
          mutex.Unlock()
       }(r)
    }

    // Ожидаем завершения всех горутин
    wg.Wait()

    fmt.Println(state)
}
*/
