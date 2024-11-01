package principles

/*
4. Interface Segregation Principle (ISP)
Интерфейсы должны быть специфичными, а не "большими".
*/

type Printer interface {
	Print() string
}

type Scanner interface {
	Scan() string
}

type MultiFunctionDevice interface {
	Printer
	Scanner
}

type Xerox struct{}

func (x Xerox) Print() string {
	return "Printing"
}

func (x Xerox) Scan() string {
	return "Scanning"
}
