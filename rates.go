package exchange

import (
	"fmt"
	"sort"
	"time"
)

type Rate struct {
	Name       string
	Exchange   float64
	Difference float64
}

type RateDay struct {
	Date  time.Time
	Rates Rates
}

type RateDays []RateDay

type Rates []Rate

func (r Rates) GetByName(name string) (*Rate, error) {
	for i, v := range r {
		if v.Name == name {
			return &r[i], nil
		}
	}
	return nil, fmt.Errorf("name not found: %s", name)
}

// Sort and returns rates from history
func (r *RawRateHistory) GetRates() (RateDays, error) {
	var rateDays RateDays
	for d, v := range *r {
		var rate []Rate
		for n, e := range v {
			rate = append(rate, Rate{Name: n, Exchange: e})
		}
		date, err := stringToDate(d)
		if err != nil {
			return nil, err
		}
		rateDays = append(rateDays, RateDay{Date: *date, Rates: rate})
	}

	sort.SliceStable(rateDays, func(i, j int) bool {
		return rateDays[i].Date.Day() < rateDays[j].Date.Day()
	})
	sort.SliceStable(rateDays, func(i, j int) bool {
		return rateDays[i].Date.Month() < rateDays[j].Date.Month()
	})
	sort.SliceStable(rateDays, func(i, j int) bool {
		return rateDays[i].Date.Year() < rateDays[j].Date.Year()
	})
	return rateDays, nil
}

// Calculates difference between every exchange rate
func (r RateDays) CalcDifference() {
	if len(r) < 2 {
		return
	}
	var prev = make(map[string]float64)
	for _, dfv := range r[len(r)-1].Rates {
		prev[dfv.Name] = dfv.Exchange
	}

	for di, dv := range r[1:] {
		for n, v := range dv.Rates {
			r[di+1].Rates[n].Difference = v.Exchange - prev[v.Name]
			prev[v.Name] = v.Exchange
		}
	}
}

func (r *RawRate) GetRates() RateDays {
	var rateDays RateDays
	var rate []Rate
	for n, v := range *r {
		rate = append(rate, Rate{Name: n, Exchange: v})
	}
	rateDays = append(rateDays, RateDay{Date: time.Now(), Rates: rate})
	return rateDays
}
