package data

import (
	"encoding/json"
	"io"
)

type Rates struct {
	Rates map[string]float64 `json:"rates"`
	Base string `json:"base"`
	Date string `json:"date"`
}

func (rates *Rates) FromJson(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(rates)
}
