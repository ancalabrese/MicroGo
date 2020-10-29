package server

import (
	"context"

	"github.com/ancalabrese/MicroGo/Currency/data"
	protos "github.com/ancalabrese/MicroGo/Currency/protos/currency"
	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	log   hclog.Logger
	rates *data.ExchangeRate
}

//NewCurrencyServer returns a Currency server instance
func NewCurrencyServer(r *data.ExchangeRate, l hclog.Logger) *Currency {
	return &Currency{l,r}
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
