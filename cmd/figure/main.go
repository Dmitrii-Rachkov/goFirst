package main

import "fmt"

/*
Спроектировать систему на уровне интерфейсов, с помощью которой можно было бы рисовать и раскрашивать геометрические фигуры,
такие как круг, квадрат, треугольник (количество фигур может быть больше).
*/

// Figure - интерфейс геометрических фигур
type Figure interface {
	Draw() string
	Paint(color string)
}

// Circle - структура круга с радиусом и цветом
type Circle struct {
	radius float64
	color  string
}

// Draw - реализуем интерфейс Figure Draw для рисования круга
func (c *Circle) Draw() string {
	return fmt.Sprintf("Circle of radius: %.2f", c.radius)
}

// Paint - реализуем интерфейс Figure Paint для раскрашивания круга
func (c *Circle) Paint(color string) {
	c.color = color
}

// Square - структура квадрата со стороной и цветом
type Square struct {
	side  float64
	color string
}

// Draw - реализуем интерфейс Figure Draw для рисования квадрата
func (s *Square) Draw() string {
	return fmt.Sprintf("Square of side: %.2f", s.side)
}

// Paint - реализуем интерфейс Figure Paint для раскрашивания квадрата
func (s *Square) Paint(color string) {
	s.color = color
}

// Triangle - структура треугольника с основанием, двумя сторонами и цветом
type Triangle struct {
	base  float64
	sideA float64
	sideB float64
	color string
}

// Draw - реализуем интерфейс Figure Draw для рисования треугольника
func (t *Triangle) Draw() string {
	return fmt.Sprintf("Triangle of base: %.2f, sideA: %.2f, sideB: %.2f", t.base, t.sideA, t.sideB)
}

// Paint - реализуем интерфейс Figure Paint для раскрашивания треугольника
func (t *Triangle) Paint(color string) {
	t.color = color
}

func main() {
	// Параметры фигур
	circle := Circle{radius: 5.1}
	square := Square{side: 7.2}
	triangle := Triangle{base: 3.0, sideA: 6.7, sideB: 8.9}

	// Рисуем фигуры
	fmt.Println(circle.Draw())
	fmt.Println(square.Draw())
	fmt.Println(triangle.Draw())

	// Раскрашиваем фигуры
	circle.Paint("red")
	fmt.Println(circle.color)

	square.Paint("blue")
	fmt.Println(square.color)

	triangle.Paint("green")
	fmt.Println(triangle.color)
}
