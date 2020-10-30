package products

import (
	"net/http"

	data "github.com/ancalabrese/MicroGo/Products/data/product"
	"github.com/ancalabrese/MicroGo/Products/middleware"
)

//AddProducts to the dataset
func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(middleware.ProductKey{}).(data.Product)
	p.l.Debug("Add new record", "Product", prod)
	p.productsDB.AddProduct(prod)
	rw.WriteHeader(http.StatusOK)
}
