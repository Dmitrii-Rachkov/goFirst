package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var money int64 // usd
var bank int64  // usd
var mu sync.Mutex

// payHandler - оплачиваем покупку
func payHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		msg := "failed to read body" + err.Error()
		fmt.Println(msg)
		w.WriteHeader(http.StatusBadRequest)
		_, errWr := w.Write([]byte(msg))
		if errWr != nil {
			fmt.Println("failed to write Response: ", errWr)
		}
		return
	}

	bodyString := string(body)

	price, err := strconv.Atoi(bodyString)
	if err != nil {
		msg := "failed to convert bodyString in price" + err.Error()
		fmt.Println(msg)
		w.WriteHeader(http.StatusBadRequest)
		_, errWr := w.Write([]byte(msg))
		if errWr != nil {
			fmt.Println("failed to write Response: ", errWr)
		}
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if money-int64(price) >= 0 {
		money -= int64(price)
		msg := fmt.Sprintf("pay was successfully, money: %d", money)
		fmt.Println(msg)
		w.WriteHeader(http.StatusOK)
		_, errWr := w.Write([]byte(msg))
		if errWr != nil {
			fmt.Println("failed to write Response: ", errWr)
		}
		return
	}

	fmt.Println("not enough money")
	w.WriteHeader(http.StatusBadRequest)
	_, errWr := w.Write([]byte("not enough money"))
	if errWr != nil {
		fmt.Println("failed to write Response with err: ", errWr)
	}
}

// saveHandler - сохраняем деньги в банк
func saveHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		msg := "failed to read body" + err.Error()
		fmt.Println(msg)
		w.WriteHeader(http.StatusBadRequest)
		_, errWr := w.Write([]byte(msg))
		if errWr != nil {
			fmt.Println("failed to write Response: ", errWr)
		}
		return
	}

	bodyString := string(body)

	amountMoney, err := strconv.Atoi(bodyString)
	if err != nil {
		msg := "failed to convert bodyString in amountMoney" + err.Error()
		fmt.Println(msg)
		w.WriteHeader(http.StatusBadRequest)
		_, errWr := w.Write([]byte(msg))
		if errWr != nil {
			fmt.Println("failed to write Response: ", errWr)
		}
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if money >= int64(amountMoney) {
		money -= int64(amountMoney)
		bank += int64(amountMoney)
		w.WriteHeader(http.StatusOK)
		msg := fmt.Sprintf("balance was successfully replenished, bank: %d money: %d", bank, money)
		fmt.Println(msg)
		_, errWr := w.Write([]byte(msg))
		if errWr != nil {
			fmt.Println("failed to write Response: ", errWr)
		}
		return
	}

	errMsg := "not enough money"
	fmt.Println(errMsg)
	w.WriteHeader(http.StatusBadRequest)
	_, errWr := w.Write([]byte(errMsg))
	if errWr != nil {
		fmt.Println("failed to write Response: ", errWr)
	}
}

func main() {
	// Пример запроса в google и обработка ответа
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("failed to fetch google url")
	}
	defer func() { _ = resp.Body.Close() }()

	// Проверяем статус код ответа
	if resp.StatusCode != http.StatusOK {
		fmt.Println("bad status code")
	}

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("failed to read response body")
	}
	fmt.Println(string(body[:100]))

	// Здесь уже начинается логика нашего сервера
	money = 1000

	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	err = http.ListenAndServe(":9092", nil)
	if err != nil {
		fmt.Println("Произошла ошибка при запуске HTTP сервера: ", err)
	}
}
