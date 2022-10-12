package http_handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, code int, message string) {
	Response(w, code, map[string]string{"error": message})
}
func Response(w http.ResponseWriter, code int, payload interface{}) {
	res, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
}
