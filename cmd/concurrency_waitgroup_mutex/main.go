package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Wait group

	// Функция без синхронизации работы горутин
	// выводятся не все числа
	// withoutWait()

	// Функция с использованием синхронизации wg
	// выводятся все числа
	// withWait()

	// Функция для демонстрации ошибки с работой wg
	// wrongAdd()

	// Mutex

	// Функция без использования конкурентности
	// выполняется 500 000 операций
	// writeWithoutConcurrent()

	// Функция аналогичная функции ниже, только с использованием горутин и wg
	// на моей машине она работает медленней чем функция без горутин
	// и при этом выполняются все 500 000 операций, но из-за data race при обращении к счётчику
	// из разных горутин происходит перезапись
	// writeWithoutMutex()

	// Функция аналогичная выше с блокировкой доступа к счётчику
	// writeWithMutex()

	// Функция чтения и записи с mutex
	readWithMutex()

	// Функция чтения и записи с RWMutex
	readWithRWMutex()
}

// Функция, которая в цикле создаёт 10 горутин, просто распечатываем в консоль счётчик
// выходим из функции с exit сообщением. В итоге получим рандомный вывод, не все 10 значений выводятся,
// просто некоторые горутины не успевают отработать.
func withoutWait() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i + 1)
	}

	fmt.Println("exit")
}

// Функция аналогичная выше, только здесь мы используем синхронизацию wg
func withWait() {
	// Создаём переменной с wg группой горутин
	var wg sync.WaitGroup

	// Метод Add() принимает количество задач, которое нужно добавить в нашу группу
	// Если нам заранее известно количество итерации в цикле, то Add() можно вызвать до цикла
	// Если не знаем кол-во задач, то внутри цикла вызываем Add(1), до момента создания горутины !!!
	wg.Add(10)

	// Ниже в цикле мы идём от 0 до 10 и на каждую итерацию добавляем по одной задаче Add(1)
	// Также на каждую итерацию запускаем одну горутину
	for i := 0; i < 10; i++ {
		go func(i int) {
			//defer wg.Done()
			fmt.Println(i + 1)

			// Метод Done() сообщает о том, что задача выполнилась
			wg.Done()
		}(i)

	}

	// При вызове Wait() основная горутина main() блокируется и она ждёт пока в waitGroup
	// выполнятся все задачи
	wg.Wait()
	fmt.Println("exit")
}

// Здесь мы вызываем Add(1) внутри цикла, но уже в самой горутине
func wrongAdd() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {

		// Получается, если горутина успевает начать выполнение, то она добавится в wg
		// если не успеет запуститься, то не добавится в wg
		go func(i int) {
			wg.Add(1)

			defer wg.Done()

			fmt.Println(i + 1)
		}(i)

	}

	wg.Wait()
	fmt.Println("exit")
}

// Функция без конкурентности
func writeWithoutConcurrent() {
	// Получаем текущую отметку времени
	start := time.Now()

	// Счётчик равный 0
	var counter int

	// Запускаем цикл в котором есть 500 000 итераций
	// каждая итерация спит 1 наносекунду
	// далее увеличиваем счётчик на 1 ед.
	for i := 0; i < 500000; i++ {
		time.Sleep(time.Nanosecond)
		counter++
	}

	// Выводим start дату
	fmt.Println(start)

	// Выводим счётчик
	fmt.Println(counter)

	// Выводим разницу во времени между началом и завершением выполнения функции
	fmt.Println(time.Now().Sub(start).Seconds())
}

// Эта функция аналогичная функции выше, но здесь мы используем wg
// и запускаем 500 000 горутин
func writeWithoutMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup

	wg.Add(500000)

	for i := 0; i < 500000; i++ {
		go func() {
			// Удобно использовать defer, если есть if или switch
			defer wg.Done()
			time.Sleep(time.Nanosecond)
			counter++ // counter = counter + 1 // 555 + 1 = 556 // 555 + 1 = 556
		}()
	}
	wg.Wait()

	fmt.Println(start)
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

// Функция с использованием mutex, которая позволяет избежать data race
func writeWithMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(500000)

	for i := 0; i < 500000; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)

			// Блокируем доступ других горутин к счётчику
			mu.Lock()

			counter++

			// Разблокируем доступ к счётчику другим горутинам
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(start)
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

// Чтение счётчика и запись счётчика по 50 операций
// Две горутины создаём на каждой итерации цикла
func readWithMutex() {
	start := time.Now()
	var (
		counter int
		wg      sync.WaitGroup
		mu      sync.Mutex
	)

	// Общее количество операций на чтение и запись 100
	wg.Add(100)

	for i := 0; i < 50; i++ {
		// Здесь мы не меняем значение счётчика, а просто читаем его
		// однако все горутины выстраиваются в очередь из-за mutex
		go func() {
			defer wg.Done()

			// Имитируем чтение
			mu.Lock()

			time.Sleep(time.Nanosecond)
			_ = counter

			mu.Unlock()
		}()

		go func() {
			defer wg.Done()

			// Имитируем запись
			mu.Lock()

			time.Sleep(time.Nanosecond)
			counter++

			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(start)
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

// RWMutex позволяет разделить блокировку на чтение и запись
// все горутины на чтение ждут, пока разблокируется горутина на запись
func readWithRWMutex() {
	start := time.Now()
	var (
		counter int
		wg      sync.WaitGroup
		mu      sync.RWMutex
	)

	wg.Add(100)

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()

			// Блокировка на чтение
			mu.RLock()

			time.Sleep(time.Nanosecond)
			_ = counter

			// Разблокировка на чтение
			mu.RUnlock()
		}()

		go func() {
			defer wg.Done()

			mu.Lock()

			time.Sleep(time.Nanosecond)
			counter++

			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(start)
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}
