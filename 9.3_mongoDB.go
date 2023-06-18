package main

/*
MongoDB

MongoDB не является реляционный СУБД, но тем не менее это тоже довольно распространенная
система управления базами данных, которую можно использовать в Go.

Для работы с MongoDB нам потребуется драйвер mgo.

Вначале установим драйвер, выполнив в командной строке/терминале следующую команду:
go get gopkg.in/mgo.v2
*/

/*
Подключение

Для подключения к серверу MongoDB необходимо использовать функцию mgo.Dial(),
в которую передается адрес сервера:
func Dial(url string) (*Session, error)

Например, подключение к серверу на локальном компьютере:
session, err := mgo.Dial("mongodb://127.0.0.1")

Функция возвращает указатель на объект Session, который представляет текущую сессию.

Используя метод DB данного объекта, мы можем получить указатель на объект Database,
который представляет конкретную базу данных на сервере.
func (s *Session) DB(name string) *Database

Все данные в базах данных MongoDB структурированы по коллекциям.
Фактически коллекция - это аналог таблицы в реляционных базах данных, представлена типом Collection.
И чтобы получить обратиться к нужной коллекции, необходимо использовать метод C():
func (db *Database) C(name string) *Collection

Например, получим коллекцию "products", которая расположена в базе данных "productdb":
// открываем соединение
session, err := mgo.Dial("mongodb://127.0.0.1")
if err != nil {
    panic(err)
}
defer session.Close()

// получаем коллекцию products в базе данных productdb
productCollection := session.DB("productdb").C("products")

Получив коллекцию, мы сможем добавлять, получать данные и проводить с ними иные операции.
И по завершении работы с сервером необходимо закрыть подключение методом Close().
*/

/*
Добавление данных
Для добавления данных в коллекцию применяется метод Insert() объекта Collection:
func (c *Collection) Insert(docs ...interface{}) error

Этот метод принимает неопределенное количество добавляемых в коллекцию объектов.
package main
import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)
type Product struct{
    Id bson.ObjectId `bson:"_id"`
    Model string `bson:"model"`
    Company string `bson:"company"`
    Price int `bson:"price"`
}
func main() {

    // открываем соединение
    session, err := mgo.Dial("mongodb://127.0.0.1")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // получаем коллекцию
    productCollection := session.DB("productdb").C("products")

    p1 := &Product{Id:bson.NewObjectId(), Model:"iPhone 8", Company:"Apple", Price:64567}
    // добавляем один объект
    err = productCollection.Insert(p1)
    if err != nil{
        fmt.Println(err)
    }

    p2 := &Product{Id:bson.NewObjectId(), Model:"Pixel 2", Company:"Google", Price:58000}
    p3 := &Product{Id:bson.NewObjectId(), Model:"Xplay7", Company:"Vivo", Price:49560}
    // добавляем два объекта
    err = productCollection.Insert(p2, p3)
    if err != nil{
        fmt.Println(err)
    }
}

Прежде всего вначале импортируем два пакета драйвера, которые содержат весь необходимый нам функционал:
"gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"

Для представления данных здесь определена структура Product. Определение каждой переменной
структуры кроме названия и типа данных содержит название поля в коллекции, с которым данная
переменная будет сопоставляться.
Например,
Model string `bson:"model"`

Переменная Model будет сопоставляться с полем "model" в коллекции.
Причем между названием переменной и поля коллекции необязательно должно быть соответствие.

Также стоит отметить, что идентификатор объекта в MongoDB представляет специальный тип bson.ObjectId,
а в базе данных ему соответствует поле "_id".

Для добавления создаем три объекта - фактически три указателя на объекты Product.
Для создания уникального идентификатора применяется функция bson.NewObjectId().
Затем добавляем объекты в коллекцию:
err = productCollection.Insert(p1)
err = productCollection.Insert(p2, p3)
*/

/*
Получение данных
Для получения данных из коллекции применяется метод Find():
func (c *Collection) Find(query interface{}) *Query

В качестве параметра он принимает критерий выборки и возвращает объект *Query.
Среди методов этого объекта следует выделить методы All() и One, которые возвращают соответственно
все объекты выборки и один объект из выборки:
func (q *Query) All(result interface{}) error
func (q *Query) One(result interface{}) (err error)

В качестве параметра оба метода принимают указатель на объект, в который будет сохраняться
результат выборки.

Например, получим ранее сохраненные объекты:
package main
import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)
type Product struct{
    Id bson.ObjectId `bson:"_id"`
    Model string `bson:"model"`
    Company string `bson:"company"`
    Price int `bson:"price"`
}
func main() {

    // открываем соединение
    session, err := mgo.Dial("mongodb://127.0.0.1")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // получаем коллекцию
    productCollection := session.DB("productdb").C("products")
    // критерий выборки
    query := bson.M{}
    // объект для сохранения результата
    products := []Product{}
    productCollection.Find(query).All(&products)

    for _, p := range products{

        fmt.Println(p.Model, p.Company, p.Price)
    }
}

Критерий выборки представляет объект bson.M{}. Пустой объект bson.M{} охватывает все документы
в коллекции. Все полученные документы передаются в объект products. И затем данные выводятся на консоль:

Также мы можем конкретизировать выборку:
query := bson.M{
    "price" : bson.M{
        "$gt":50000,
    },
}
products := []Product{}
productCollection.Find(query).All(&products)

for _, p := range products{

    fmt.Println(p.Model, p.Company, p.Price)
}
В данном случае ищем все документы, у которых поле "price" имеет значение больше 50000.
*/

/*
Обновление данных
Для обновления данных применяются методы Update()/UpdateAll() объекта Collection:
func (c *Collection) Update(selector interface{}, update interface{}) error
func (c *Collection) UpdateAll(selector interface{}, update interface{}) (info *ChangeInfo, err error)

Первый параметр методов выборки представляет критерий выборки документов, которые будут обновляться.
Второй параметр указывает, как эти документы будут обновляться.
Оба параметра задаются с помощью объекта bson.M.
Однако если метод Update обновляет только один документ, который соответствует первому параметру,
то метод UpdateAll - обновляет все элементы.

Например, изменим цену смартфона "iPhone 8":
package main
import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)
type Product struct{
    Id bson.ObjectId `bson:"_id"`
    Model string `bson:"model"`
    Company string `bson:"company"`
    Price int `bson:"price"`
}
func main() {

    // открываем соединение
    session, err := mgo.Dial("mongodb://127.0.0.1")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // получаем коллекцию
    productCollection := session.DB("productdb").C("products")

    // обновляем данные
    err = productCollection.Update(bson.M{"model": "iPhone 8"}, bson.M{"$set":bson.M{"price":45000}})
    if err != nil{
        fmt.Println(err)
    }

    products := []Product{}
    productCollection.Find(bson.M{}).All(&products)

    for _, p := range products{

        fmt.Println(p.Model, p.Company, p.Price)
    }
}
Первый аргумент метода Update - bson.M{"model": "iPhone 8"} указывает, что выбираются все элементы,
у которых поле "model" равно "iPhone 8". Второй аргумент - bson.M{"$set":bson.M{"price":45000}}
с помощью параметра $set устанавливает, какие значения будут иметь те или иные поля
(в данном случае поле "price").
*/

/*
Удаление документов
Для удаления документов из коллекции применяется методы Remove()/RemoveAll() объекта Collection:
func (c *Collection) Remove(selector interface{}) error
func (c *Collection) RemoveAll(selector interface{}) (info *ChangeInfo, err error)
Оба метода в качестве параметра принимают критерий выборки документов, которые будут удаляться.
Только метод Remove удаляет только один документ из выборки, а метод RemoveAll удаляет все
элементы выборки.

Например, удалим все смартфоны компании Vivo:
package main
import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)
type Product struct{
    Id bson.ObjectId `bson:"_id"`
    Model string `bson:"model"`
    Company string `bson:"company"`
    Price int `bson:"price"`
}
func main() {

    // открываем соединение
    session, err := mgo.Dial("mongodb://127.0.0.1")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // получаем коллекцию
    productCollection := session.DB("productdb").C("products")

    // удаляем все документы с company = "Vivo"
    _, err = productCollection.RemoveAll(bson.M{"company": "Vivo"})
    if err != nil{
        fmt.Println(err)
    }

    products := []Product{}
    productCollection.Find(bson.M{}).All(&products)

    for _, p := range products{

        fmt.Println(p.Model, p.Company, p.Price)
    }
}

*/
