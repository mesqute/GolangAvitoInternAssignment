package main

import "net/http"

func routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/update-cash", updateCashHandler)
	mux.HandleFunc("/transfer-cash", transferCashHandler)
	mux.HandleFunc("/get-current-cash", getCurrentCashHandler)

	return mux
}
