package server

import (
	"context"

	"github.com/ancalabrese/MicroGo/Currency/data"
	"github.com/ancalabrese/MicroGo/Currency/protos/currency"
	protos "github.com/ancalabrese/MicroGo/Currency/protos/currency"
	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	log   hclog.Logger
	rates *data.ExchangeRate
	codes *data.Currencies
}

//NewCurrencyServer returns a Currency server instance
func NewCurrencyServer(r *data.ExchangeRate, c *data.Currencies, l hclog.Logger) *Currency {
	return &Currency{l, r, c}
}

//GetRate returns the current rate for the requested currency
func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRquest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())
	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}
	return &protos.RateResponse{Rate: rate}, nil
}

func (c *Currency) GetCurrencyCodes(ctx context.Context, cc *currency.CurrenciesRequest) (*protos.CurrenciesResponse, error) {
	c.log.Info("Handle GetCurrencyCodes")
	codes, err := c.codes.GetCurencies()
	if err != nil {
		return nil, err
	}
	return &protos.CurrenciesResponse{Currencies: codes}, nil
}
