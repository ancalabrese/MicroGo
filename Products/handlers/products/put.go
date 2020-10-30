package products

import (
	"net/http"

	data "github.com/ancalabrese/MicroGo/Products/data/product"
	"github.com/ancalabrese/MicroGo/Products/middleware"
)

//UpdateProduct with new product passed in the reuqest body
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	id, err := getProductId(r)
	p.l.Debug("Update record", "id", id)
	if err != nil {
		p.l.Error("Cannot find record", "error", err)
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}

	prod := r.Context().Value(middleware.ProductKey{}).(data.Product)
	err = p.productsDB.UpdateProduct(id, prod)
	if err != nil {
		p.l.Error("Cannot insert record", "error", err)
		http.Error(rw, "Product not found", http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
