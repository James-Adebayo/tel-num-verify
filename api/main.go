package main

import (
	"net/http"
	"num-verify/controller"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("../index.html"))
	})
	http.HandleFunc("/validate", controller.ValidateController)
	// fmt.Println("Server started on http://localhost:8080/")
	http.ListenAndServe(":80", nil)
}
