package principles

import "fmt"

/*
O) Open–closed principle / Принцип открытости-закрытости

Программные объекты должны быть открыты для расширения, но закрыты для модификации.
Речь о том, что нельзя переопределять методы или классы, просто добавляя дополнительные функции по мере необходимости.

Хороший способ решения этой проблемы – использование наследования. В JavaScript эта проблема решается с помощью композиции.

Простое правило: если вы изменяете сущность, чтобы сделать ее расширяемой, вы впервые нарушили этот принцип.
*/

// Допустим, у меня есть задача построить платежную систему, которая сможет обрабатывать платежи по кредитным картам.
// Она также должна быть достаточно гибкой, чтобы принимать различные типы платежных методов в будущем.

type PaymentMethod interface {
	Pay()
}

type Payment struct{}

func (p Payment) Process(pm PaymentMethod) {
	pm.Pay()
}

type CreditCard struct {
	amount float64
}

func (cc CreditCard) Pay() {
	fmt.Printf("Paid %.2f using CreditCard", cc.amount)
}

func main() {
	p := Payment{}
	cc := CreditCard{12.23}
	p.Process(cc)
}

/*
Согласно OCP, моя Payment структура открыта для расширения и закрыта для изменения. Поскольку я использую
PaymentMethod интерфейс, мне не нужно редактировать Payment поведение при добавлении новых способов оплаты.
Добавление чего-то вроде PayPal будет выглядеть так:
*/

type PayPal struct {
	amount float64
}

func (pp PayPal) Pay() {
	fmt.Printf("Paid %.2f using PayPal", pp.amount)
}

// then in main()
// pp := PayPal{22.33}
// p.Process(pp)
