package utils

import (
	"net/http"
)

type eventResponse struct {
	Data any `json:"result"`
}

type errorResponse struct {
	Err string `json:"error"`
}

func Send(w http.ResponseWriter, data any, status int) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	var toEncode any
	if err, ok := data.(error); ok {
		toEncode = errorResponse{
			Err: err.Error(),
		}
	} else {
		toEncode = eventResponse{
			Data: data,
		}
	}

	if err := Serialize(w, toEncode); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
