package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func temp() {
	fmt.Print("hello")
}
