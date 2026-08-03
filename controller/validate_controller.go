package controller

import (
	"encoding/json"
	"net/http"
	"num-verify/model"
	"num-verify/response"
	"num-verify/service"
)

func ValidateController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var validate model.ValidRequest

	err := json.NewDecoder(r.Body).Decode(&validate)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.Response{
			Success: false,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	result, err := service.ValidateNumber(validate.Number)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response.Response{
			Success: false,
			Message: string(err.Error()),
			Data:    nil,
		})

		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response.Response{
		Success: true,
		Message: "Request success",
		Data:    result,
	})
}
