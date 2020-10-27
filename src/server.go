package src

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rommel96/rakuten-appication-api-server-golang/src/db"
	"github.com/rommel96/rakuten-appication-api-server-golang/src/routes"
)

const HISTORICAL_RATES = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml"

func Run() {
	router := mux.NewRouter()
	routes.Routes(router)
	startUp()
	PORT := os.Getenv("PORT")
	log.Println("Listen and Server on port: ", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}

type envelope struct { //provide by https://www.onlinetool.io/xmltogo/
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Gesmes  string   `xml:"gesmes,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Subject string   `xml:"subject"`
	Sender  struct {
		Text string `xml:",chardata"`
		Name string `xml:"name"`
	} `xml:"Sender"`
	Cube struct {
		Text string `xml:",chardata"`
		Cube []struct {
			Text string `xml:",chardata"`
			Time string `xml:"time,attr"`
			Cube []struct {
				Text     string  `xml:",chardata"`
				Currency string  `xml:"currency,attr"`
				Rate     float32 `xml:"rate,attr"`
			} `xml:"Cube"`
		} `xml:"Cube"`
	} `xml:"Cube"`
}

func startUp() {
	resp, err := http.Get(HISTORICAL_RATES)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var data envelope
	err = xml.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	savedDataDB := make([]db.Rates, 0)
	for _, day := range data.Cube.Cube {
		for _, rate := range day.Cube {
			date, err := time.Parse("2006-01-02", day.Time)
			if err != nil {
				panic(err)
			}
			d := db.Rates{
				Time:     date,
				Currency: rate.Currency,
				Rate:     rate.Rate,
			}
			savedDataDB = append(savedDataDB, d)
		}
	}
	err = db.InsertInitialData(savedDataDB)
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}
