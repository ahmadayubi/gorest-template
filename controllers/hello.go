package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status int `json:"status"`
}

func Hello(w http.ResponseWriter, r *http.Request){
	res := Response{
		Message: "Hello",
		Status: http.StatusOK,
	}
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())
	return
}
