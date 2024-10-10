package flyweight

import "fmt"

/*
Концептуальный пример
В игре Counter-Strike Террористы и Контртеррористы имеют различные типы мундира. Для простоты допустим, что и Террористы,
и Контртеррористы имеют по одному типу мундира. Объект «мундир» вписан в объект «игрок» следующим образом:

Ниже приведена структура игрока. Как видим, объект «мундир» вписан в структуру игрока:
type player struct {
    dress      dress
    playerType string // Может быть T или CT
    lat        int
    long       int
}

Припустим, что у нас есть 5 Террористов и 5 Контртеррористов, то есть всего 10 игроков. Тогда мы имеем два возможных
варианта создания мундиров:

1. Каждый из 10 объектов игроков создает отдельный объект мундира и встраивает его.
Всего создается 10 объектов мундиров.

2. Мы создаем 2 объекта мундиров: - Единый Объект Мундира Террориста – его будут использовать 5 Террористов.
- Единый Объект Мундира Контртеррориста – его будут использовать 5 Контртеррористов.

Как мы видим, в Подходе 1прийдется создать 10 объектов мундиров, тогда как в Подходе 2 мы создаем только 2 объекта.
Второй подход – это суть паттерна проектирования Легковес. Два объекта мундиров, созданные нами, называют легковесными
объектами.

Паттерн Легковес находит одинаковые элементы и создает легковесные объекты. Эти легковесные объекты (мундиры) в
дальнейшем могут быть распространены между несколькими объектами (игроки). Такая практика значительно уменьшает
количество объектов мундиров, а главное – даже если мы создадим больше игроков, им все равно будет достаточно только
двух объектов мундиров.

Используя паттерн Легковес, мы сохраняем легковесные объекты в полях карты. Когда создаются другие объекты,
разделяющие между собой легковесные объекты, легковесы загружаются из карты.

Теперь давайте подумаем над тем, какие части этой системы будут относиться к «внутреннему» или «внешнему состоянию»:
- Внутреннее состояние: Мундир входит во внутреннее состояние, так как он используется несколькими объектами
Террористов и Контртеррористов.
- Внешнее состояние: Местонахождение и оружие игрока относятся ко внешнему состоянию, поскольку у каждого
объекта они разные.
*/

// dressFactory.go: Фабрика легковесов
const (
	//TerroristDressType terrorist dress type
	TerroristDressType = "tDress"
	//CounterTerrroristDressType terrorist dress type
	CounterTerrroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &DressFactory{
		dressMap: make(map[string]Dress),
	}
)

type DressFactory struct {
	dressMap map[string]Dress
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}

	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	}
	if dressType == CounterTerrroristDressType {
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}

	return nil, fmt.Errorf("Wrong dress type passed")
}

func getDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}

// dress.go: Интерфейс легковеса
type Dress interface {
	getColor() string
}

// terroristDress.go: Конкретный легковесный объект
type TerroristDress struct {
	color string
}

func (t *TerroristDress) getColor() string {
	return t.color
}

func newTerroristDress() *TerroristDress {
	return &TerroristDress{color: "red"}
}

// counterTerroristDress.go: Конкретный легковесный объект
type CounterTerroristDress struct {
	color string
}

func (c *CounterTerroristDress) getColor() string {
	return c.color
}

func newCounterTerroristDress() *CounterTerroristDress {
	return &CounterTerroristDress{color: "green"}
}

// player.go: Контекст
type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *Player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *Player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

// game.go: Клиентский код
type game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func newGame() *game {
	return &game{
		terrorists:        make([]*Player, 1),
		counterTerrorists: make([]*Player, 1),
	}
}

func (c *game) addTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	c.terrorists = append(c.terrorists, player)
	return
}

func (c *game) addCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	c.counterTerrorists = append(c.counterTerrorists, player)
	return
}

// main.go: Клиентский код
func main() {
	game := newGame()

	//Add Terrorist
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	//Add CounterTerrorist
	game.addCounterTerrorist(CounterTerrroristDressType)
	game.addCounterTerrorist(CounterTerrroristDressType)
	game.addCounterTerrorist(CounterTerrroristDressType)

	dressFactoryInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}

// output.txt: Результат выполнения
/*
DressColorType: ctDress
DressColor: green
DressColorType: tDress
DressColor: red
*/
