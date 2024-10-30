package main

import "fmt"

// Runner интерфейс
type Runner interface {
	Run() string
}

// Swimmer интерфейс
type Swimmer interface {
	Swim() string
}

// Flyer интерфейс
type Flyer interface {
	Fly() string
}

// Ducker - встраивание интерфейсов
type Ducker interface {
	Runner
	Swimmer
	Flyer
}

type Human struct {
	Name string
}

// Run - имплементируем интерфейс Runner
func (h Human) Run() string {
	return fmt.Sprintf("Человек %s бегает", h.Name)
}

func (h Human) writeCode() {
	fmt.Println("Человек пишет код")
}

type Duck struct {
	Name, Surname string
}

func (d Duck) Run() string {
	return "Утка бегает"
}

func (d Duck) Swim() string {
	return "Утка плавает"
}

func (d Duck) Fly() string {
	return "Утка летает"
}

func main() {
	//interfaceValues()
	typeAssertionAndPolymorphism()
}

// interfaceValues - интерфейсные значения, хранят в себе знания о конкретном типе и
// знание о значении конкретного типа
func interfaceValues() {
	var runner Runner
	// Если выведем на печать, то увидим Type: <nil> Value: <nil>
	// Это дефолтные значения
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)

	// Если нет конкретного типа и значения этого конкретного типа, то интерфейс равен nil
	if runner == nil {
		fmt.Println("Runner is nil")
	}

	// Мы не можем присвоить значение переменной runner, т.к. интерфейс не имплементировали
	//runner = int64(1)

	// Безопасно вызвать метод у интерфейсного типа мы можем только в том случае, если
	// интерфейсное значение не равно nil и значение конкретного типа тоже не равно nil
	// иначе будет паника
	//runner.Run()

	// Тип Human имплементирует интерфейс Runner, значит мы можем его подставить в переменную runner
	var unnamedRunner *Human
	// unnamedRunner имеет конкретный тип
	fmt.Printf("Type: %T Value: %#v\n", unnamedRunner, unnamedRunner)

	runner = unnamedRunner
	// Теперь мы видим, что runner имеет конкретный тип
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)
	// Теперь runner не равен nil, т.е. в проверке на nil сравнивается есть ли конкретный тип или нет
	if runner == nil {
		fmt.Println("Runner is nil")
	}

	namedRunner := &Human{Name: "Джек"}
	fmt.Printf("Type: %T Value: %#v\n", namedRunner, namedRunner)

	runner = namedRunner
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)

	// empty interface{}

	var emptyInterface interface{} = unnamedRunner
	fmt.Printf("Type: %T Value: %#v\n", emptyInterface, emptyInterface)

	emptyInterface = runner
	fmt.Printf("Type: %T Value: %#v\n", emptyInterface, emptyInterface)

	emptyInterface = int64(1)
	fmt.Printf("Type: %T Value: %#v\n", emptyInterface, emptyInterface)

	emptyInterface = true
	fmt.Printf("Type: %T Value: %#v\n", emptyInterface, emptyInterface)
}

func typeAssertionAndPolymorphism() {
	var runner Runner
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)

	john := &Human{"John"}
	runner = john
	polymorphism(john)
	typeAssertion(john)

	donald := &Duck{Name: "Donald", Surname: "Duck"}
	runner = donald
	polymorphism(donald)
	typeAssertion(donald)
}

// Принимаем любое значение, которое имплементирует интерфейс Runner, например Human или Duck
func polymorphism(runner Runner) {
	// вызываем метод Run в зависимости от имплементации интерфейса Runner
	fmt.Println(runner.Run())
}

// В зависимости от типа вызываем разные функции
func typeAssertion(runner Runner) {
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)
	if human, ok := runner.(*Human); ok {
		fmt.Printf("Type: %T Value: %#v\n", human, human)
		human.writeCode()
	}
	if duck, ok := runner.(*Duck); ok {
		fmt.Printf("Type: %T Value: %#v\n", duck, duck)
		fmt.Println(duck.Fly())
	}

	switch v := runner.(type) {
	case *Human:
		fmt.Println(v.Run())
	case *Duck:
		fmt.Println(v.Swim())
	default:
		fmt.Printf("Type: %T Value: %#v\n", v, v)
	}
}

/*

1. Интерфейс - спец тип в го, представляющий из себя набор сигнатур методов,
которые нужно реализовать его имплементации

2. Встраивание интерфейсов (interface embedding)

3. Явной имплементации нет, нужно просто иметь такие же методы
Утиная типизация (duck typing).
Для того, что имплементировать интерфейс, нам нужно просто реализовать все его методы

4. Интерфейсное значение внутри хранит информацию о конкретном (неинтерфейсном) типе
и его значении
                                            ConcreteType                 ConcreteType
                                           /                            /
interfaceValue = nil       interfaceValue                 interfaceValue
                                           \                            \
                                            nil                          Concrete value

5. Nil interface: нет ни типа, ни значения. Паникует (падает) при вызове методов

6. Пустой интерфейс: interface{} - содержит любое значение, ибо каждый реализует
набор из 0 методов

7. Полиморфизм

8. Утверждение типа (type assertion) - пытаемся получить значение конкретного типа
concreteValue := interfaceValue.(concreteType)
concreteValue, ok := interfaceValue.(concreteType)

9. Type switches:
    switch v := interfaceValue.(type) {
    case concreteType1:
        // some code
    case concreteType2:
        // some code
    default:
        v has the same type as interfaceValue
    }

*/
