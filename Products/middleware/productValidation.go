package middleware

import (
	"context"
	"net/http"

	"github.com/ancalabrese/MicroGo/Products/data"
)

//ProductKey to retrieve the product in  the body request
type ProductKey struct{}

//Validate casts the request argument in Product and validate object
func Validate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		e := prod.FromJSON(r.Body)
		if e != nil {
			http.Error(rw, "Bad argument", http.StatusBadRequest)
			return
		}
		e = prod.Validate()
		if e != nil {
			http.Error(rw, "Bad argument", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), ProductKey{}, prod)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
