package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// baseSelect()
	gracefulShutdown()
}

// Select - это по сути switch case только для каналов
func baseSelect() {
	// Пример 1

	// Буферизированный канал с емкостью 1
	bufferedChan := make(chan string, 3)

	// Записываем в этот канал значение, при этом мы не блокируемся и занимаем буфер полностью
	bufferedChan <- "first"

	// Пишем ключевое слово select, а дальше case после которого должна быть работа с каналом
	// В примере ниже мы пытаемся считать данные из канала bufferedChan и положить в него другое значение
	// Select различает три типа операции: block, unblock, default
	select {
	// Этот case - не блокирующая операция, значение в буфере уже есть и мы его просто достаем и выводим в консоль
	case str := <-bufferedChan:
		fmt.Println("read", str)
	// Если мы попытаемся записать данные "second" в канал, а буфер заполнен, то операция будет блокирующей
	case bufferedChan <- "second":
		fmt.Println("write", <-bufferedChan, <-bufferedChan)
	}
	// 1. Select сначала анализирует все case и он ищет неблокирующие операции, чтение или запись в канал,
	// которые не блокируются.
	// 2. Следующую операцию select выполнит блокирующую, будет ждать когда освободится буфер
	// В нашем кейсе выполняется неблокирующая операция
	// 3. Ветка default нам нужна для того, чтобы мы не блокировались. То есть если у нас будет кейс с блокирующей
	// операцией и мы не хотим чтобы select завис на время ожидания, мы можем написать ветку дефолт и в таком случае
	// она выполнится.
	// 4. Если есть несколько не блокирующих операций, то select вызывает какой-то рандомный case

	// Пример 2

	// Не буферизированный канал
	unbufChan := make(chan int)

	// Записываем в канал значение через 1 секунду
	go func() {
		time.Sleep(time.Second)
		unbufChan <- 1
	}()

	select {
	// Запись в буферизированный канал и вывод на консоль - неблокирующая
	case bufferedChan <- "third":
		fmt.Println("unblocking writing")
	// Операция чтения из канала - блокирующая операция
	case val := <-unbufChan:
		fmt.Println("blocking reading", val)
	// Передаем в After duration (промежуток времени), по истечении которого будет отправлено значение в канал
	// Получается, в этом кейсе мы ждём момент когда сможем прочитать данные из канала Time
	// Данные появляются по истечении duration
	case <-time.After(time.Millisecond * 1500):
		fmt.Println("time`s up")
	default:
		fmt.Println("default case")
	}

	// Пример 3

	// Предположим есть 1000 операция на каждую из которых приходится 1 nanosecond
	// Мы не хотим ждать пока выполнятся все операции, а хотим выполнить столько операций сколько успеем
	// Если мы запускаем select в цикле, то нужно timer выносить наружу цикла, это важно
	// иначе на каждую операцию мы буде обновлять наш timer и таким образом мы никогда его не дождемся
	resultChan := make(chan int)
	timer := time.After(time.Second) // timer outside loop

	// Запускаем горутину, которая пытается записать в наш канал 1000 значений по очереди.
	// Если timer заканчивается, то перестаем выполнять запись в первом case и выходим из цикла.
	// Если timer ещё не закончен выполняем default операцию, пишем в канал данные.
	// Закрываем канал и выводим на консоль все данные
	go func() {
		defer close(resultChan)

		for i := 1; i <= 1000; i++ {

			select {
			case <-timer:
				fmt.Println("time`s up")
				return
			default:
				//time.Sleep(time.Nanosecond)
				resultChan <- i
			}
		}
	}()

	for v := range resultChan {
		fmt.Println(v)
	}
}

// Пример 4

// gracefulShutdown - возможность завершить программу таким образом, чтобы не просто её сразу убить полностью
// весь процесс, а дать ей время на выполнение каких-то операций: закрытие соединений и т.д.
func gracefulShutdown() {
	// Создаём буферизированный канал, который получает интерфейс signal - системные сигналы
	sigChan := make(chan os.Signal, 1)

	// Подписываемся на два сигнала. Когда мы запустим программу, она будет слушать эти системные сигналы
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Запускаем программу и либо она сама завершится через 10 секунд, либо мы посылаем нашей программе сигнал
	// и она завершает свою работу.
	timer := time.After(10 * time.Second)

	select {
	// Здесь используем timer который сработает через 10 секунд
	case <-timer:
		fmt.Println("time`s up")
		return
	// Здесь пытаемся получить значение из канала, в который приходит сигнал
	// напирмер нажмём ctrl + c когда запускаем программу
	case sig := <-sigChan:
		fmt.Println("Stopped by signal:", sig)
		return
	}
}
