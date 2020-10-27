package db

func InsertInitialData(rates []Rates) error {
	for _, r := range rates {
		err := db.Create(&r).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func FindLatestRates() ([]Rates, error) {
	var rates []Rates
	err := db.Debug().Order("time DESC").Limit(32).Find(&rates).Error
	if err != nil {
		return nil, err
	}
	return rates, nil
}

func FindDateRates(filter string) ([]Rates, error) {
	var rates []Rates
	err := db.Debug().Find(&rates, "time LIKE ?", filter).Error
	if err != nil {
		return nil, err
	}
	return rates, nil
}

func FindAnalyzeRates() ([]AnalyzeRates, error) {
	var result []AnalyzeRates
	err := db.Debug().Table("rates").Group("currency").Select("currency ,MIN(rate) AS min, MAX(rate) AS max, AVG(rate) AS avg").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
