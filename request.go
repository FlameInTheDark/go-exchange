package exchange

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const layoutISO = "2006-01-02"
const domain = "https://api.exchangeratesapi.io/"

type ResponseLatest struct {
	Rates   RawRate `json:"rates"`
	StartAt string  `json:"start_at"`
	EndAt   string  `json:"end_at"`
	Base    string  `json:"base"`
	Error   string  `json:"error"`
}

type ResponseHistory struct {
	Rates RawRateHistory `json:"rates"`
	Date  string         `json:"date"`
	Base  string         `json:"base"`
	Error string         `json:"error"`
}

type RawRate map[string]float64

type RawRateHistory map[string]RawRate

func getRatesHistory(base string, symbols []string, from, to time.Time) (*ResponseHistory, error) {
	sym := prepareSymbols(symbols)
	resp, err := http.Get(fmt.Sprintf(domain+"history?symbols=%s&base=%s&start_at=%s&end_at=%s",
		sym, base, dateToString(from), dateToString(to)))
	if err != nil {
		return nil, fmt.Errorf("api request: %s", err)
	}
	defer resp.Body.Close()

	var response ResponseHistory

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("parsing json body: %s", err)
	}

	if response.Error != "" {
		return nil, fmt.Errorf("API error response: %s", response.Error)
	}

	return &response, nil
}

func getRatesLatest(base string, symbols []string) (*ResponseLatest, error) {
	sym := prepareSymbols(symbols)
	resp, err := http.Get(fmt.Sprintf(domain+"latest?symbols=%s&base=%s",
		sym, base))
	if err != nil {
		return nil, fmt.Errorf("api request: %s", err)
	}
	defer resp.Body.Close()

	var response ResponseLatest

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("parsing json body: %s", err)
	}

	if response.Error != "" {
		return nil, fmt.Errorf("API error response: %s", response.Error)
	}

	return &response, nil
}

func prepareSymbols(symbols []string) string {
	for i, s := range symbols {
		symbols[i] = strings.ToUpper(s)
	}
	return strings.Join(symbols, ",")
}

func dateToString(t time.Time) string {
	return fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
}

func stringToDate(s string) (*time.Time, error) {
	t, err := time.Parse(layoutISO, s)
	if err != nil {
		return nil, fmt.Errorf("string to date: %s", err)
	}
	return &t, nil
}
