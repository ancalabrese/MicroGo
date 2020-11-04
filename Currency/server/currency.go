package server

import (
	"context"
	"io"
	"time"

	"github.com/ancalabrese/MicroGo/Currency/data"
	"github.com/ancalabrese/MicroGo/Currency/protos/currency"
	protos "github.com/ancalabrese/MicroGo/Currency/protos/currency"
	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	log         hclog.Logger
	rates       *data.ExchangeRate
	codes       *data.Currencies
	subscribers map[protos.Currency_SubscribeServer][]*protos.RateRquest
}

//NewCurrencyServer returns a Currency server instance
func NewCurrencyServer(r *data.ExchangeRate, c *data.Currencies, l hclog.Logger) *Currency {
	currency := &Currency{l, r, c, make(map[protos.Currency_SubscribeServer][]*protos.RateRquest)}

	go currency.handleRatesUpdates()

	return currency
}

//GetRate returns the current rate for the requested currency
func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRquest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())
	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}
	return &protos.RateResponse{Base: rr.Base, Destination: rr.Destination, Rate: rate}, nil
}

//GetCurrencyCodes returns all the availbale currencies codes
func (c *Currency) GetCurrencyCodes(ctx context.Context, cc *currency.CurrenciesRequest) (*protos.CurrenciesResponse, error) {
	c.log.Info("Handle GetCurrencyCodes")
	codes, err := c.codes.GetCurencies()
	if err != nil {
		return nil, err
	}
	return &protos.CurrenciesResponse{Currencies: codes}, nil
}


//Subscribe implement the gRPC bidirectional streaming method for the server
func (c *Currency) Subscribe(src protos.Currency_SubscribeServer) error {
	for {
		rr, err := src.Recv()
		if err == io.EOF {
			c.log.Error("Client has closed the connection", "remote addr")
			break
		} else if err != nil {
			c.log.Error("Cannot read from client", "Error", err)
			return err
		}

		c.log.Info("Handling client request", "Base currency", rr.Base, "Destination currency", rr.Destination)
		rrSub, isCached := c.subscribers[src]
		if !isCached {
			rrSub = []*protos.RateRquest{}
		}
		rrSub = append(rrSub, rr)
		c.subscribers[src] = rrSub
	}
	return nil
}

func (c *Currency) handleRatesUpdates() {
	rateUpdate := c.rates.MonitorRates(5 * time.Second)
	for range rateUpdate {
		c.log.Info("Got updated rates")

		//Looping over subscribed cleints
		for key, val := range c.subscribers {
			//Looping over rates
			for _, rr := range val {
				rate, err := c.rates.GetRate(rr.Base.String(), rr.Destination.String())
				if err != nil {
					c.log.Error("Unable to get rate", "Base",rr.Base.String(), "Destination", rr.Destination.String())
				}
				key.Send(&protos.RateResponse{
					Base:rr.Base, 
					Destination: rr.Destination,
					Rate:  rate,
				})
			}
		}

	}
}
