package data

import (
	"fmt"

	model "github.com/ancalabrese/MicroGo/Currency/data/model"
	protos "github.com/ancalabrese/MicroGo/Currency/protos/currency"
	"github.com/hashicorp/go-hclog"
)

type Currencies struct {
	log  hclog.Logger
	list model.Currencies
}

func NewCurrencies(l hclog.Logger) *Currencies {
	c := model.Currencies{}
	for code, _ := range protos.Currencies_value {
		c.Codes = append(c.Codes, code)
	}
	return &Currencies{log: l, list: c}
}

func (c *Currencies) GetCurencies() ([]string, error) {
	if len(c.list.Codes) == 0 {
		c.log.Error("No currencines records found")
		return nil, fmt.Errorf("No currencines records found")
	}
	return c.list.Codes, nil
}
