package main

import "fmt"

/*
Соответствие интерфейсу

Чтобы тип данных соответствал интерфейсу, он должен реализовать все методы этого интерфейса. Например:
*/
type Stream interface {
	read() string
	write(string)
	close()
}

func writeToStream(stream Stream, text string) {
	stream.write(text)
}
func closeStream(stream Stream) {
	stream.close()
}

// структура файла
type File struct {
	text string
}

// структура папка
type Folder struct{}

// реализация методов для типа *File
func (f *File) read() string {
	return f.text
}

func (f *File) write(message string) {
	f.text = message
	fmt.Println("Запись в файл строки", message)
}

func (f *File) close() {
	fmt.Println("Файл закрыт")
}

// релизация методов для типа *Folder
func (f *Folder) close() {
	fmt.Println("Папка закрыта")
}

func interfaceCompliance() {
	myFile := &File{}
	myFolder := &Folder{}

	writeToStream(myFile, "hello world")
	closeStream(myFile)
	// closeStream(myFolder)     // Ошибка: тип *Folder не реализует интерфейс Stream
	myFolder.close() // Так можно
}

/*
Здесь определен интерфейс Stream, который условно представляет некоторый поток данных и определяет
три метода: close, read и write. И также есть две структуры File и Folder, которые представляют
соответственно файл и папку. Для типа *File реализованы все три метода интерфейса Stream.
А для типа *Folder реализован только один метод интерфейса Stream. Поэтому тип *File реализует интерфейс
Stream и соответствует этому интерфейсу, а тип *Folder не соответствует интерфейсу Stream.
Поэтому везде, где требуется объект Stream мы можем использовать объект типа *File, но никак не *Folder.

Например, в функцию closeStream(), которая условно закрывает поток, в качестве параметра передается
объект Stream. При вызове в эту функцию можно передать объект типа *File, который соответствует
интерфейсу Stream:
closeStream(myFile)

А вот объект *Folder передать нельзя:
closeStream(myFolder)       // Ошибка: тип *Folder не реализует интерфейс Stream

Но мы по-прежнему можем вызывать у объекта *Folder метод close, но он будет рассматриваться
как собственный метод, который не имеет никакого отношения к интерфейсу Stream.

Еще важно отметить, что в данном случае методы реализованы именно для типа *File, то есть указателя
на объект File, а не для типа File. Это два разных типа. Поэтому тип File по-прежнему
НЕ соответствует интерфейсу Stream. И мы, к примеру, не можем написать следующее:
myFile2 := File{}
closeStream(myFile2)        // ! Ошибка: тип File не соответствует интерфейсу Stream
*/

/*
Реализация нескольких интерфейсов

При этом тип данных необязательно должен реализовать только методы интерфейса, для типа данных
можно определить его собственные методы или также реализовать методы других интерфейсов. Например:
package main

import "fmt"

type Reader interface{
    read()
}

type Writer interface{
    write(string)
}

func writeToStream(writer Writer, text string){
    writer.write(text)
}
func readFromStream(reader Reader){
    reader.read()
}

type File struct{
    text string
}

func (f *File) read(){
    fmt.Println(f.text)
}
func (f *File) write(message string){
    f.text = message
    fmt.Println("Запись в файл строки", message)
}

func main() {

    myFile := &File{}
    writeToStream(myFile, "hello world")
    readFromStream(myFile)
}
В данном случае для типа *File реализованы методы обоих интерфейсов - Reader и Writer.
Соответственно мы можем использовать объекты типа *File в качестве объектов Reader и Writer.
*/

/*
Вложенные интерфейсы

Одни интерфейсы могут содержать другие интерфейсы.
type Reader interface{
    read()
}

type Writer interface{
    write(string)
}

type ReaderWriter interface{
    Reader
    Writer
}

В этом случае для соответствия подобному интерфейсу типы данных должны реализовать
все его вложенные интерфейсы. Например:
package main

import "fmt"

type Reader interface{
    read()
}

type Writer interface{
    write(string)
}

type ReaderWriter interface{
    Reader
    Writer
}

func writeToStream(writer ReaderWriter, text string){
    writer.write(text)
}
func readFromStream(reader ReaderWriter){
    reader.read()
}

type File struct{
    text string
}

func (f *File) read(){
    fmt.Println(f.text)
}
func (f *File) write(message string){
    f.text = message
    fmt.Println("Запись в файл строки", message)
}

func main() {

    myFile := &File{}
    writeToStream(myFile, "hello world")
    readFromStream(myFile)
    writeToStream(myFile, "lolly bomb")
    readFromStream(myFile)
}
В данном случае определено три интерфейса. Причем интерфейс ReaderWriter содержит интерфейсы
Reader и Writer. Чтобы структура File соответствовала интерфейсу ReaderWriter, она должна реализовать
методы read и write, то есть методы обоих вложенных интерфейсов, что в принципе здесь и сделано.
*/

/* Интерфейс может быть пустым, а это значит что у него может быть любой тип данных.
Чтобы избежать ошибок от незнания с каким типом данных мы работаем, мы можем проверить
соответствует ли тип данных тому что нам нужно:

// создаем пустой интерфейс
var i interface{} = "hello"

// проверяем на соответствие типу данных с помощью 'ok' - это булеан если true то соответствует
v, ok := i.(int)
fmt.Printf("%T %v %v\n, v, v, ok")
*/
