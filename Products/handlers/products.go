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
	"log"
	"net/http"
	"strconv"

	proto "github.com/ancalabrese/MicroGo/Currency/protos/currency"
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
	l  *log.Logger
	cc proto.CurrencyClient
}

//NewProducts returns a new instance of Products
func NewProducts(l *log.Logger, cc proto.CurrencyClient) *Products {
	return &Products{l, cc}
}

// swagger: route GET /products prodects listProducts
// Returns the list of all products
// response:
// 200

//GetProducts return the products in the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := products.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}
}

//GetProduct returns a single product by its ID
func (p *Products) GetProduct(rw http.ResponseWriter, r *http.Request) {
	id, err := getProductId(r)
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}
	prod, _, err := data.GetProduct(id)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
	}
	err = prod.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
	}
}

//AddProducts to the dataset
func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(middleware.ProductKey{}).(data.Product)
	data.AddProduct(prod)
	rw.WriteHeader(http.StatusOK)
}

//UpdateProduct with new product passed in the reuqest body
func (p Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	id, err := getProductId(r)
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}

	prod := r.Context().Value(middleware.ProductKey{}).(data.Product)
	e := data.UpdateProduct(id, &prod)
	if e != nil {
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

// func (p *Products) grabURLProdID(url string) int {
// 	reg := regexp.MustCompile(`/([0-9+])`)
// 	subMatches := reg.FindAllStringSubmatch(url, -1)
// 	if len(subMatches) == 0 || len(subMatches) > 1 || len(subMatches[0]) < 2 || len(subMatches[0]) > 2 {
// 		p.l.Println("Invalid URL: (PUT) -> ", url)
// 		return -1
// 	}
// 	id, err := strconv.Atoi(subMatches[0][1])
// 	if err != nil {
// 		p.l.Println("Invalid URL: (PUT) -> ", url)
// 		return -1
// 	}
// 	return id
// }
