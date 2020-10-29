package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	protos "github.com/ancalabrese/MicroGo/Currency/protos/currency"
	"github.com/go-playground/validator"
	"github.com/hashicorp/go-hclog"
)

//Error Types
var ErrProductNotFound = fmt.Errorf("Product ID not found")

//Defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	SKU         string  `json:"sku" validate:"required,sku"`
	Price       float64 `json:"price" validate:"required"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

var products = Products{
	&Product{
		ID:          1,
		Name:        "Espresso",
		Description: "Italian style coffee",
		SKU:         "P-00001",
		Price:       1.00,
		CreatedOn:   time.Now().String(),
		UpdatedOn:   time.Now().String(),
	},
	&Product{
		ID:          2,
		Name:        "Cappuccino",
		Description: "Coffee with milk",
		SKU:         "P-00002",
		Price:       3.00,
		CreatedOn:   time.Now().String(),
		UpdatedOn:   time.Now().String(),
	},
}

type Products []*Product

type ProductsDB struct {
	log            hclog.Logger
	currencyClient protos.CurrencyClient
}

func NewProductsDB(l hclog.Logger, cc protos.CurrencyClient) *ProductsDB {
	return &ProductsDB{l, cc}
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

//GetProducts return every product in the data base
func (p *ProductsDB) GetProducts(currency string) (Products, error) {
	if currency == "" {
		return products, nil
	}
	rate, err := p.getRate(currency)
	if err != nil {
		p.log.Error("Unable to get rate", "currency requested", currency, "error", err)
		return nil, err
	}
	pr := Products{}
	for i, _ := range products {
		pr = append(pr, p.convertPrice(rate, i))
	}
	return pr, nil
}

//GetProduct returns the position of the product if found
func (p *ProductsDB) GetProduct(id int, currency string) (*Product, error) {
	_, pos, err := findProductById(id)
	if err != nil {
		p.log.Error("Unable to find product", "id", id, "error", err)
		return nil, err
	}
	if currency == "" {
		return products[pos], nil
	}
	rate, err := p.getRate(currency)
	if err != nil {
		p.log.Error("Unable to get rate", "currency requested", currency, "error", err)
		return nil, err
	}
	return p.convertPrice(rate, pos), nil
}

// DeleteProduct deletes a product from the database
func (p *ProductsDB) DeleteProduct(id int) error {
	_, pos, err := findProductById(id)
	if err != nil {
		p.log.Error("Unable to find product", "id", id, "error", err)
		return err
	}

	products = append(products[:pos], products[pos+1])
	return nil
}

//AddProducts adds new products to the DB in bulk
func (p *ProductsDB) AddProducts(prods Products) {
	for _, newP := range prods {
		p.AddProduct(*newP)
	}
}

//AddProducts adds new single product to the DB
func (p *ProductsDB) AddProduct(pr Product) {
	pr.ID = getNextId()
	products = append(products, &pr)
}

func (p *ProductsDB) UpdateProduct(id int, newProduct Product) error {
	_, pos, err := findProductById(id)
	if err != nil {
		p.log.Error("Unable to find product", "id", id, "error", err)
		return err
	}
	newProduct.ID = id
	products[pos] = &newProduct
	return nil
}


func getNextId() int {
	return products[len(products)-1].ID + 1
}

func findProductById(id int) (*Product, int, error) {
	for i, p := range products {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func validateSKU(f validator.FieldLevel) bool {
	rg := regexp.MustCompile(`[A-Z]-[0-9]+`)
	matches := rg.FindAllString(f.Field().String(), -1)
	if matches == nil || len(matches) == 0 {
		return false
	}
	return true
}

func (p *ProductsDB) getRate(currency string) (float64, error) {
	rr := &protos.RateRquest{
		Base:        protos.Currencies(protos.Currencies_value["EUR"]),
		Destination: protos.Currencies(protos.Currencies_value[currency]),
	}
	response, err := p.currencyClient.GetRate(context.Background(), rr)
	return response.Rate, err
}

func (p *ProductsDB) convertPrice(rate float64, index int) *Product {
	np := *products[index]
	np.Price = np.Price * rate
	return &np
}
