package utils

import (
	"encoding/json"
	"net/http"
)

type returnResponse struct {
	Code int 
	Message string
}

func RequestResponse(w http.ResponseWriter, code int, message string) {
	json.NewEncoder(w).Encode(returnResponse{
		Code: code,
		Message: message})
}