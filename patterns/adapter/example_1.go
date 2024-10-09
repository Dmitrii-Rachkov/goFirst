package adapter

/*
Концептуальный пример
Мы имеем код клиента, ожидающий определенных от объекта определенных качеств (порт Lightning),
но также мы имеем другой объект под названием adaptee (ноутбук на Windows), который предоставляет тот же функционал,
но через другой интерфейс (USB порт).

В такой ситуации нам подойдет паттерн Адаптер. Мы создадим структуру adapter, которая будет:

- Реализовать тот же интерфейс, который ожидает клиент (порт Lightning).

- Переводить запрос от клиента к адаптируемому объекту в форме, которую он ожидает.
Адаптер принимает коннектор Lightning, после чего переводит его сигналы в формат USB в ноутбуке на Windows.
*/

//  client.go: Клиентский код
import "fmt"

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}

// computer.go: Интерфейс клиента
type Computer interface {
	InsertIntoLightningPort()
}

// mac.go: Сервис
type Mac struct {
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

// windows.go: Неизвестный сервис
type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

// windowsAdapter.go: Адаптер
type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}

// main.go
func main() {

	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}

//  output.txt: Результат выполнения
/*
Client inserts Lightning connector into computer.
Lightning connector is plugged into mac machine.
Client inserts Lightning connector into computer.
Adapter converts Lightning signal to USB.
USB connector is plugged into windows machine.
*/
