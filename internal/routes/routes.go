package routes

import (
	"net/http"

	"insurance-api/internal/handlers"
)

func RegisterRoutes(
	mux *http.ServeMux,
	customerHandler *handlers.CustomerHandler,
	policyHandler *handlers.PolicyHandler,
	customerPolicyHandler *handlers.CustomerPolicyHandler,
	claimHandler *handlers.ClaimHandler,
) {
	RegisterCustomerRoutes(mux, customerHandler)
	RegisterPolicyRoutes(mux, policyHandler)
	RegisterCustomerPolicyRoutes(mux, customerPolicyHandler)
	RegisterClaimRoutes(mux, claimHandler)
}
