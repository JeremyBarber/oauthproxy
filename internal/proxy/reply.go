/*
Copyright © 2018-2021 Neil Hemming
*/
package proxy

import (
	"encoding/json"
	"log"
	"net/http"
)

func replyServiceUnavailable(w http.ResponseWriter) {
	replyWithError(w, http.StatusServiceUnavailable, "Service unavailable")
}

func replyNotFound(w http.ResponseWriter) {
	replyWithError(w, http.StatusNotFound, "Not found")
}

func replyInvalid(w http.ResponseWriter) {
	replyWithError(w, http.StatusBadRequest, "bad request")
}

func replyWithError(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	data := make(map[string]interface{})
	data["error"] = msg
	data["error_description"] = msg
	data["error_code"] = statusCode

	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("Reply:", err, statusCode, msg)
	}
}
