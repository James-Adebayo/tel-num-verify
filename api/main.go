package main

import (
	"net/http"
	"num-verify/controller"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	mux.HandleFunc("/validate", controller.ValidateController)

	mux.ServeHTTP(w, r)
}
