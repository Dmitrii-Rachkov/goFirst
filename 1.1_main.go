/*Программа на языке Go хранится в одном или нескольких файлах.
Каждый файл с программным кодом должен принадлежать какому-нибудь пакету.
И вначале каждого файла должно идти объявление пакета, к которому этот
файл принадлежит. Пакет объявляется с помощью ключевого слова 'package'*/

/*В файле может использоваться функционал из других пакетов.
В этом случае используемые пакеты надо импортировать с помощью ключевого слова import.
Импортируемые пакеты должны идти после объявления пакета для текущего файла:
Например, в данном случае текущий файл будет находиться в пакете main.
И далее он подключает пакет fmt.
Причем главный пакет программы должен называться "main".
Так как именно данный пакет определяет, что будет создаваться исполняемый файл приложения,
который после компиляции можно будет запускать на выполнение.*/

/*После подключения других пакетов располагаются объявления типов, переменных, функций, констант.
При этом входной точкой в приложения является функция с именем main.
 Она обязательно должна быть определена в программе.
 Все, что выполняется в программе, выполняется именно в функции main.*/

/*Базовым элементом программы являются инструкции.
Например, вызов функции fmt.Println("Hello Go!") представляет отдельную инструкцию.
Каждая инструкция выполняет определенное действие и размещается на новой строке.
Можно размещать несколько инструкций и на одной строке, но тогда их надо отделять точкой запятой*/

package main

import (
	"fmt"
	"goFirst/computation"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println("Hello go")
	variables1()                 // переменные
	typeData()                   // типы данных
	constantEx()                 // константы
	arithmeticOperations()       // арифметические операции
	relation()                   // условные операторы
	logicOperations()            // логичесик операции
	shift()                      // операции сдвига
	bitwise()                    // поразрядные операции
	arrays()                     // массивы
	ifElse()                     // условные конструкции if, else
	switchSimple()               // условная конструкция switch простая
	switchSum()                  // условная конструкция switch sum (сложение)
	switchDefault()              // условная конструкция switch с default
	switchMany()                 // условная конструкция switch с несколькими аргументами
	for1()                       // простой цикл for
	for2()                       // цикл for без объявления переменной
	for3()                       // цикл for с счётчиком в теле цикла
	for4()                       // цикл который использует только условие
	forFor()                     // цикл вложенный в другой цикл
	forArray()                   // перебор массива
	forUnderscore()              // перебор массива без индекса
	forStandard()                // стандартный перебор массива
	forContinue()                // оператор continue
	forBreak()                   // оператор break
	add(4, 5)                    // функция с параметрами
	addType(5, 3, 6.3, 1.1, 2.2) // указываем один раз тип аргументов функции
	addMany(6, 7)                // функция принимает неизвестное кол-во аргументов
	var a = returnSum(2, 2)      // возвращение результата из функции
	fmt.Println(a)
	var b = returnName(6, 4) // именованные возвращаемые результаты функции
	fmt.Println(b)
	var age, name = returnMany(4, 5, "Tom", "Simpson") // возврат разных значений
	fmt.Println(age, name)
	typeInt()              // переменная которая является функцией в другой функции
	changeFunc()           // переменная указывает сначала на одну функцию потом на другую
	action(10, 25, sumInt) // функция которая принимает на вход другую функцию
	action(5, 6, multiplication)
	validNumber()             // считаем сумму чисел среза если соответствует критериям
	anyFunc()                 // функция в которой возвращается другая функция
	anonimSum()               // анонимная функция
	anonimArgument()          // анонимная функция как аргумент
	anonimResult()            // анонимная функция как результат другой функции
	anonimEnv()               // анонимная функция с доступом к окружению
	fmt.Println(factorial(5)) // вычисление факториала рекурсивной функцией
	fmt.Println(fibonachi(6)) // вычисление фибоначи рекрсией
	deferFunc()               // функция defer срабатывает в конце программы не смотря на объявление
	// в начале программы
	deferMany()                // много функции defer
	fmt.Println(divide(10, 5)) // работа оператора panic если второе число 0
	oneSlice()                 // работа среза
	makeSlice()                // функция make в срезах
	addSlice()                 // добавление в срез
	operatorSlice()            // использование оператора i:j для среза
	deleteSlice()              // удаляем элемент среза
	mapOne()                   // отображения или map (ключ: значение)
	pointers()                 // указатель
	newPointer()               // функция new в указателях
	startChange()              // функция работает с копией переменной
	newVar()                   // изменяем исходную переменную указателем
	returnPointer()            // функция возвращает указатель
	declarationType()          // объявление производных типов
	structurePeople()          // создание структуры people
	nestedStruct()             // вложенные структуры
	var lib library = library{"Book1", "Book2", "Book3"}
	lib.print()                      // работа метода для определенного типа
	tomPerson()                      // метод структуры
	tomPersonPointer()               // метод указателя структуры
	fmt.Println(computation.Fact(5)) // используем функцию из другого пакета
	interfaceWork()                  // интерфейс в работе
	interfaceCompliance()            // соответствие интерфейсу
	goroutine()                      // функция горутин
	goroutineInput()                 // функция горутин с ожиданием ввода в консоль
	blockGoroutine()                 // блокировка горутин пока не получат данные
	buferChannel()                   // буферезированные каналы
	capLen()                         // емкость и количество элементов в канале
	closeChannel()                   // закрытие канала
	closeValid()                     // проверяем закрыт ли канал
	sync()                           // синхронизация горутин
	streamData()                     // передаем поток данных
	readPhone()                      // чтение телефонных номеров
	writeFile()                      // запись строки в файл
	readFile()                       // чтение из файла
	outStream()                      // выодим данные из файлы с помощью потока
	outFormat()                      // форматированный вывод
	readConsole()                    // чтение с консоли
	bufioWrite()                     // запись в файл с помощью буфера
	bufioRead()                      // чтение из файла с помощью буфера
	// request()                     // отправляем запрос на сайт golang
	getGoogle() // get запрос на сайт гугл
	reqGoogle() // сокращенная отправка запроса на сайт гугл
}

/*Установка*/

// Скачиваем go компилятор с сайта https://go.dev
// Запускаем скачанный файл и устанавливаем по умолчанию
// Устанавливаем VS Code
// В разделе Extensions набираем 'go' и устанавливаем плагин от google
// Сочетание клавиш CTRL + SHIFT + P открываем команды и пишем: 'go install'
// Ставим все инструменты и внизу в углу IDE тоже там будут ошибки
// их также устраняем путём установки всего по подсказкам
// Чтобы запускать через 'Run and Debug' надо создать go.mod
// Команда в терминале: 'go mod init goFirst', где goFirst имя проекта
// Далее не знаю зачем команда 'go mod tidy'
// Запуск программы через терминал: 'go run 1.1_main.go'
