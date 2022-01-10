package main

import "net/http"

func routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/updatecash", updateCashHandler)
	mux.HandleFunc("/transfercash", transferCashHandler)
	mux.HandleFunc("/getcurrentcash", getCurrentCashHandler)

	return mux
}
