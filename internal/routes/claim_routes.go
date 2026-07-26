package routes

import (
	"net/http"

	"insurance-api/internal/handlers"
	"insurance-api/internal/middleware"
)

func RegisterClaimRoutes(
	mux *http.ServeMux,
	handler *handlers.ClaimHandler,
) {
	mux.Handle("POST /claims", middleware.ContentTypeValidation(middleware.RequestLogging(http.HandlerFunc(handler.Create))))
	mux.Handle("GET /claims", middleware.RequestLogging(middleware.Recovery(http.HandlerFunc(handler.GetAll))))
	mux.HandleFunc("GET /claims/{id}", handler.GetByID)
	mux.Handle("PUT /claims/{id}", middleware.ContentTypeValidation(http.HandlerFunc(handler.Update)))
	mux.Handle("DELETE /claims/{id}", middleware.ContentTypeValidation(http.HandlerFunc(handler.Delete)))
}
