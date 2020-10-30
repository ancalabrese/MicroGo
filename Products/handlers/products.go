// Package classification of Product API
//
// Documenatation for Product API
//
// Schemes: http
// BasePath : /
// Version: 1.0.0
//
// Consumes:
// 	 - application/json
//
// Produces:
// 	- application/json
// swagger:meta

package handlers

import (
	"net/url"
	"net/http"
	"strconv"

	"github.com/hashicorp/go-hclog"

	"github.com/ancalabrese/MicroGo/Products/data"
	"github.com/gorilla/mux"
)

// A list of products in the API response
type productsReponse struct {
	//All products in the data store
	// in: body
	Body []data.Product
}

type Products struct {
	l          hclog.Logger
	productsDB *data.ProductsDB
}

//NewProducts returns a new instance of Products
func NewProducts(l hclog.Logger, pdb *data.ProductsDB) *Products {
	return &Products{l, pdb}
}

func (p Products) operationNotSupported(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "Method not supported", http.StatusMethodNotAllowed)
}

func getProductId(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return id, err
	}
	return id, nil
}

func getCurrency(url *url.URL) string{
	return url.Query().Get("currency")
}
