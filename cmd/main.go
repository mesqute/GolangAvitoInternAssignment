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

	//обявление указателя структуры "сервер", содержащей данные о сетевом адресе и маршрутах
	srv := &http.Server{
		Addr:    *addr,
		Handler: routes(),
	}

	// запуск HTTP-сервера
	log.Printf("Запуск сервера на %s", *addr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
