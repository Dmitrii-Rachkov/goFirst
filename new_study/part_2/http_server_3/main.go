package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Payment struct {
	// Описание покупки
	Description string `json:"description"`
	// Сумма покупки USD
	Amount int `json:"amount"`
	// ФИО человека, совершающего покупку
	FullName string `json:"full_name"`
	// Место прописки человека, совершающего покупку
	Address string `json:"address"`
}

var money = 1000 // usd
var paymentHistory = make([]Payment, 0)
var mtx sync.Mutex

func payHandler(w http.ResponseWriter, r *http.Request) {
	var payment Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mtx.Lock()
	defer mtx.Unlock()

	if money-payment.Amount >= 0 {
		money -= payment.Amount
		paymentHistory = append(paymentHistory, payment)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(fmt.Sprintf("payment successfully processed, history: %v, \nbalance: %d", paymentHistory, money)))
		return
	}

	http.Error(w, "failed to pay", http.StatusInternalServerError)
	return
}

func paramsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(query.Encode()))
}

func main() {
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/params", paramsHandler)

	if err := http.ListenAndServe(":9093", nil); err != nil {
		fmt.Println("Ошибка запуска HTTP сервера: ", err)
	}
}
