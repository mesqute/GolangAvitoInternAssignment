package main

import (
	"net/http"
)

// Обработчик метода изменения кол-ва средств на балансе
func updateCashHandler(w http.ResponseWriter, r *http.Request) {
	//проверка, использует ли запрос метод POST.
	//Если нет, то формирует ответное сообщение об ошибке
	//с описанием разрешенных методов в заголовке.
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод запрещен", http.StatusMethodNotAllowed)
		return
	}
}

// Обработчик метода перевода средств между счетами
func transferCashHandler(w http.ResponseWriter, r *http.Request) {
	//проверка, использует ли запрос метод POST.
	//Если нет, то формирует ответное сообщение об ошибке
	//с описанием разрешенных методов в заголовке.
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод запрещен", http.StatusMethodNotAllowed)
		return
	}
}

// Обработчик метода получения текущего баланса
func getCurrentCashHandler(w http.ResponseWriter, r *http.Request) {
	//проверка, использует ли запрос метод GET.
	//Если нет, то формирует ответное сообщение об ошибке
	//с описанием разрешенных методов в заголовке.
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Метод запрещен", http.StatusMethodNotAllowed)
		return
	}
}
