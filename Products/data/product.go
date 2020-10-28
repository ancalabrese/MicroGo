package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

//Error Types
var ErrProductNotFound = fmt.Errorf("Product ID not found")

//Defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	SKU         string  `json:"sku" validate:"required,sku"`
	Price       float32 `json:"price" validate:"required"`
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

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func GetProducts() Products {
	return products
}

func AddProducts(p Products) {
	for _, newP := range p {
		newP.ID = getNextId()
		products = append(products, newP)
	}
}

func AddProduct(p Product) {
	p.ID = getNextId()
	products = append(products, &p)
}

func UpdateProduct(id int, newProduct *Product) error {
	_, pos, err := findProductById(id)
	if err != nil {
		return err
	}
	newProduct.ID = id
	products[pos] = newProduct
	return nil
}

func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	err := d.Decode(p)
	return err
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
