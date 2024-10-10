package bridge

import "fmt"

/*
Концептуальный пример

Представим, что у вас есть два типа компьютеров: Mac и Windows, а также два типа принтеров: Epson и HP.
Компьютеры и принтеры должны работать между собой в любых комбинациях. Клиент не хочет думать об особенностях
подключения принтеров к компьютерам.

Мы не хотим, чтобы при введении в эту систему новых принтеров количество кода увеличивалось по экспоненте.
Вместо создания четырех структур для 2*2 комбинаций, мы создадим две иерархии:
- Иерархия абстракции: сюда будут входить наши компьютеры
- Иерархия реализации: сюда будут входить наши принтеры

Эти две иерархии общаются между собой посредством Моста, в котором Абстракция (компьютер) содержит ссылку на Реализацию (принтер).
И абстракцию, и реализацию можно разрабатывать отдельно, не влияя друг на друга.
*/

// computer.go: Абстракция
type Computer interface {
	Print()
	SetPrinter(Printer)
}

// mac.go: Расширенная абстракция
type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

// windows.go: Расширенная абстракция
type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

// printer.go: Реализация
type Printer interface {
	PrintFile()
}

// epson.go: Конкретная реализация
type Epson struct {
}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

// hp.go: Конкретная реализация
type Hp struct {
}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}

// main.go: Клиентский код
func main() {

	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}

// output.txt: Результат выполнения
/*
Print request for mac
Printing by a HP Printer

Print request for mac
Printing by a EPSON Printer

Print request for windows
Printing by a HP Printer

Print request for windows
*/
