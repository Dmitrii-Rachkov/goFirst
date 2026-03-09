package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payment struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	var pay Payment
	if err := json.NewDecoder(r.Body).Decode(&pay); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello world!"))
	if err != nil {
		fmt.Println("Error writing response: ", err)
	}
}

func main() {
	http.HandleFunc("/default", defaultHandler)

	err := http.ListenAndServe(":9094", nil)
	if err != nil {
		panic(err)
	}
}
