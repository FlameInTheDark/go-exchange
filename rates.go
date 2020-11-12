package exchange

import (
	"sort"
	"time"
)

type Rate struct {
	Name       string
	Exchange   float64
	Difference float64
}

type RateDay struct {
	Date time.Time
	Rate []Rate
}

type Rates []RateDay

func (r *RawRateHistory) GetRates() (Rates, error) {
	var rateDays Rates
	for d, v := range *r {
		var rate []Rate
		for n, e := range v {
			rate = append(rate, Rate{Name: n, Exchange: e})
		}
		date, err := stringToDate(d)
		if err != nil {
			return nil, err
		}
		rateDays = append(rateDays, RateDay{Date: *date, Rate: rate})
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

func (r Rates) CalcDifference() {
	if len(r) < 2 {
		return
	}
	var prev = make(map[string]float64)
	for _, dfv := range r[len(r)-1].Rate {
		prev[dfv.Name] = dfv.Exchange
	}

	for di, dv := range r[1:] {
		for n, v := range dv.Rate {
			r[di+1].Rate[n].Difference = v.Exchange - prev[v.Name]
			prev[v.Name] = v.Exchange
		}
	}
}

func (r *RawRate) GetRates() Rates {
	var rateDays Rates
	var rate []Rate
	for n, v := range *r {
		rate = append(rate, Rate{Name: n, Exchange: v})
	}
	rateDays = append(rateDays, RateDay{Date: time.Now(), Rate: rate})
	return rateDays
}
