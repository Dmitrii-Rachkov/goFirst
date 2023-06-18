package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
Буферизированный ввод-вывод

Большиство встроенных операций ввода-вывода не используют буфер.
Это может иметь отрицательный эффект для производительности приложения.
Для буферизации потоков чтения и записи в Go опредеелены ряд возможностей, которые сосредоточены
в пакете bufio.

# Запись через буфер

Для записи в источник данных через буфер в пакете bufio определен тип Writer.
Чтобы записать данные, можно воспользоваться одним из его методов:
func (b *Writer) Write(p []byte) (nn int, err error)
func (b *Writer) WriteByte(c byte) error
func (b *Writer) WriteRune(r rune) (size int, err error)
func (b *Writer) WriteString(s string) (int, error)

Write(): записывает срез байтов
WriteByte(): записывает один байт
WriteRune(): записывает один объект типа rune
WriteString(): записывает строку

При выполнении этих методов данные вначале накапливаются в буфере, а чтобы сбросить их
в источник данных, необходимо вызвать метод Flush().

Для создания потока вывода через буфер применяется функция bufio.NewWriter():
func NewWriter(w io.Writer) *Writer

Она принимает объект io.Writer - это может быть любой объект, в который идет запись:
os.Stdout, файл и т.д. В качестве результата возвращается объект bufio.Writer:
*/
func bufioWrite() {
	fmt.Println("Записываем данные в файл с помощью буфера")
	rows := []string{
		"Hello Go!",
		"Welcome to Golang",
	}

	file, err := os.Create("some.dat")
	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for _, row := range rows {
		writer.WriteString(row)  // запись строки
		writer.WriteString("\n") // перевод строки
	}
	writer.Flush() // сбрасываем данные из буфера в файл some
}

/*
В данном случае в файл через буферизированный поток вывода записываются две строки.
*/

/*
Чтение через буфер
Для чтения из источника данных через буфер в пакете bufio определен тип Reader.
Для чтения данных можно воспользоваться одним из его методов:
func (b *Reader) Read(p []byte) (n int, err error)
func (b *Reader) ReadByte() (byte, error)
func (b *Reader) ReadBytes(delim byte) ([]byte, error)
func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
func (b *Reader) ReadRune() (r rune, size int, err error)
func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
func (b *Reader) ReadString(delim byte) (string, error)

Read(p []byte): считывает срез байтов и возвращает количество прочитанных байтов
ReadByte(): считывает один байт
ReadBytes(delim byte): считывает срез байтов из потока, пока не встретится байт delim
ReadLine(): считывает строку в виде среза байт
ReadRune(): считывает один объект типа rune
ReadSlice(delim byte): считывает срез байтов из потока, пока не встретится байт delim
ReadString(delim byte): считывает строку, пока не встретится байт delim

Для создания потока ввода через буфер применяется функция bufio.NewReader():
func NewReader(rd io.Reader) *Reader

Она принимает объект io.Reader - это может быть любой объект, с которого производится
чтение: os.Stdin, файл и т.д. В качестве результата возвращается объект bufio.Reader:
*/
func bufioRead() {
	fmt.Println("Считываем данные из файла some")
	file, err := os.Open("some.dat")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		fmt.Print(line)
	}
}

/*
В данном случае идет считывания из ранее записанного файла. Для этого объект файла os.File передается
в функцию bufio.NewReader, на основании которого создается объект bufio.Reader.
Поскольку идет построчное считывание, то каждая строка считывается из потока, пока не будет
обнаружен символ перевода строки \n.
*/
