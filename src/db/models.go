package db

import "time"

type Rates struct {
	Id       uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Time     time.Time `json:"time"`
	Currency string    `json:"currency"`
	Rate     float32   `json:"rate"`
}

type AnalyzeRates struct {
	Currency string  `gorm:"currency"`
	Min      float32 `json:"min"`
	Max      float32 `json:"max"`
	Avg      float32 `json:"avg"`
}

type ValuesAnalyze struct {
	Min float32 `json:"min" gorm:"min"`
	Max float32 `json:"max" gorm:"max"`
	Avg float32 `json:"avg" gorm:"avg"`
}
