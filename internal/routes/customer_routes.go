package routes

import (
	"net/http"

	"insurance-api/internal/handlers"
)

func RegisterCustomerRoutes(
	mux *http.ServeMux,
	handler *handlers.CustomerHandler,
) {
	mux.HandleFunc("POST /customers", handler.Create)
	mux.HandleFunc("GET /customers", handler.GetAll)
	mux.HandleFunc("GET /customers/{id}", handler.GetByID)
	mux.HandleFunc("PUT /customers/{id}", handler.Update)
	mux.HandleFunc("DELETE /customers/{id}", handler.Delete)
}
