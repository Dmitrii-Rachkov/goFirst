package main

import (
	"fmt"
	"time"
)

/*
Каналы - это инструмент коммуникации, который позволяет обмениваться данными между горутинами.
*/

func main() {
	//nilChan()
	//unbufferedChan()
	//bufferedChan()
	forRange()
}

// Функция для создания nil канала
func nilChan() {
	// Канал в который можно записывать и читать данные int типа
	var nilChannel chan int
	fmt.Printf("Len: %d Cap: %d\n", len(nilChannel), cap(nilChannel))

	// Запись/Чтение nil канала будет заблокирована навечно, нельзя так делать
	// nilChannel <- 1

	// read from nil channel blocks forever
	// <-nilChannel

	// При попытке закрыть nil канал возникает паника
	// closing nil channel will raise a panic
	// close(nilChannel)
}

// Небуферизированный канал
func unbufferedChan() {
	// Создаём канал с типом int на чтение/запись
	unbufferedChannel := make(chan int)
	fmt.Printf("Len: %d Cap: %d\n", len(unbufferedChannel), cap(unbufferedChannel))

	// В момент записи в небуферизированный канал, размер его буфера(len) = 0.
	// Горутина становится в очередь на запись.
	// Чтобы поместить данные в канал, у канала обязательно должно быть читатели и писатели.
	// Поэтому здесь происходит блокировка.
	// unbufferedChannel <- 1

	// Также если мы просто будем только лишь читать из канала, то произойдёт блокировка
	// т.к. нет писателя
	// <-unbufferedChannel

	// Пример 1
	// Создаём горутину, которая будет спать секунду, а затем запишет 1 в канал
	// Также важно, что двунаправленный канал можно передавать как однонаправленный chanForWriting chan<- int
	go func(chanForWriting chan<- int) {
		time.Sleep(time.Second)
		// <-chanForWriting
		unbufferedChannel <- 1
	}(unbufferedChannel)

	// Здесь в основной нашей горутине main мы читаем данные в переменную value
	value := <-unbufferedChannel
	fmt.Println(value)

	// Пример 2
	// Создаём горутину, которая будет спать одну секунду, а затем читает из канала в переменную value
	go func(chanForReading <-chan int) {
		time.Sleep(time.Second)
		value := <-chanForReading
		fmt.Println(value)
	}(unbufferedChannel)

	// В основной горутине main мы пишем данные в канал
	unbufferedChannel <- 2

	// Пример 3
	// Как работать с каналом после его закрытия

	// Горутина, которая закрывает канал через полсекунды
	go func() {
		time.Sleep(time.Millisecond * 500)
		close(unbufferedChannel)
	}()

	// Горутина, которая читает из канала через 1 секунду
	go func() {
		time.Sleep(time.Second)
		fmt.Println(<-unbufferedChannel)
	}()
	// Пишем в основной горутине данные в канал, но он уже закрыт и получаем панику
	unbufferedChannel <- 3

	// Пример 4
	// Нельзя закрыть уже закрытый канал, будет паника
	//close(unbufferedChannel)
	//close(unbufferedChannel)

	// Самая лучшая практика закрывать канал там, где в него пишутся данные
}

// Буферизированный канал
func bufferedChan() {
	// Создание буферизированного канала с cap = 2
	// Len - количество элементов в канале
	// Cap - емкость канала
	bufferedChannel := make(chan int, 2)
	fmt.Printf("Len: %d Cap: %d\n", len(bufferedChannel), cap(bufferedChannel))

	// В отличие от небуферизированного канала bufferedChannel <- 1 не приведёт к блокировке,
	// а просто запишет 1 в буфер.
	// Затем запишем ещё одно значение, т.к. cap = 2
	bufferedChannel <- 1
	bufferedChannel <- 2

	fmt.Printf("Len: %d Cap: %d\n", len(bufferedChannel), cap(bufferedChannel))

	// Если канал заполнен и мы попытаемся в него ещё сделать запись, то получим блокировку deadlock
	//bufferedChannel <- 3

	// Здесь операция чтения не блокируется, т.к. у нас есть данные в канале, которые можно прочитать
	// А в небуферизированном мы бы заблокировались и ждали пока произойдёт запись
	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)

	fmt.Printf("Len: %d Cap: %d\n", len(bufferedChannel), cap(bufferedChannel))

	// Когда мы попытаемся почитать в третий раз из канала, получим блокировку deadlock,
	// т.к. больше нет данных в канале
	//fmt.Println(<-bufferedChannel)
}

// Использование циклов для получения значений из канала
func forRange() {
	// Создаём буферизированный канал
	bufferedChannel := make(chan int, 3)

	numbers := []int{5, 6, 7, 8}

	// Пример 1
	// Создаём горутину, которая будет записывать значения из slice в канал
	go func() {
		for _, num := range numbers {
			bufferedChannel <- num
		}
		// Закрываем канал, тем самым сообщаем остальным что данных больше не будет
		// Однако можно читать данные из закрытого канала, и там если нет значений, то будут дефолтные int
		close(bufferedChannel)
	}()

	// Читаем значения из канала в бесконечном цикле
	for {
		// проверяем закрыт канал или нет
		v, ok := <-bufferedChannel
		fmt.Println(v, ok)

		// Если закрыт, прекращаем цикл
		if !ok {
			break
		}
	}

	// Пример 2
	// Здесь используем цикл for range
	// Нам не нужно проверять закрыт канал или нет
	bufferedChannel = make(chan int, 3)

	go func() {
		for _, num := range numbers {
			bufferedChannel <- num
		}
		close(bufferedChannel)
	}()

	for v := range bufferedChannel {
		fmt.Println("buffered", v)
	}

	// Пример 3
	// Небуферизированный канал
	// Получается здесь есть блокировки, меняем чтение и запись
	unbufferedChannel := make(chan int)

	// Получается это пишущая горутина
	go func() {
		for _, num := range numbers {
			unbufferedChannel <- num
		}
		close(unbufferedChannel)
	}()

	// а это читающая горутина получается
	for value := range unbufferedChannel {
		fmt.Println("unbuffered", value)
	}
}
