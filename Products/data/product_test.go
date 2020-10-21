package data

import "testing"

func TestCheckProductValues(t *testing.T) {

	p := &Product{
		Name:  "Frappuccino",
		Price: 1.50,
		SKU:   "P-0001",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
