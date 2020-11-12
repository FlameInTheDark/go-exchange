package exchange

import (
	"time"
)

func GetExchangeHistory(base string, symbols []string, from, to time.Time) (*Rates, error) {
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

func GetExchangeLatest(base string, symbols []string) (*Rates, error) {
	raw, err := getRatesLatest(base, symbols)
	if err != nil {
		return nil, err
	}

	rates := raw.Rates.GetRates()

	return rates, nil
}
