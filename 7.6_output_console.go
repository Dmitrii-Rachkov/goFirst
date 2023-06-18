package main

/*
Вывод на консоль

Стандартным потоком вывода в Go является объект os.Stdout, который фактически представляет консоль.
Например, мы могли бы вывести в этот поток данные следующим образом:
package main
import (
    "fmt"
    "os"
)

func main() {
    fmt.Fprintln(os.Stdout, "hello cold")
}
Здесь используется рассмотренная в прошлой теме функция Fprintln(), которая выводит в поток
вывода набор значений. То есть фактически в данном случае идет запись или вывод на консоль.

Однако поскольку запись в стандартный поток os.Stdout - довольно распространенная задача,
то вместо функций Fprint/Fprintln/Fprintf применяются их двойники: Println(), Print() и Printf()
соответственно, которые по умолчанию выводят данные в os.Stdout:
package main
import "fmt"

type person struct {
   name string
   age int32
   weight float64
}
func main() {
    tom := person {
        name:"Tom",
        age: 24,
        weight: 68.5,
    }
    fmt.Printf("%-10s %-10d %-10.3f\n",
               tom.name, tom.age, tom.weight)
    fmt.Print("Hello ")
    fmt.Println("cold!")
}
*/
