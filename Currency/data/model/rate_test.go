package model

import (
	"bytes"

	"testing"
)

func TestRate(t *testing.T) {
	s := []byte(`{"rates": {"CAD": 1.5589,"HKD": 9.1699,"ISK": 165.3},"base": "EUR","date": "2020-10-27"}`)
	r := &Rates{}
	err := r.FromJson(bytes.NewReader(s))
	if err != nil {
		t.Errorf("Error unmarshaling json %#v", err)
	}

	if _, found := r.Rates["CAD"]; !found {
		t.Errorf("Key not found %#vf", r)
	}
}
