package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
Отправка запросов по HTTP

Особую область применения Go представляют запросы по протоколу HTTP.
Протокол HTTP работает поверх TCP, и технически мы можем написать приложение, которое принимает
или отправляет запросы по протоколу TCP и тем самым отправлять и получать запросы и по протоколу HTTP.
Однако в связи с тем, что данный протокол и в целом сфера веб играет большую роль,
то все соответствующие функции по работе с http были выделены в отдельный пакет net/http.

Для отправки запросов в пакете net/http определен ряд функций:
func Get(url string) (resp *Response, err error)
func Head(url string) (resp *Response, err error)
func Post(url string, contentType string, body io.Reader) (resp *Response, err error)
func PostForm(url string, data url.Values) (resp *Response, err error)

Get(): отправляет запрос GET
Head(): отправляет запрос HEAD
Post(): отправляет запрос POST
PostForm(): отправляет форму в запросе POST

Рассмотрим выполнение самого простого запроса - запроса GET, для которого применяется одноименный метод:
*/
func getGoogle() {
	fmt.Println("Get запрос на сайт гугл")
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	for true {

		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)
		fmt.Println(string(bs[:n]))

		if n == 0 || err != nil {
			break
		}
	}
}

/*
Метод Get() в качестве параметра принимает адрес ресурса, к которому надо выполнить запрос,
и возвращает объект *http.Response, который инкапсулирует ответ.
Поле Body cтруктуры http.Response представляет ответ от веб-ресурса и при этом также представляет
интерфейс io.ReadCloser. А это значит, что это поле по сути является потоком для чтения,
и мы можем считать пришедшие данные через метод Read. И кроме того, для тобо, чтобы закрыть поок,
необходимо вызвать метод Close. Поэтому после запроса вызывается метод defer resp.Body.Close() и в
цикле считываем через метод Read данные и выводим на консоль.

Поскольку в данном случае ответ от веб-ресурса все равно выводится на консоль, то мы можем сократить код:
*/
func reqGoogle() {
	fmt.Println("Сокращенная отправка запроса на сайт гугл")
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
