package main

/*
Форматируемый ввод

Пакет fmt чтение из объекта, который реализует интерфейс io.Reader.
Для этого применяются следующие функции: Fscan(), Fscanln() и Fscanf().

Функции Fscan и Fscanln
Через параметры функций Fscan и Fscanln можно получить вводимые значения:с
func Fscan(r io.Reader, a ...interface{}) (n int, err error)
func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
В качестве первого параметра передается объект io.Reader, из которого надо считывать данные,
а второй параметр представляет объекты, в которые считываются данные.
В качестве результата обе функции возвращают количество считанных байт и информацию об ошибке.
Например:
package main
import (
    "fmt"
    "os"
)

type person struct {
   name string
   age int32
   weight float64
}
func main() {
    filename := "hello2.txt"
    writeData(filename)
    readData(filename)
}

func writeData(filename string){
    // начальные данные
    var name string = "Tom"
    var age int = 24

    file, err := os.Create(filename)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()

    fmt.Fprintln(file, name)
    fmt.Fprintln(file, age)
}
func readData(filename string){

    var name string
    var age int

    file, err := os.Open(filename)
    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()

    fmt.Fscanln(file, &name)
    fmt.Fscanln(file, &age)
    fmt.Println(name, age)
}
В данном случае вначале записываем две переменных в файл с помощью fmt.Fprintln,
а затем считываем записанные значения с помощью fmt.Fscanln.
*/

/*
Fscanf
Функция fmt.Fscanf() считывает данные с применением форматирования:
func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
Первый параметр функции представляет объект io.Reader.
Второй параметр - строка форматирования, которая содержит спецификаторы и определяет последовательность
считывания данных. Третий параметр - набор объектов, в которые надо считать данные.
Например:
package main
import (
    "fmt"
    "os"
)

type person struct {
   name string
   age int32
   weight float64
}
func main() {
    filename := "person.dat"
    writeData(filename)
    readData(filename)
}

func writeData(filename string){
    // начальные данные
    tom := person { name:"Tom", age: 24, weight: 68.5 }

    file, err := os.Create(filename)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()

    // сохраняем данные в файл
    fmt.Fprintf(file, "%s %d %.2f\n", tom.name, tom.age, tom.weight)
}
func readData(filename string){

    // переменные для считывания данных
    var name string
    var age int
    var weight float64

    file, err := os.Open(filename)
    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()

    // считывание данных из файла
    _, err = fmt.Fscanf(file, "%s %d %f\n", &name, &age, &weight)
    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Printf("%-8s %-8d %-8.2f\n", name, age, weight)
}
Здесь вначале данные структуры person записываются в файл, а затем считываются из него в три переменных.
При записи данных файл мы знаем его структуру. Поэтому мы можем взять строку форматирования с той
же последовательностью спецификаторов и выполнить обратное действие - считать данные.
При считывание в объекты в функцию передаются их адреса:
fmt.Fscanf(file, "%s %d %f\n", &name, &age, &weight)
При определении строки форматирования и передаче объктов для считывания действуют те же правила,
что и при записи с помощью fmt.Fprintf. Так, первый спецификатор связан с первым объектом,
второй спецификатор - со вторым объектом и так далее. И также спецификаторы должны соответствовать
объетам по типу.
В итоге при выполнении этой программы на консоль будет выведено:
Tom	24	68.50
При этом объекты, в которые производится считывание, необязательно должны представлять переменные примитивных типов.
Например, это может быть и структура:
func readData(filename string){

    // переменная для считывания данных
    tom := person{}

    file, err := os.Open(filename)
    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()

    // считывание данных из файла
    _, err = fmt.Fscanf(file, "%s %d %f\n", &tom.name, &tom.age, &tom.weight)

    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Printf("%-8s %-8d %-8.2f\n", tom.name, tom.age, tom.weight)
}
Рассмотрим более сложный пример, когда файл содержит набор данных структур person:
package main
import (
    "fmt"
    "os"
    "io"
)

type person struct {
   name string
   age int32
   weight float64
}
func main() {
    filename := "people.dat"
    writeData(filename)
    readData(filename)
}

func writeData(filename string){
    // начальные данные
    var people = []person{
        { "Tom", 24, 68.5 },
        { "Bob", 25, 64.2 },
        { "Sam", 27, 73.6 },
    }

    file, err := os.Create(filename)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()

    for _, p := range people{
        fmt.Fprintf(file, "%s %d %.2f\n", p.name, p.age, p.weight)
    }
}
func readData(filename string){

    var name string
    var age int
    var weight float64

    file, err := os.Open(filename)
    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()

    for{
        _, err = fmt.Fscanf(file, "%s %d %f\n", &name, &age, &weight)
        if err != nil{
            if err == io.EOF{
                break
            } else{
                fmt.Println(err)
                os.Exit(1)
            }
        }
        fmt.Printf("%-8s %-8d %-8.2f\n", name, age, weight)
    }
}
Сначала функция writeData записывает в файл набор объектов person.
А затем в функции readData из файла считываются данные в бесконечном цикле. Когда файл закончится,
функция Fscanf возвратит ошибку io.EOF.
*/
