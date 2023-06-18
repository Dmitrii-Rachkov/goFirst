package main

/*
Создание и открытие файлов

Для работы с файлами мы можем использовать функциональность пакета os.
Все файлы в Go представлены типом os.File. Этот тип реализует ряд интерфейсов, например,
io.Reader и io.Writer, которые позволяют читать содержимое файла и сохранять данные в файл.

С помощью функции os.Create() можно создать файл по определенному пути.
Путь к файлу передается в качестве параметра. Если подобный файл уже существует,
то он перезаписывается:
file, err := os.Create("hello.txt")
Функция возвращает объект os.File для работы с файлом и информацию об ошибке,
которая может возникнуть при создании файла.

Ранее созданный файл можно открыть с помощью функции os.Open():
file, err := os.Open("hello.txt")
Эта функция также возвращает объект os.File для работы с файлом и информацию об ошибке,
которая может возникнуть при открытии файла.

Также в нашем распоряжении есть функция os.OpenFile(), которая открывает файл,
а если файла нет, то создает его. Она принимает три параметра:

путь к файлу
редим открытия файла (для чтения, для записи и т.д.)
разрешения для доступа к файлу

Например:
// открытие файла для чтения
f1, err := os.OpenFile("sometext.txt", os.O_RDONLY, 0666)
// открытие файла для записи
f2, err := os.OpenFile("common.txt", os.O_WRONLY, 0666)

После окончания работы с файлом его следует закрыть с помощью метода Close().
package main
import (
"fmt"
"os"
)

func main() {
    file, err := os.Create("hello.txt")     // создаем файл
    if err != nil{                          // если возникла ошибка
        fmt.Println("Unable to create file:", err)
        os.Exit(1)                          // выходим из программы
    }
    defer file.Close()                      // закрываем файл
    fmt.Println(file.Name())                // hello.txt
}
С помощью функции os.Exit() можно выйти из программы.
А метод Name(), определенный для типа os.File, позволяет получить имя файла.
*/
