package singleton

/*
Существуют и другие методы создания экземпляра одиночки в Go:

1. Функция init
Мы можем создавать экземпляр одиночки внутри функции init.
Это возможно только в тех случаях, когда ранняя инициализация экземпляра не является проблемой.
Функция init вызывается единожды для каждого файла в пакете, поэтому мы можем быть уверенны в том,
что будет создан только один экземпляр.

2. sync.Once
sync.Once выполнит операцию лишь один раз. Смотрите код ниже:
*/

// Singleton
import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleInstance = &single{}
			})
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
*/
