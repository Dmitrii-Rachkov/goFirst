package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// В пакете sync/atomic много методов и они различаются в зависимости от типа данных, так сделано
// потому что раньше не было generics.
// Фактически есть 5 основных методов, просто они для разных типов.

func main() {
	// Обычный пример с mutex
	// AddMutex()

	// Такой же пример с низкоуровневой реализацией пакета Atomic
	// AddAtomic()

	// Примеры работы с переменными
	// StoreLoadSwap()

	// Сравниваем значение переменной в горутине и меняем её если true
	// compareAndSwap()

	// Использование atomic.Value
	atomicVal()
}

// AddMutex пример в котором замеряем начальное время и конечное и замеряем сколько времени работала функция
func AddMutex() {
	start := time.Now()

	var (
		counter int64
		wg      sync.WaitGroup
		mu      sync.Mutex
	)

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println("With mutex:", time.Now().Sub(start).Seconds())
}

// AddAtomic - Такая же функция как AddMutex только с более низкоуровневой реализацией.
// В этом пакете операции выполняются атомарно - либо полностью выполняется целиком, либо не выполняется,
// операции не может быть выполнена частично.
// В этом пакете используются операции на уровне процессора.
// Минус пакета в том, что он поволяет выполнить какую-то одну операцию и он работает обычно только с числовыми
// типами данных в отличии от простого mutex.
func AddAtomic() {
	start := time.Now()
	var (
		counter int64
		wg      sync.WaitGroup
	)

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			// Передаём счётчик в функцию и delta (значение на которое нам нужно изменить наш счётчик)
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println("With atomic:", time.Now().Sub(start).Seconds())
}

func StoreLoadSwap() {
	var counter int64

	// Load позволяет получить значение счётчика атомарно
	fmt.Println(atomic.LoadInt64(&counter))

	// Store позволяет получить установить нужное значение нашей переменной
	atomic.StoreInt64(&counter, 5)
	fmt.Println(atomic.LoadInt64(&counter))

	// Swap кладёт новое значение в переменную и возвращает старое значение
	fmt.Println(atomic.SwapInt64(&counter, 10))
	fmt.Println(atomic.LoadInt64(&counter))
}

// Допустим запускаем 100 горутин и при этом мы хотим чтобы одну конкретную операцию
// выполнила одна конкретная горутина.
func compareAndSwap() {
	var (
		counter int64
		wg      sync.WaitGroup
	)
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()

			// Передаём счётчик, старое значение и новое значение
			// Возвращается true если удалось сменить old на new
			if !atomic.CompareAndSwapInt64(&counter, 0, 1) {
				return
			}

			fmt.Println("Swapped goroutine number is", i)
		}(i)
	}

	wg.Wait()
	fmt.Println(counter)
}

// Можем использовать atomic.Value и класть туда другие значения
func atomicVal() {
	var (
		value atomic.Value
	)

	value.Store(1)
	fmt.Println(value.Load())

	fmt.Println(value.Swap(2))
	fmt.Println(value.Load())

	fmt.Println(value.CompareAndSwap(2, 3))
	fmt.Println(value.Load())
}
