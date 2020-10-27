package routes

import (
	"encoding/json"
	"net/http"
)

func responseJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

type sampleResponse struct {
	Base  string             `json:"base"`
	Rates map[string]float32 `json:"rates"`
}

type ErrorMessage string

type MessageResponse string

const (
	BAD_DATE_REQUEST  ErrorMessage    = "Bad Date Request, try Again"
	BAD_MONT_REQUEST  ErrorMessage    = "Bad Month Request, try Again"
	BAD_DAY_REQUEST   ErrorMessage    = "Bad Day Request, try Again"
	NOT_CHANGES_RATES MessageResponse = "There not exists exchanges rate on specific date"
)
