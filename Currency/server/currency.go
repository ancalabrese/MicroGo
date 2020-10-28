package server

import (
	"context"

	protos "github.com/ancalabrese/MicroGo/Currency/protos/currency"
	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	log hclog.Logger
}

//NewCurrencyServer returns a Currency server instance
func NewCurrencyServer(l hclog.Logger) *Currency {
	return &Currency{l}
}

//GetRate returns the current rate for the requested currency
func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRquest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}
