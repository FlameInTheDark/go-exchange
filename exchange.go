package exchange

import (
	"time"
)

// Get exchange history for specified currencies between two dates
func GetExchangeHistory(base string, symbols []string, from, to time.Time) (Rates, error) {
	raw, err := getRatesHistory(base, symbols, from, to)
	if err != nil {
		return nil, err
	}

	rates, err := raw.Rates.GetRates()
	if err != nil {
		return nil, err
	}
	rates.CalcDifference()
	return rates, nil
}

// Get the latest exchange for the specified currencies
func GetExchangeLatest(base string, symbols []string) (Rates, error) {
	raw, err := getRatesLatest(base, symbols)
	if err != nil {
		return nil, err
	}

	rates := raw.Rates.GetRates()

	return rates, nil
}

// Convert the amount of the first currency to a second
func Convert(from, to string, amount float64) (float64, error) {
	rates, err := GetExchangeLatest(from, []string{to})
	if err != nil {
		return 0, err
	}

	return rates[0].Rate[0].Exchange * amount, nil
}
