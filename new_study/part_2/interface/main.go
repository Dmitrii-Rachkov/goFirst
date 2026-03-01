package main

import "fmt"

type BMW struct{}

func (b *BMW) Run() {
	fmt.Println("Давлю на газ со всей силы BMW")
}

type Audi struct{}

func (a *Audi) Run() {
	fmt.Println("Полный привод Audi")
}

type Car interface {
	Run()
}

func DriveCar(c Car) {
	c.Run()
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	bmw := &BMW{}
	audi := &Audi{}
	DriveCar(bmw)
	DriveCar(audi)

	slice := []int{1, 2, 3}
	fmt.Println(slice[3]) // panic
}
