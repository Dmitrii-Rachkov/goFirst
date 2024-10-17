package main

import "fmt"

/*
Интерфейсы

# Введение в интерфейсы

Интерфейсы представляют абстракцию поведения других типов. Интерфейсы позволяют определять функции,
которые не привязаны к конкретной реализации. То есть интерфейсы определяют некоторый функционал,
но не реализуют его.

Для определения интерфейса применяется ключевое слово interface:

	type имя_интерфейса interface{
	    определения_функций
	}

Например, простейшее определение интерфейса:
*/
type Vehicle interface {
	move()
}

/*
Данный интерфейс называется vehicle. Допустим, данный интерфейс представляет некоторое транспортное
средство. Он определяет функцию move(), которая не принимает никаких параметров и ничего не возвращает.

При этом важно понимать, что интерфейс - это именно абстракция, а не конкретный тип, как int, string
или структуры. К примеру, мы не можем напрямую создать объект интерфейса:
var v vehicle = vehicle{}

Интерфейс представляет своего рода контракт, которому должен соответствовать тип данных.
Чтобы тип данных соответствовал некоторому интерфейсу, данный тип должен реализовать в виде методов
все функции этого интерфейса. Например, определим две структуры:
*/

// структура "Автомобиль"
type Car struct{}

// структура "Самолет"
type Aircraft struct{}

func (c Car) move() {
	fmt.Println("Автомобиль едет")
}

func (a Aircraft) move() {
	fmt.Println("Самолёт летит")
}

func interfaceWork() {
	var tesla Vehicle = Car{}
	var boieng Vehicle = Aircraft{}
	tesla.move()
	boieng.move()
}

/*
Здесь определены две структуры: Car и Aircraft, которые, предположим, представляют, автомобиль
и самолет соответственно. Для каждой из структур определен метод move(), который имитирует перемещение\
транспортного средства. Этот метод move соответствует функции move интерфейса vehicle по типу
параметров и типу возвращаемых значений. Поскольку между методом структур и функций в интерфейсе
есть соответствие, то подобные структуры неявно реализуют данный интерфейс.

В Go интерфейс реализуется неявно. Нам не надо специально указывать, что структуры применяют
определенный интерфейс, как в некоторых других языках программирования. Для реализации типу данных
достаточно реализовать методы, которые определяет интерфейс.

Поскольку структуры Car и Aircraft реализуют интерфейс Vehicle, то мы можем определить переменные
данного интерфейса, передав им объекты структур:
var tesla Vehicle = Car{}
var boing Vehicle = Aircraft{}

Где нам могут помочь интерфейсы? Интерфейсы позволяют определить какую-то обобщенную реализацию
без привязки к конкретному типу. Например, рассмотрим следующую ситуацию:
package main

import "fmt"

type Car struct{ }
type Aircraft struct{}


func (c Car) move(){
    fmt.Println("Автомобиль едет")
}
func (a Aircraft) move(){
    fmt.Println("Самолет летит")
}

func driveCar(c Car){
    c.move()
}
func driveAircraft(a Aircraft){
    a.move()
}

func main() {

    var tesla Car = Car{}
    var boing Aircraft = Aircraft{}
    driveCar(tesla)
    driveAircraft(boing)
}
}
Допустим, в данном случае определены две структуры Car и Aircraft, которые представляют автомобиль
и самолет. Для каждой из структур определен метод перемещения move(), который условно перемещает
транспортное средство. И также определены две функции driveCar() и driveAircraft(), которые принимают
соответственно структуры Car и Aircraft и предназначены для вождения этих транспортных средств.

И отчетливо видно, что обе функции driveCar и driveAircraft фактически идентичны, они выполняют один
и те же действия, только для разных типов. И было бы неплохо, если можно было бы определить одну
обобщенную функцию для разных типов. Особенно учитывая, что у нас может быть и больше транспортных
средств - велосипед, корабль и т.д. И для вождения каждого транспортного средства придется определять
свой метод, что не очень удобно. И как раз в этом случае можно воспользоваться интерфейсами:
package main
import "fmt"

type Vehicle interface{
    move()
}

func drive(vehicle Vehicle){
    vehicle.move()
}

type Car struct{ }
type Aircraft struct{}


func (c Car) move(){
    fmt.Println("Автомобиль едет")
}
func (a Aircraft) move(){
    fmt.Println("Самолет летит")
}

func main() {

    tesla := Car{}
    boing := Aircraft{}
    drive(tesla)
    drive(boing)
}
Теперь вместо двух функций определена одна общая функция - drive(), которая в качесте параметра
принимает значение типа Vehicle. Поскольку этому интерфейсу соответствуют обе структуры
	Car и Aircraft, то мы можем передавать эти структуры в функцию drive в качесте аргументов.
*/
