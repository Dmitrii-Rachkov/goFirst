package singleton

/*
Концептуальный пример
Обычно, экземпляр одиночки создается во время начальной инициализации структуры.
Для этого, мы определяем для структуры метод getInstance. Этот метод будет создавать и возвращать экземпляры одиночки.
После создания первого экземпляра, при каждом вызове метода getInstance будет возвращаться именно он.

А что касательно потоков goroutine? Структура одиночки должна возвращать один и тот же экземпляр в случаях,
когда разные потоки пытаются получить доступ к этому экземпляру.
По этой причине можно легко ошибиться и неправильно реализовать паттерн Одиночка.
Пример ниже показывает, как правильно создать Одиночку.

Некоторые важные детали, о которых нужно помнить:

- В начале нужна nil-проверка, с ее помощью мы убеждаемся, что первый экземпляр singleInstance — пустой.
Благодаря этому мы можем избежать ресурсоемких операций блокировки при каждом вызове getInstance. Если эта проверка не пройдена, тогда поле singleInstance уже заполнено.

- Структура singleInstance создается внутри блокировки.

- После блокировки используется еще одна nil-проверка. В случаях, когда первую проверку проходит более одного потока,
вторая обеспечивает создание экземпляра одиночки единым потоком.
В противном случае, все потоки создавали бы свои экземпляры структуры одиночки.
*/

// Singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

// Клиентский код
func main() {

	for i := 0; i < 30; i++ {
		go getInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}

//  output.txt: Результат выполнения
/*
Creating single instance now.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
*/
