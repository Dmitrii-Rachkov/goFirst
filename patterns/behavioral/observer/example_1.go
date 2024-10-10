package observer

import "fmt"

/*
Концептуальный пример
На сайте интернет-магазина периодически может заканчиваться определенный товар. В то же время некоторые пользователи
могут быть заинтересованы в этом предмете, которого пока что нет в наличии. У этой проблемы может быть 3 варианта
решения:

1. Покупатель самостоятельно периодически проверяет наличие товара.
2. Интернет-магазин засыпает пользователей оповещениями о поступлениях всех новых товаров в наличие.
3. Пользователь подписывается только на тот конкретный предмет, который его интересует, и получает оповещение
о его возвращении на полки магазина. Также, на один и тот же продукт могут подписаться несколько покупателей.

Вариант 3 звучит наиболее эффективно, и фактически это и есть суть паттерна Наблюдатель. Главные элементы этого
паттерна проектирования следующие:

- Издатель — публикует событие, когда что-то происходит.
- Наблюдатель — подписывается на события субъекта и получает оповещения в случае их возникновения.
*/

// subject.go: Издатель
type Subject interface {
	register(observer Observer)
	deregister(observer Observer)
	notifyAll()
}

// item.go: Конкретный издатель
type Item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}
func (i *Item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}
func (i *Item) register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) deregister(o Observer) {
	i.observerList = removeFromslice(i.observerList, o)
}

func (i *Item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

func removeFromslice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

// observer.go: Наблюдатель
type Observer interface {
	update(string)
	getID() string
}

//  customer.go: Конкретный наблюдатель

type Customer struct {
	id string
}

func (c *Customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *Customer) getID() string {
	return c.id
}

// main.go: Клиентский код
func main() {

	shirtItem := newItem("Nike Shirt")

	observerFirst := &Customer{id: "abc@gmail.com"}
	observerSecond := &Customer{id: "xyz@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
}

//  output.txt: Результат выполнения
/*
Item Nike Shirt is now in stock
Sending email to customer abc@gmail.com for item Nike Shirt
Sending email to customer xyz@gmail.com for item Nike Shirt
*/
