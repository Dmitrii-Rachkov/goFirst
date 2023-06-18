package main

/*
http.Client

Для осуществления HTTP-запросов также может применяться структура http.Client.
Чтобы отправить запрос к веб-ресурсу, можно использовать один из ее методов:
func (c *Client) Do(req *Request) (*Response, error)
func (c *Client) Get(url string) (resp *Response, err error)
func (c *Client) Head(url string) (resp *Response, err error)
func (c *Client) Post(url string, contentType string, body io.Reader) (resp *Response, err error)
func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error)

Во многом они аналогичны тем одноименным функциям (за исключением метода Do), которые определены
в пакете net/http и которые были рассмотрены в прошлой теме.
Например, выполнение самого простого запроса GET:
package main
import (
    "fmt"
    "net/http"
    "io"
    "os"
)
func main() {
   client := http.Client{}
   resp, err := client.Get("https://google.com")
   if err != nil {
         fmt.Println(err)
         return
   }
   defer resp.Body.Close()
   io.Copy(os.Stdout, resp.Body)
}
*/

/*
Настройка клиента

Структура http.Client имеет ряд полей, которые позволяют настроить ее поведение:

Timeout: устанавливает таймаут для запроса
Jar: устанавливает куки, отправляемые в запросе
Transport: определяет механиз выполнения запроса

Установка таймаута:
package main
import (
    "fmt"
    "net/http"
    "io"
    "os"
    "time"
)
func main() {
    client := http.Client{
        Timeout: 6 * time.Second,
    }
    resp, err := client.Get("https://google.com")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()
    io.Copy(os.Stdout, resp.Body)
}
Свойство Timeout представляет объект time.Duration, и в данном случае оно равно 6 секундам.
*/

/*
Request

Для управления запросом и его параметрами в Go используется объект http.Request.
Он позволяет установить различные настройки, добавить куки, заголовки, определить тело запроса.
Для создания объекта http.Request применяется функция http.NewRequest():
func NewRequest(method, url string, body io.Reader) (*Request, error)
Функция принимает три параметра. Первый параметр - тип запроса в виде строки ("GET", "POST").
Второй параметр - адрес ресурса. Третий параметр - тело запроса.

Для отправки объекта Request можно применять метод Do():
Do(req *http.Request) (*http.Response, error)

Например:
package main
import (
    "fmt"
    "net/http"
    "io"
    "os"
)
func main() {
    client := &http.Client{}
    req, err := http.NewRequest(
         "GET", "https://google.com", nil,
    )
    // добавляем заголовки
    req.Header.Add("Accept", "text/html")   // добавляем заголовок Accept
    req.Header.Add("User-Agent", "MSIE/15.0")   // добавляем заголовок User-Agent

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()
    io.Copy(os.Stdout, resp.Body)
}
*/
