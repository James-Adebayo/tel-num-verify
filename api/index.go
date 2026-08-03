package handler

import (
	"log"
	"net/http"
	"num-verify/controller"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	log.Println("Validate route reached")
	mux.HandleFunc("/validate", controller.ValidateController)

	mux.ServeHTTP(w, r)
}
