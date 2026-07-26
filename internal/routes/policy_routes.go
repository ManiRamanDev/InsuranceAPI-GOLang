package routes

import (
	"net/http"

	"insurance-api/internal/handlers"
)

func RegisterPolicyRoutes(
	mux *http.ServeMux,
	handler *handlers.PolicyHandler,
) {
	mux.HandleFunc("POST /policies", handler.Create)
	mux.HandleFunc("GET /policies", handler.GetAll)
	mux.HandleFunc("GET /policies/{id}", handler.GetByID)
	mux.HandleFunc("PUT /policies/{id}", handler.Update)
	mux.HandleFunc("DELETE /policies/{id}", handler.Delete)
}
