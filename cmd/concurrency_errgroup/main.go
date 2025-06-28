package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	// Канал как promise (обещание)
	// chanAsPromise()

	// Канал как mutex
	// chanAsMutex()

	// Пример без ErrGroup
	// withoutErrGroup()

	// Пример ErrGroup
	errGroup()
}

// Функция makeRequest принимает число, просто для различия запросов.
// Из функции возвращаем канал из которого можно читать данные.
// Таким образом имитируем http запрос.
func makeRequest(num int) <-chan string {
	responseChan := make(chan string)

	go func() {
		time.Sleep(time.Second)
		responseChan <- fmt.Sprintf("response number %d", num)
	}()
	return responseChan
}

// Предположим нам нужно сделать два запроса и мы не хотим блокироваться, ожидая от них ответа,
// мы хотим делать код дальше, и потом в конце выполнения программы, когда нам эти данные из
// двух запросов понадобятся, мы их попытаемся получить.
// Для этого реализована функция makeRequest и в функции chanAsPromise мы не блокируемся и ничего не ждём
func chanAsPromise() {
	firstResponseChan := makeRequest(1)
	secondResponseChan := makeRequest(2)
	// do something else
	fmt.Println("non blocking")

	fmt.Println(<-firstResponseChan, <-secondResponseChan)
}

// Есть счётчик counter, мы хотим запустить 1000 горутин и каждая горутина увеличивает счётчик.
// Основная проблема в том, что несколько горутин берут одно и тоже значение counter и пытаются туда записать.
// Мы потеряем часть записей.
func chanAsMutex() {
	// Счётчик
	var counter int

	// Создаём буферизированный канал mutexChan
	mutexChan := make(chan struct{}, 1)
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// Размер структуры определяется количеством полей,
			// если полей нет - пустая структура ничего не весит
			// Пытаемся записать в канал и при этом мы блокируемся так как ёмкость канала 1,
			// Мы пишем один раз и блокируемся, ждём пока освободится канал <-mutexChan
			mutexChan <- struct{}{}

			counter++

			<-mutexChan
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

// В примере ниже мы запускаем три горутины и если в друг в одной из них возникнет ошибка
// мы хотим получить эту ошибку и при этом прекратить выполнение всех остальных горутин.
func withoutErrGroup() {
	var err error
	// Создаём контекст который можно отменить
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	wg.Add(3)

	// Первая горутина спит 1 секунду и смотрит если контекст закрыт, то ничего не делает,
	// если иначе, то пишем в консоль сообщение и спим 1 секунду.
	go func() {
		time.Sleep(time.Second)
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("first started")
			time.Sleep(time.Second)
		}
	}()

	// Вторая горутина если контекст завершен, то ничего не делает
	// иначе пишет сообщение в консоль, пишет ошибку и закрывает контекст
	go func() {
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("second started")
			err = fmt.Errorf("any error")
			cancel()
		}
	}()

	// Третья горутина если контекст закрыт, то ничего не делает, иначе
	// пишет сообщение в консоль и спит 1 секунду
	go func() {
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("third started")
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	fmt.Println(err)
}

// errGroup следит за тем, чтобы ошибку мы получили только самую первую и не было race condition.
// errGroup позволяет отменять задачу, если отменился контекст при ошибке и соответственно
// все последующие горутины не запустятся.
func errGroup() {
	// Вызываем в пакете ErrGroup в которой есть контекст
	// В этом пакете есть метод Wait который ждёт завершения всех горутин и также
	// метод Go который может принимать callback который возвращает ошибку
	g, ctx := errgroup.WithContext(context.Background())

	// Первая горутина также если контекст завершился, то ничего не делаем и возвращаем пустую ошибку
	// иначе пишем сообщение в консоль и спим 1 секунду
	g.Go(func() error {
		time.Sleep(time.Second)

		select {
		case <-ctx.Done():
			return nil
		default:
			fmt.Println("first started")
			time.Sleep(time.Second)
			return nil
		}
	})

	// Вторая горутина возвращает ошибку
	g.Go(func() error {
		fmt.Println("second started")
		return fmt.Errorf("unexpected error in request 2")
	})

	// Третья горутина также если контекст завершился, то ничего не делаем и возвращаем пустую ошибку
	// иначе пишем сообщение в консоль и спим 1 секунду.
	// Return можно писать в каждом case, а можно писать один раз после всех case
	g.Go(func() error {
		select {
		case <-ctx.Done():
		default:
			fmt.Println("third started")
			time.Sleep(time.Second)
		}
		return nil
	})

	// Вызываем метод Wait который возвращает ошибку
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
