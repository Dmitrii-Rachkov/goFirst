package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// baseKnowledge()
	workerPool()
}

// Context - это объект, который служит для двух целей:
// 1. Хранить какие-то значения, чтобы передавать их дальше в другие функции или горутины
// 2. Сообщать о завершении
func baseKnowledge() {
	// На старте можно создать Background и TODO
	// Смысл в том, что изначально мы создаём контекст где-то наверху в корне нашего приложения
	// например приходит http запрос и там у нас формируется контекст и он передаётся дальше по флоу.
	// Также можем в запросе положить какие-то значения

	// Создаём корневой контекст
	ctx := context.Background()
	fmt.Println(ctx)

	// Этот контекст который пригоден только в тестах или вы проектируйте какой-то метод или функцию
	// и хотите заложить туда контекст, но не знайте понадобится он там или нет
	toDo := context.TODO()
	fmt.Println(toDo)

	// Можем положить значение в контекст
	withValue := context.WithValue(ctx, "name", "vasya")
	fmt.Println(withValue.Value("name"))

	// Контекст который умеет завершаться
	// CancelFunc это функция которая может отменить наш контекст
	// Контекст нужно закрывать там где его создавали, на том же уровне
	withCancel, cancel := context.WithCancel(ctx)
	fmt.Println(withCancel.Err())
	cancel()
	fmt.Println(withCancel.Err())

	// Контекс которому можно устанавливать deadline
	// Передаем в функцию время, когда контекст должен завершиться
	withDeadline, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
	defer cancel()
	fmt.Println(withDeadline.Deadline())
	fmt.Println(withDeadline.Err())
	fmt.Println(<-withDeadline.Done())

	// Таймут, задаем 2 секунды и через две секунды он завершится
	withTimeout, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	fmt.Println(withTimeout.Done())
}

// Один из паттернов в разработке go
func workerPool() {

	// Создаём контекст который можно отменить
	ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel = context.WithTimeout(context.Background(), time.Millisecond*20)
	defer cancel()

	// Создаём два буферизированных канала
	wg := &sync.WaitGroup{}
	numbersToProcess, processedNumbers := make(chan int, 5), make(chan int, 5)

	// Делаем цикл столько сколько нам доступно ядер процессора
	// Смысл этого цикла запустить наши воркеры
	for i := 0; i <= runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, numbersToProcess, processedNumbers)
		}()
	}

	go func() {
		for i := 0; i < 1000; i++ {
			/*if i == 500 {
				cancel()
			}*/
			numbersToProcess <- i
		}
		close(numbersToProcess)
	}()

	go func() {
		wg.Wait()
		close(processedNumbers)
	}()

	var counter int
	for resultValue := range processedNumbers {
		counter++
		fmt.Println(resultValue)
	}

	fmt.Println(counter)
}

func worker(ctx context.Context, toProcess <-chan int, processed chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case value, ok := <-toProcess:
			if !ok {
				return
			}
			time.Sleep(time.Millisecond)
			processed <- value * value
		}
	}
}
