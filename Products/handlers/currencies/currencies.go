package currencies

import (
	"net/http"

	"github.com/ancalabrese/MicroGo/Products/data"
	currencies "github.com/ancalabrese/MicroGo/Products/data/currency"
	"github.com/hashicorp/go-hclog"
)

type Currencies struct {
	log        hclog.Logger
	currencyDB *currencies.CurrencyDB
}

func NewCurrencyH(l hclog.Logger, cdb *currencies.CurrencyDB) *Currencies {
	return &Currencies{log: l, currencyDB: cdb}
}

//GetCurrencies returns all supported currencies
func (c *Currencies) GetCurrencies(rw http.ResponseWriter, r *http.Request) {
	c.log.Info("Handling GetCurrencies")
	currencies, err := c.currencyDB.GetCurrencies()
	if err != nil {
		c.log.Error("Cannot retrieve currencies", "error", err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	data.ToJSON(currencies, rw)
}
