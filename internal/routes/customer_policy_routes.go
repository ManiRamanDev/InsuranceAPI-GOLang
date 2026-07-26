package routes

import (
	"net/http"

	"insurance-api/internal/handlers"
)

func RegisterCustomerPolicyRoutes(
	mux *http.ServeMux,
	handler *handlers.CustomerPolicyHandler,
) {
	mux.HandleFunc("POST /customer-policies", handler.Create)
	mux.HandleFunc("GET /customer-policies", handler.GetAll)
	mux.HandleFunc("GET /customer-policies/{id}", handler.GetByID)
	mux.HandleFunc("PUT /customer-policies/{id}", handler.Update)
	mux.HandleFunc("DELETE /customer-policies/{id}", handler.Delete)
}
