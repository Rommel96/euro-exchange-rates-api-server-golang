package routes

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/rommel96/rakuten-appication-api-server-golang/src/db"
)

func Routes(router *mux.Router) {
	router.HandleFunc("/rates/latest", getLatestRates).Methods("GET")
	router.HandleFunc("/rates/{year:[0-9]+}-{month:[0-9]+}-{day:[0-9]+}", getDateRates).Methods("GET")
	router.HandleFunc("/rates/analyze", getAnalyzeRates).Methods("GET")
}

func getLatestRates(w http.ResponseWriter, r *http.Request) {
	rates, err := db.FindLatestRates()
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, err)
		return
	}
	m := make(map[string]float32)
	for _, r := range rates {
		m[r.Currency] = r.Rate
	}
	responseJSON(w, http.StatusOK, sampleResponse{
		Base:  "EUR",
		Rates: m,
	})
}

func getDateRates(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	year, _ := strconv.Atoi(vars["year"])
	month, _ := strconv.Atoi(vars["month"])
	day, _ := strconv.Atoi(vars["day"])
	if year > time.Now().Year() || year < time.Now().Year()-1 {
		responseJSON(w, http.StatusBadRequest, BAD_DATE_REQUEST)
		return
	}
	if month > 12 || month == 0 {
		responseJSON(w, http.StatusBadRequest, BAD_MONT_REQUEST)
		return
	}
	if day > 31 || day == 0 {
		responseJSON(w, http.StatusBadRequest, BAD_DAY_REQUEST)
		return
	}
	dateString := strings.Join([]string{vars["year"], vars["month"], vars["day"]}, "-")
	rates, err := db.FindDateRates(dateString + "%")
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, err)
		return
	}
	if len(rates) == 0 {
		responseJSON(w, http.StatusBadRequest, map[string]interface{}{
			"base":  "EUR",
			"rates": NOT_CHANGES_RATES,
		})
		return
	}
	m := make(map[string]float32)
	for _, r := range rates {
		m[r.Currency] = r.Rate
	}
	responseJSON(w, http.StatusOK, sampleResponse{
		Base:  "EUR",
		Rates: m,
	})
}
func getAnalyzeRates(w http.ResponseWriter, r *http.Request) {
	rates, err := db.FindAnalyzeRates()
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, err)
		return
	}
	m := make(map[string]db.ValuesAnalyze)
	for _, r := range rates {
		m[r.Currency] = db.ValuesAnalyze{
			Min: r.Min,
			Max: r.Max,
			Avg: r.Avg,
		}
	}
	responseJSON(w, http.StatusOK, map[string]interface{}{
		"base":          "EUR",
		"rates_analyze": m,
	})
}
