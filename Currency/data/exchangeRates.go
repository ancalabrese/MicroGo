package data

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	model "github.com/ancalabrese/EXPerimenta/GoMic/Currency/data/model"
	"github.com/hashicorp/go-hclog"
)

type ExchangeRate struct {
	lock  sync.Mutex
	log   hclog.Logger
	rates model.Rates
}

var er *ExchangeRate

//Singleton
func NewExchangeRate(l hclog.Logger) (*ExchangeRate, error) {
	if er != nil && er.rates.Date == time.Now().Format("2020-10-26") {
		er.log.Info("Using cached rates", "last Update", er.rates.Date)
		return er, nil
	}
	er.log.Info("Request new rates", "last Update", er.rates.Date)
	er.lock.Lock()
	defer er.lock.Unlock()
	er = &ExchangeRate{log: l, rates: model.Rates{}}
	err := er.getCurrentRates()
	if err != nil {
		return er, err
	}
	return er, nil
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
	er.log.Debug("Got currencies", er.rates)
	return nil
}

// func (er *ExchangeRate) fromJson( r io.Reader) {
// 	dec := json.NewDecoder(r)
// 	dec.
// }
