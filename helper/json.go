package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, v any) {
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		panic(err)
	}
}
