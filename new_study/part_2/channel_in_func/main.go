package main

import (
	"fmt"
	"sync"
	"time"
)

/*
// mine - в нее передаем wg обязательно по ссылке, не копия. Иначе просто копируем: счетчик горутин, сколько ждут, семафор
// Также mutex тоже передаем по ссылке, иначе мы просто копируем: state (состояние), семафор

func mine(wg *sync.WaitGroup, transferPoint chan int, n int) {
	defer wg.Done()
	fmt.Println("Поход в шахту номер: ", n)
	time.Sleep(1 * time.Second)
	fmt.Println("Завершен поход в шахту: ", n)

	transferPoint <- 10
}

func main() {
	coal := 0
	initTime := time.Now()
	transferPoint := make(chan int, 3)
	var wg sync.WaitGroup

	wg.Add(3)
	go mine(&wg, transferPoint, 1)
	go mine(&wg, transferPoint, 2)
	go mine(&wg, transferPoint, 3)

	go func() {
		wg.Wait()
		close(transferPoint)
	}()

	for result := range transferPoint {
		coal += result
	}

	fmt.Println("Добыли угля: ", coal)
	fmt.Println("Время на добычу: ", time.Since(initTime))
}

*/

/*
// Простейший счётный СЕМАФОР
func main() {
	// Создаем семафор на 3 места
	sem := make(chan struct{}, 3)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg, sem, i)
	}

	wg.Wait()
	fmt.Println("Все воркеры завершили работу")
}

func worker(wg *sync.WaitGroup, sem chan struct{}, id int) {
	defer wg.Done()
	// Захватываем слот (ждем, если все заняты)
	sem <- struct{}{}        // P операция (wait)
	defer func() { <-sem }() // V операция (signal)

	// Работа с ограниченным ресурсом
	fmt.Printf("Воркер %d начал работу\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Воркер %d закончил\n", id)
}
*/

// Простейший БИНАРНЫЙ СЕМАФОР
func main() {
	// Бинарный семафор (только 1 место!)
	binarySem := make(chan struct{}, 1)
	var wg sync.WaitGroup

	// Запускаем 5 воркеров, но только 1 может работать одновременно
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(&wg, binarySem, i)
	}

	wg.Wait()
	fmt.Println("Все воркеры завершили работу")
}

func worker(wg *sync.WaitGroup, sem chan struct{}, id int) {
	defer wg.Done()

	// Пытаемся захватить семафор (занять единственное место)
	sem <- struct{}{}        // P операция (wait)
	defer func() { <-sem }() // V операция (signal)

	fmt.Printf("🔴 Воркер %d НАЧАЛ работу (семафор ЗАНЯТ)\n", id)
	time.Sleep(1 * time.Second) // Имитация работы
	fmt.Printf("✅ Воркер %d ЗАКОНЧИЛ работу (семафор СВОБОДЕН)\n\n", id)
}
