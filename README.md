# go-exchange
exchangeratesapi.io client on Golang

### Usage

Install

`go get github.com/FlameInTheDark/go-exchange`

Example

```go
package main

import (
    "log"
    "time"

    "github.com/FlameInTheDark/go-exchange"
)

func main() {
    // Get exchange history for specified currencies between two dates
    from := time.Date(2020, 9, 10, 0, 0, 0, 0, &time.Location{})
    to := time.Date(2020, 11, 11, 0, 0, 0, 0, &time.Location{})
    ratesHistory, err := exchange.GetExchangeHistory("USD", []string{"RUB", "EUR"}, from, to)
    if err != nil {
        log.Fatal(err)
    }

    // Get the latest exchange for the specified currencies
    ratesLatest, err := exchange.GetExchangeLatest("USD", []string{"RUB", "EUR"})
    if err != nil {
        log.Fatal(err)
    }

    // Convert the amount of the first currency to a second
    converted, err := exchange.Convert("EUR", "RUB", 40)
    if err != nil {
        log.Fatal(err)
    }
}
```