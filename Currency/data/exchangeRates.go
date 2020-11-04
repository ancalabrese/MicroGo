package data

import (
	"fmt"
	"math/rand"
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

// MonitorRates checks the rates in the EXChangerate API every interval and sends a message to the
// returned channel when there are changes
//
// Note: the ECB API only returns data once a day, this function only simulates the changes
// in rates for demonstration purposes
func (e *ExchangeRate) MonitorRates(interval time.Duration) chan struct{} {
	ret := make(chan struct{})

	go func() {
		ticker := time.NewTicker(interval)
		for {
			select {
			case <-ticker.C:
				// just add a random difference to the rate and return it
				// this simulates the fluctuations in currency rates
				for k, v := range e.rates.Rates {
					// change can be 10% of original value
					change := (rand.Float64() / 10)
					// is this a postive or negative change
					direction := rand.Intn(1)

					if direction == 0 {
						// new value with be min 90% of old
						change = 1 - change
					} else {
						// new value will be 110% of old
						change = 1 + change
					}

					e.rates.Rates[k] = v * change
				}

				// notify updates, this will block unless there is a listener on the other end
				ret <- struct{}{}
			}
		}
	}()

	return ret
}
