package principles

/*
D) Dependency inversion principle / Принцип инверсии зависимостей

Этот принцип невозможно переоценить. Мы должны полагаться на абстракции, а не на конкретные реализации.
Компоненты ПО должны иметь низкую связность и высокую согласованность.

Этот принцип гласит, что высокоуровневые модули не должны зависеть от низкоуровневых модулей, а наоборот, и те, и другие
должны зависеть от абстракций. Это помогает уменьшить связь между компонентами и сделать код более гибким и поддерживаемым.
*/

// Предположим, у нас есть структура Worker, представляющая работника в компании,
// и структура Supervisor, представляющая руководителя:

type Worker struct {
	ID   int
	Name string
}

func (w Worker) GetID() int {
	return w.ID
}

func (w Worker) GetName() string {
	return w.Name
}

type Supervisor struct {
	ID   int
	Name string
}

func (s Supervisor) GetID() int {
	return s.ID
}

func (s Supervisor) GetName() string {
	return s.Name
}

// Теперь, для анти-паттерна, предположим, что у нас есть высокоуровневый модуль Department,
// который представляет отдел в компании и должен хранить информацию о работниках и руководителях,
// которые считаются низкоуровневыми модулями:

type Department struct {
	Workers     []Worker
	Supervisors []Supervisor
}

// Согласно принципу инверсии зависимостей, высокоуровневые модули не должны зависеть от низкоуровневых модулей.
// Вместо этого, оба должны зависеть от абстракций. Чтобы исправить мой пример анти-шаблона, я могу создать интерфейс Employee,
// который представляет и работника, и супервизора:

type Employee interface {
	GetID() int
	GetName() string
}

// Теперь я могу обновить Departmentструктуру, чтобы она больше не зависела от низкоуровневых модулей:

type Department struct {
	Employees []Employee
}

// И для полного рабочего примера:

package main

import "fmt"

type Worker struct {
	ID   int
	Name string
}

func (w Worker) GetID() int {
	return w.ID
}

func (w Worker) GetName() string {
	return w.Name
}

type Supervisor struct {
	ID   int
	Name string
}

func (s Supervisor) GetID() int {
	return s.ID
}

func (s Supervisor) GetName() string {
	return s.Name
}

type Employee interface {
	GetID() int
	GetName() string
}

type Department struct {
	Employees []Employee
}

func (d *Department) AddEmployee(e Employee) {
	d.Employees = append(d.Employees, e)
}

func (d *Department) GetEmployeeNames() (res []string) {
	for _, e := range d.Employees {
		res = append(res, e.GetName())
	}
	return
}

func (d *Department) GetEmployee(id int) Employee {
	for _, e := range d.Employees {
		if e.GetID() == id {
			return e
		}
	}
	return nil
}

func main() {
	dep := &Department{}
	dep.AddEmployee(Worker{ID: 1, Name: "John"})
	dep.AddEmployee(Supervisor{ID: 2, Name: "Jane"})

	fmt.Println(dep.GetEmployeeNames())

	e := dep.GetEmployee(1)
	switch v := e.(type) {
	case Worker:
		fmt.Printf("found worker %+v\n", v)
	case Supervisor:
		fmt.Printf("found supervisor %+v\n", v)
	default:
		fmt.Printf("could not find an employee by id: %d\n", 1)
	}
}

/*
Это демонстрирует принцип инверсии зависимости, поскольку Department структура зависит от абстракции
(Employee интерфейса), а не от конкретной реализации (структуры Worker или Supervisor).
Это делает код более гибким и простым в обслуживании, поскольку изменения в реализации рабочих и супервизоров н
е повлияют на Department структуру.
 */