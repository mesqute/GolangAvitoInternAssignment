package main

import (
	"flag"
	"log"
	"net/http"
)

// Обработчик метода изменения кол-ва средств на балансе
func updateCashHandler(w http.ResponseWriter, r *http.Request) {

}

// Обработчик метода перевода средств между счетами
func transferCashHandler(w http.ResponseWriter, r *http.Request) {

}

// Обработчик метода получения текущего баланса
func getCurrentCashHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//переменная,содержащая сетевой адрес дислокации сервера,
	//получающая значение из флагов командной сроки при запуске приложения
	//(по умолчанию ":8080")
	addr := flag.String("addr", ":8080", "сетевой адрес HTTP")

	mux := http.NewServeMux()
	mux.HandleFunc("/updatecash", updateCashHandler)
	mux.HandleFunc("/transfercash", transferCashHandler)
	mux.HandleFunc("/getcurrentcash", getCurrentCashHandler)

	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}
