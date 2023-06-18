package main

import (
	"fmt"
	"os"
)

/*
Чтение с консоли

В Go имеется объект os.Stdin, который реализует интерфейс io.Reader и позволяет считывать данные с консоли.
Например, мы можем использовать функцию fmt.Fscan() для считывания с консоли с помощью os.Stdin:
*/
func readConsole() {
	fmt.Println("Чтение с консоли")
	var name string
	var age int
	fmt.Println("Введите имя: ")
	fmt.Fscan(os.Stdin, &name)

	fmt.Println("Введите возраст: ")
	fmt.Fscan(os.Stdin, &age)

	fmt.Println(name, age)
}

/*
При запуске программы мы сможем вводить данные с консоли, и они перейдут в переменные name и age:

Однако также для получения ввода с консоли мы можем использовать встроенные функции
fmt.Scan(), fmt.Scanln() и fmt.Scanf(), которые аналогичны соответственно функциям
fmt.Fscan(), fmt.Fscanln() и fmt.Fscanf():

func Scan(a ...interface{}) (n int, err error)
func Scanf(format string, a ...interface{}) (n int, err error)
func Scanln(a ...interface{}) (n int, err error)

Все эти функции уже по умолчанию считывают данные с потока os.Stdin:
package main
import (
    "fmt"
    "os"
)

func main() {
    var name string
    var age int
    fmt.Print("Введите имя: ")
    fmt.Scan(&name)
    fmt.Print("Введите возраст: ")
    fmt.Scan(&age)

    fmt.Println(name, age)
}

или так

package main
import (
    "fmt"
    "os"
)

func main() {
    var name string
    var age int
    fmt.Print("Введите имя и возраст: ")
    fmt.Scan(&name, &age)
    fmt.Println(name, age)

    // альтернативный вариант
    //fmt.Println("Введите имя и возраст:")
    //fmt.Scanf("%s %d", &name, &age)
    //fmt.Println(name, age)
}
В случае если вводятся сразу несколько значений, то разделителем между ними является пробел.
Хотя теоретически строка может включать внутренние пробелы, тем не менее данные функции считывают
значение строки и других типов данных до пробела:
*/
