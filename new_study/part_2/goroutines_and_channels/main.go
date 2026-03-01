package main

import (
	"fmt"
	"sync"
	"time"
)

// Поход рабочего в шахту
func mine(n int) int {
	fmt.Println("Начал поход в шахту: ", n)
	time.Sleep(1 * time.Second)
	fmt.Println("Завершен поход в шахту: ", n)

	return n
}

/*
func main() {
	// Количество угля
	coal := 0
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	initTime := time.Now()

	// wg.Add(1) до запуска горутины, если внутри горутины,
	// то она может не успеть запуститься и wg.Wait() её ждать не будет
	// Когда мы разделяем получение результата result и сложение coal с блокировкой, то это быстрей,
	// чем блокировать вызов функции и сложение вместе
	wg.Add(1)
	go func() {
		defer wg.Done()

		result := mine(1)
		mu.Lock()
		coal += result
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		result := mine(2)
		mu.Lock()
		coal += result
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		result := mine(3)
		mu.Lock()
		coal += result
		mu.Unlock()
	}()

	wg.Wait()

	fmt.Println("Добыли угля: ", coal)
	fmt.Println("Прошло времени: ", time.Since(initTime))
}
*/

/*
	РЕШЕНИЕ ЧЕРЕЗ КАНАЛЫ
	Важно понимать, что главная горутина main блокируется на этапе чтения из канала в цикле
	до тех пор пока канал не закрыт, это фишка for range при работе с каналами
func main() {
	ch := make(chan int, 3)
	var wg sync.WaitGroup
	start := time.Now()

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			ch <- mine(n)
		}(i)
	}

	// Горутина координатор, которая ждет завершения всех горутин и закрывает канал
	go func() {
		wg.Wait()
		close(ch)
	}()

	coal := 0
	for result := range ch {
		coal += result
	}

	fmt.Println("Добыли угля: ", coal)
	fmt.Println("Прошло времени: ", time.Since(start))
}
*/

// РЕШЕНИЕ КАНАЛ КАК MUTEX
// Важно, если мы в рамках одной горутины и пишем в канал и читаем из канала, то это должен быть буферизированный канал,
// иначе получим блокировку. Пока в буфере есть место мы можем в одной горутине и писать и читать

// Небуферизированный канал должен иметь писателей и читателей, которые работают в РАЗНЫХ горутинах !!!
func main() {
	ch := make(chan struct{}, 1)
	var wg sync.WaitGroup
	start := time.Now()

	coal := 0

	wg.Add(1)
	go func() {
		defer wg.Done()
		result := mine(1)

		ch <- struct{}{}
		coal += result
		<-ch
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		result := mine(2)

		ch <- struct{}{}
		coal += result
		<-ch
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		result := mine(3)

		ch <- struct{}{}
		coal += result
		<-ch
	}()

	wg.Wait()

	fmt.Println("Добыли угля: ", coal)
	fmt.Println("Прошло времени: ", time.Since(start))
}

// АКСИОМЫ КАНАЛОВ

/*
				nil channel		closed channel

close(ch)			panic			panic

read				block			default value
val: <- ch

write				block			panic
ch <- val
*/

// ЧТЕНИЕ ИЗ КАНАЛА В ЦИКДЕ ПОД КАПОТОМ
/*
for {
    value, ok := <-ch  // Пытаемся прочитать из канала
    if !ok {            // Если канал закрыт и пуст
        break           // Выходим из цикла
    }
    fmt.Println(value)  // Обрабатываем значение
}
*/
