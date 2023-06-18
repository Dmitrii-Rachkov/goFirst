package main

/*
Базы данных

Работа с реляционными база данных

Для работы с реляционными базами данных в языке Go применяется встроенный пакет database/sql.
Однако он не используется сам по себе. Он лишь предоставляет универсальный интерфейс для работы
с базами данных. Для работы с конкретной СУБД нам также необходим драйвер.
Список доступных драйверов можно найти здесь. Однако поскольку драйвера реализуют одни и те
же интерфейсы, то в принципе работа с различными СУБД будет идентична.

Для того, чтобы начать работу с базой данных, необходимо открыть подключение с помощью функции Open():
func Open(driverName, dataSourceName string) (*DB, error)

Эта функция принимает в качестве параметров имя драйвера и имя источника данных, к которому надо
подключаться. Возвращает функция объект DB - по сути базу данных, с которой мы можем работать.
Если неудалось подключить к источнику данных, то в объекте error мы сможем найти сведения об ошибке.

Затем взаимодействие с базой данных осуществляется посредством методов объекта DB.
func (db *DB) Exec(query string, args ...interface{}) (Result, error)
func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
func (db *DB) QueryRow(query string, args ...interface{}) *Row
func (db *DB) Close() error         // закрывает подключение

Метод Exec() выполняет некоторое sql-выражение, которое передается через первый параметр,
не возвращая никакого результата. Метод также принимает дополнительные параметры, с помощью которых
можно передать значения в выполняемое sql-выражение. Например, абстрактная операция добавления данных
в БД, которая предполагает выполнение команды INSERT:
result, err := db.Exec("INSERT INTO Products (model, company, price) VALUES ('iPhone X', 'Apple', 72000)")

Способ вставки дополнительных параметров в SQL-выражение зависит от конкретного драйвера.
Также этот метод подходит для выполнения команд UPDATE (обновление) и DELETE (удаление).

Метод возвращает объект Result. Это интерфейс определяет два метода:
LastInsertId() (int64, error)   // возвращает id последней строки, которая была добавлена/обновлена/удалена
RowsAffected() (int64, error)   // возвращает количество затронутых строк

Метод Query() для выполнения запроса, который возвращает какие-либо данные.
Обычно это запросы, которые содержат команду SELECT.
rows, err := db.Query("SELECT name FROM users WHERE age=23")

Результатом запроса является объект *Rows - по сути набор строк.
С помощью ряда его методов можно извлечь полученные данные:
func (rs *Rows) Columns() ([]string, error)     // возвращает названия столбцов набора
func (rs *Rows) Next() bool                     // возвращает true если в наборе есть еще одна строка и переходит к ней
func (rs *Rows) Scan(dest ...interface{}) error     // считывает данные строки в переменные
func (rs *Rows) Close() error                   // закрывает объект Rows для дальнейшего чтения

Общий принцип чтения набора строк выглядит примерно следующим образом:
rows, err := db.Query("SELECT ...")
...
defer rows.Close()
for rows.Next() {
    var id int
    var name string
    rows.Scan(&id, &name)
    fmt.Println(id, name)
}

Вначале выполняем запрос к базе данных с помощью метода db.Query, затем с помощью метода Next()
последовательно считываем все строки из набора. Если строк в наборе нет, то метод возвращает false,
и происходит ыход из цикла. Если строки еще есть, то указатель *Rows переходит к следующей строке.
И затем мы можем считать в переменные с помощью метода Scan() данные из текущей строки.

Метод QueryRow() возвращает одну строку в виде объекта *Row. Как правило, этот метод применяется
для получение единичного объекта, например, по id. Этот объект имеет метод Scan(),
который позволяет извлечь данные из строки:
func (r *Row) Scan(dest ...interface{}) error

Также стоит отметить, что язык Go поддерживает создание запросов с помощью объекта Stmt,
в который можно вводить различные данные и который повышает производительность.
И также в Go имеется поддержка транзакций в виде объекта Tx.

Все эти вещи по разному реализуются в различных драйверах для конкретных систем управления
базами данных. Но общие принципы будут одни и те же. То есть общая структура работы с различными
база данных благодаря единому интерфейсу будут совпадать.
*/
