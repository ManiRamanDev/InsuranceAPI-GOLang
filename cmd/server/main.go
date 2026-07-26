package main

import (
	"fmt"
	"insurance-api/internal/config"
	"insurance-api/internal/database"
	"insurance-api/internal/handlers"
	"insurance-api/internal/logger"
	"insurance-api/internal/models"
	"insurance-api/internal/repository"
	"insurance-api/internal/routes"
	"insurance-api/internal/service"
	"log"
	"net/http"
)

func main() {

	// Load Configurations
	apiConfig, err := config.LoadAPIConfig()

	fmt.Println("API Config:", apiConfig)
	if err != nil {
		log.Fatal(err)
	}

	dbConfig, err := config.LoadDatabaseConfig()
	if err != nil {
		log.Fatal(err)
	}

	logConfig, err := config.LoadLogConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Logger
	appLogger, err := logger.New(logConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer appLogger.Close()

	appLogger.Info("Starting Insurance API...")

	// Connect Database
	db, err := database.Connect(dbConfig)
	if err != nil {
		appLogger.Error("Database connection failed", err)
		log.Fatal(err)
	}

	err = db.AutoMigrate(
		&models.Customer{},
		&models.Policy{},
		&models.CustomerPolicy{},
		&models.Claim{},
	)

	if err != nil {
		appLogger.Error("Database migration failed", err)
		log.Fatal(err)
	}

	// Prevent "declared and not used"
	_ = db

	// Create Router
	mux := http.NewServeMux()

	customerRepo := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepo)
	customerHandler := handlers.NewCustomerHandler(customerService)

	policyRepo := repository.NewPolicyRepository(db)
	policyService := service.NewPolicyService(policyRepo)
	policyHandler := handlers.NewPolicyHandler(policyService)

	customerPolicyRepo := repository.NewCustomerPolicyRepository(db)
	customerPolicyService := service.NewCustomerPolicyService(customerPolicyRepo)
	customerPolicyHandler := handlers.NewCustomerPolicyHandler(customerPolicyService)

	claimRepo := repository.NewClaimRepository(db)
	claimService := service.NewClaimService(claimRepo)
	claimHandler := handlers.NewClaimHandler(claimService)

	routes.RegisterRoutes(mux,
		customerHandler,
		policyHandler,
		customerPolicyHandler,
		claimHandler)

	// Temporary Health Endpoint
	registerHealthEndpoint(mux)

	appLogger.Info("Server listening on port " + apiConfig.HTTPPort)

	/*if err := http.ListenAndServe(":"+apiConfig.HTTPPort, mux); err != nil {
		appLogger.Error("Server stopped", err)
		log.Fatal(err)
	}*/

	if err := http.ListenAndServeTLS("insurance-api.company.local:8443",
		"C:/certificates/RootSSL.crt",
		"C:/certificates/RootSSL.key", mux); err != nil {
		appLogger.Error("Server stopped", err)
		log.Fatal(err)
	}
}

func registerHealthEndpoint(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Insurance API is running"))
	})
}
