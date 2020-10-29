package handlers

import "net/http"

func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	id, err := getProductId(r)
	p.l.Debug("Delete record", "id", id)
	if err != nil {
		p.l.Error("Cannot find record", "error", err)
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}
	err = p.productsDB.DeleteProduct(id)
	if err != nil {
		p.l.Error("Cannot delete record", "error", err)
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
