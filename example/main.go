package main

import (
	"fmt"
	"log"
	"time"

	"github.com/FlameInTheDark/go-exchange"
)

func main() {
	from := time.Date(2020, 9, 10, 0, 0, 0, 0, &time.Location{})
	to := time.Date(2020, 11, 11, 0, 0, 0, 0, &time.Location{})
	rates, err := exchange.GetExchangeHistory("USD", []string{"RUB", "EUR"}, from, to)
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range *rates {
		var rates string
		for _, v := range r.Rate {
			if v.Name == "RUB" {
				rates += fmt.Sprintf("%s Ex:%f, Dif:%f", v.Name, v.Exchange, v.Difference)
			}
		}
		fmt.Printf("Y:%d, M:%d, D:%d | Dif:%s\n", r.Date.Year(), r.Date.Month(), r.Date.Day(), rates)
	}

	ratesLatest, err := exchange.GetExchangeLatest("USD", []string{"RUB", "EUR"})
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range *ratesLatest {
		for _, v := range r.Rate {
			fmt.Printf("%s | Ex: %f\n", v.Name, v.Exchange)
		}
	}
}
