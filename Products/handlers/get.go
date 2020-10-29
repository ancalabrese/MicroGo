package handlers

import (
	"net/http"

	"github.com/ancalabrese/MicroGo/Products/data"
)

// swagger: route GET /products prodects listProducts
// Returns the list of all products
// response:
// 200
//GetProducts return the products in the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Debug("Get all records")
	rw.Header().Add("Content-Type", "application/json")
	products, err := p.productsDB.GetProducts("")
	err = data.ToJSON(products, rw)
	if err != nil {
		p.l.Error("Unable to serialise records", "error", err)
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}
}

//GetProduct returns a single product by its ID
func (p *Products) GetProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	id, err := getProductId(r)
	p.l.Debug("Get record", "id", id)
	if err != nil {
		p.l.Error("Cannot find record", "error", err)
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}
	prod, err := p.productsDB.GetProduct(id, "")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
	}
	err = data.ToJSON(prod, rw)
	if err != nil {
		p.l.Error("Unable to serialise record", "error", err)
		http.Error(rw, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
