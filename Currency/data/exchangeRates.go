package data

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	model "github.com/ancalabrese/MicroGo/Currency/data/model"
	"github.com/hashicorp/go-hclog"
)

type ExchangeRate struct {
	lock  sync.Mutex
	log   hclog.Logger
	rates model.Rates
}

var er *ExchangeRate

//NewExchangeRate returns a new instance of Exchange rate. Sigleton
func NewExchangeRate(l hclog.Logger) (*ExchangeRate, error) {
	if er != nil && er.rates.Date == time.Now().Format("2020-10-26") {
		er.log.Info("Using cached rates", "last Update", er.rates.Date)
		return er, nil
	}
	er = &ExchangeRate{log: l, rates: model.Rates{}}
	er.log.Info("Requesting new rates")
	er.lock.Lock()
	defer er.lock.Unlock()
	err := er.getCurrentRates()
	return er, err
}

//Get rate returns the ratio from base and destination rates for conversion between different currencies
func (e *ExchangeRate) GetRate(base, destination string) (float64, error) {
	br, ok := e.rates.Rates[base]
	if !ok {
		e.log.Error("Rate not found", "Requested base", base, "Destination", destination)
		return br, fmt.Errorf("Rate not found for currency %s", base)
	}
	dr, ok := e.rates.Rates[destination]
	if !ok {
		e.log.Error("Rate not found", "Requested base", base, "Destination", destination)
		return dr, fmt.Errorf("Rate nor found for currency %s", destination)
	}

	return dr / br, nil
}

func (er *ExchangeRate) getCurrentRates() error {
	resp, err := http.DefaultClient.Get("https://api.exchangeratesapi.io/latest")
	if err != nil {
		er.log.Error("Cannot connect to online DB", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		er.log.Error("Returned code", resp.StatusCode)
		return fmt.Errorf("Request code returned %d", resp.StatusCode)
	}
	er.rates.FromJson(resp.Body)
	er.rates.Rates["EUR"] = 1
	er.log.Info("Got new currencies")
	return nil
}
