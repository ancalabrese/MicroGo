package middleware

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-hclog"

	data "github.com/ancalabrese/MicroGo/Products/data"
	pdata "github.com/ancalabrese/MicroGo/Products/data/product"
)

//ProductKey to retrieve the product in  the body request
type ProductKey struct{}

type Validator struct {
	log hclog.Logger
}

func NewProductValidator(l hclog.Logger) *Validator {
	return &Validator{l}
}
//Validate casts the request argument in Product and validate object
func (v *Validator) Validate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := pdata.Product{}

		e := data.FromJSON(&prod, r.Body)
		if e != nil {
			v.log.Error("Cannot deserialize received obj", "error", e)
			http.Error(rw, "Bad argument", http.StatusBadRequest)
			return
		}
		e = prod.Validate()
		if e != nil {
			v.log.Error("Object validation failed", "error", e)
			http.Error(rw, "Bad argument", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), ProductKey{}, prod)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
