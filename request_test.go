package exchange

import (
	"testing"
	"time"
)

func Test_dateToString(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "date to string", args: args{t: time.Date(2020, 11, 11, 0, 0, 0, 0, &time.Location{})}, want: "2020-11-11"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dateToString(tt.args.t); got != tt.want {
				t.Errorf("prepareDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prepareSymbols(t *testing.T) {
	type args struct {
		symbols []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "single", args: args{symbols: []string{"EUR"}}, want: "EUR"},
		{name: "single lowercase", args: args{symbols: []string{"eur"}}, want: "EUR"},
		{name: "single different", args: args{symbols: []string{"eUr"}}, want: "EUR"},
		{name: "multiple", args: args{symbols: []string{"EUR", "USD"}}, want: "EUR,USD"},
		{name: "multiple lowercase", args: args{symbols: []string{"eur", "usd"}}, want: "EUR,USD"},
		{name: "multiple different", args: args{symbols: []string{"eUr", "USD"}}, want: "EUR,USD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareSymbols(tt.args.symbols); got != tt.want {
				t.Errorf("prepareSymbols() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringToDate(t *testing.T) {
	var (
		date = time.Date(2020, 11, 10, 0, 0, 0, 0, &time.Location{})
	)
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    *time.Time
		wantErr bool
	}{
		{name: "string '2020-11-10' to time", args: args{s: "2020-11-10"}, want: &date, wantErr: false},
		{name: "string 'not a time' to time", args: args{s: "not a time"}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := stringToDate(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("stringToDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got != nil) && got.String() != tt.want.String() {
				t.Errorf("stringToDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
