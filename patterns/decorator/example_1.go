package decorator

import "fmt"

// Концептуальный пример

// pizza.go: Интерфейс компонента
type IPizza interface {
	getPrice() int
}

// veggieMania.go: Конкретный компонент
type VeggieMania struct {
}

func (p *VeggieMania) getPrice() int {
	return 15
}

// tomatoTopping.go: Конкретный декоратор
type TomatoTopping struct {
	pizza IPizza
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

// cheeseTopping.go: Конкретный декоратор
type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

// main.go: Клиентский код
func main() {

	pizza := &VeggieMania{}

	//Add cheese topping
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}

// output.txt: Результат выполнения
/*
Price of veggeMania with tomato and cheese topping is 32
*/
