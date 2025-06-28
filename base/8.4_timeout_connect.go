package main

/*
Установка таймаута

При взаимодействии клиента и сервера мы можем устанавливать таймаут, по истечении которого соединение
между сервером и клиентом при отсутствии взаимодействия будет разорвано.
Для этого у типа net.Conn определены следующие методы:

SetDeadline(t time.Time) error: устанавливает таймаут на все операции ввода-вывода.
Для установки времени применяется структура time.Time

SetReadDeadline(t time.Time) error: устанавливает таймаут на операции ввода в поток

SetWriteDeadline(t time.Time) error: устанавливает таймаут на операции вывода из потока

В каком случае они могут пригодиться? В прошлой теме было рассмотрено взаимодействие сервера и клиента.
Для чтения данных от клиента сервер использовал буфер фиксированного размера:
input := make([]byte, (1024 * 4))
n, err := conn.Read(input)

Однако в ряде ситуаций это не лучший способ, особенно когда размер передаваемых данных превышает
размер буфера. Мы можем точно не знать, сколько данных возвратит нам сервер.
Поэтому определим следующий код клиента:
package main
import (
    "fmt"
    "net"
    "time"
)
func main() {

    conn, err := net.Dial("tcp", "127.0.0.1:4545")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()
    for{
        var source string
        fmt.Print("Введите слово: ")
        _, err := fmt.Scanln(&source)
        if err != nil {
            fmt.Println("Некорректный ввод", err)
            continue
        }
        // отправляем сообщение серверу
        if n, err := conn.Write([]byte(source));
        n == 0 || err != nil {
            fmt.Println(err)
            return
        }
        // получем ответ
        fmt.Print("Перевод:")
        conn.SetReadDeadline(time.Now().Add(time.Second * 5))
        for{
            buff := make([]byte, 1024)
            n, err := conn.Read(buff)
            if err !=nil{ break}
            fmt.Print(string(buff[0:n]))
            conn.SetReadDeadline(time.Now().Add(time.Millisecond * 700))
        }
        fmt.Println()
    }
}
Теперь получение данных выделено в отдельный цикл for:
for{
    buff := make([]byte, 1024)
    n, err := conn.Read(buff)
    if err !=nil{ break}
    fmt.Print(string(buff[0:n]))
    conn.SetReadDeadline(time.Now().Add(time.Millisecond * 700))
}
Поэтому даже если сервер передал больше 1024 байт, все они все равно будут обработаны.
Но кроме того, здесь также устанавливается таймаут на чтение данных.
Перед самим циклом устанавливается таймаут в 5 секунд:
conn.SetReadDeadline(time.Now().Add(time.Second * 5))

Это значит, что клиент может ожидать данные на чтение от сервера в течении 5 секунд.
По истечении этого времени операция чтения генерирует ошибку и соответственно происходит
выход из цикла, где мы пытаемся прочитать данные от сервера. 5 секунд - довольно большой период,
но в начале перед первым взаимодействием лучше устанавливать период побольше.
И после прочтения первых 1024 байт таймаут сбрасывается до 700 миллисекунд.
То есть если в течение последующих 700 милисекунд сервер не пришлет никаких данных,
то происходит выход из цикла и соответственно чтение данных заканчивается.

Важно понимать роль подобных задержек, так как они позволяют сгенерировать ошибку при чтении данных.
А значит мы можем получить эту ошибку и должным образом обработать ее, например, выйти из бесконечного
цикла. Если бы мы не использовали установку таймаута, то могла бы сложиться ситуация, когда сервер
ожидал данных от клиента в операции чтения, а клиент ожидал данных от сервера также в операции чтения.
И была бы своего рода блокировка.

Код сервера остается тем же, что и в прошлой теме:
package main
import (
    "fmt"
    "net"
)
var dict = map[string]string{
    "red": "красный",
    "green": "зеленый",
    "blue": "синий",
    "yellow": "желтый",
}

func main() {
    listener, err := net.Listen("tcp", ":4545")

    if err != nil {
        fmt.Println(err)
        return
    }
    defer listener.Close()
    fmt.Println("Server is listening...")
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println(err)
            conn.Close()
            continue
        }
        go handleConnection(conn)  // запускаем горутину для обработки запроса
    }
}
// обработка подключения
func handleConnection(conn net.Conn) {
    defer conn.Close()
    for {
        // считываем полученные в запросе данные
        input := make([]byte, (1024 * 4))
        n, err := conn.Read(input)
        if n == 0 || err != nil {
            fmt.Println("Read error:", err)
            break
        }
        source := string(input[0:n])
        // на основании полученных данных получаем из словаря перевод
        target, ok := dict[source]
        if ok == false{             // если данные не найдены в словаре
            target = "undefined"
        }
        // выводим на консоль сервера диагностическую информацию
        fmt.Println(source, "-", target)
        // отправляем данные клиенту
        conn.Write([]byte(target))
    }
}
*/
