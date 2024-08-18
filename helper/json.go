package helper

import (
	"contact-management-restful/exception"
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, v any) {
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		panic(exception.NewBadRequest("invalid request body"))
	}
}

func WriteToResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&data); err != nil {
		panic(err)
	}
}
