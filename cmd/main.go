package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	//переменная,содержащая сетевой адрес дислокации сервера,
	//получающая значение из флагов командной сроки при запуске приложения
	//(по умолчанию ":8080")
	addr := flag.String("addr", ":8080", "сетевой адрес HTTP")

	//считывание флагов из командной строки
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/updatecash", updateCashHandler)
	mux.HandleFunc("/transfercash", transferCashHandler)
	mux.HandleFunc("/getcurrentcash", getCurrentCashHandler)

	log.Printf("Запуск сервера на %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}
