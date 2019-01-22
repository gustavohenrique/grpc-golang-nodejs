package product

import (
	"net/http"
	"encoding/json"
)

type Handler struct{
	service *ProductService
}

func NewHandler(service *ProductService) *Handler {
	return &Handler{service: service}
}

func (this *Handler) FetchAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	products, err := this.service.FetchAllProducts()
	if err != nil {
		http.Error(w, "Failed getting the product list.", http.StatusInternalServerError)
		return
	}

	id := req.Header.Get("X-USER-ID")
	customer, err := this.service.FindCustomerByID(id)
	if err != nil {
		json.NewEncoder(w).Encode(products)
		return
	}

	productsWithDiscountApplied, err := this.service.ApplyDiscount(customer, products)
	if err == nil {
		json.NewEncoder(w).Encode(productsWithDiscountApplied)
		return
	}
	json.NewEncoder(w).Encode(products)
}
