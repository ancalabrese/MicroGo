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
	"net/http"
	"strconv"

	"github.com/hashicorp/go-hclog"

	"github.com/ancalabrese/MicroGo/Products/data"
	"github.com/ancalabrese/MicroGo/Products/middleware"
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

// swagger: route GET /products prodects listProducts
// Returns the list of all products
// response:
// 200

//GetProducts return the products in the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Debug("Get all records")
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

//AddProducts to the dataset
func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(middleware.ProductKey{}).(data.Product)
	p.l.Debug("Add new record", "Product", prod)
	p.productsDB.AddProduct(prod)
	rw.WriteHeader(http.StatusOK)
}

//UpdateProduct with new product passed in the reuqest body
func (p Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
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
