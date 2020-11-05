package currencies

import (
	"context"

	protos "github.com/ancalabrese/MicroGo/Currency/protos/currency"
	"github.com/hashicorp/go-hclog"
)

//TODO: Create currency obj and cache currencies. 
//TODO: Have a bi direction stram to be notified if more currencies are added to refresh cache

type CurrencyDB struct {
	log            hclog.Logger
	currencyClient protos.CurrencyClient
}

func NewCurrencyDB(l hclog.Logger, cc protos.CurrencyClient) *CurrencyDB {
	return &CurrencyDB{log: l, currencyClient: cc}
}

//GetCurrencies return all the available currencies from the currency server
func (cdb *CurrencyDB) GetCurrencies() ([]string, error) {
	cdb.log.Info("Getting all supported currencies")
	cr := &protos.CurrenciesRequest{}
	resp, err := cdb.currencyClient.GetCurrencyCodes(context.Background(), cr)
	if err != nil {
		cdb.log.Error("Cannot retrieve currencies", "error", err)
		return nil, err
	}
	return resp.Currencies, nil
}
