package integration

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"insurance-api/internal/database"
	"insurance-api/internal/handlers"
	"insurance-api/internal/repository"
	"insurance-api/internal/routes"
	"insurance-api/internal/service"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func newIntegrationAppBenchmark(b *testing.B) *integrationApp {
	b.Helper()

	dbFile := filepath.Join(b.TempDir(), "benchmark.db")
	dsn := fmt.Sprintf("file:%s?_foreign_keys=1", dbFile)
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		b.Fatalf("open sqlite db: %v", err)
	}

	if err := database.InitSchema(db); err != nil {
		b.Fatalf("init schema: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Insurance API is running"))
	})

	customerRepo := repository.NewCustomerRepository(db)
	policyRepo := repository.NewPolicyRepository(db)
	customerPolRepo := repository.NewCustomerPolicyRepository(db)
	claimRepo := repository.NewClaimRepository(db)

	routes.RegisterRoutes(
		mux,
		handlers.NewCustomerHandler(service.NewCustomerService(customerRepo)),
		handlers.NewPolicyHandler(service.NewPolicyService(policyRepo)),
		handlers.NewCustomerPolicyHandler(service.NewCustomerPolicyService(customerPolRepo)),
		handlers.NewClaimHandler(service.NewClaimService(claimRepo)),
	)

	server := httptest.NewServer(mux)

	return &integrationApp{
		server:          server,
		db:              db,
		customerRepo:    customerRepo,
		policyRepo:      policyRepo,
		customerPolRepo: customerPolRepo,
		claimRepo:       claimRepo,
	}
}

func BenchmarkIntegrationHealthEndpoint(b *testing.B) {
	app := newIntegrationAppBenchmark(b)
	defer app.close()

	client := app.server.Client()
	url := app.server.URL + "/health"

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		resp, err := client.Get(url)
		if err != nil {
			b.Fatalf("health request failed: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			b.Fatalf("expected %d, got %d", http.StatusOK, resp.StatusCode)
		}
		if _, err := io.Copy(io.Discard, resp.Body); err != nil {
			resp.Body.Close()
			b.Fatalf("read body failed: %v", err)
		}
		if err := resp.Body.Close(); err != nil {
			b.Fatalf("close body failed: %v", err)
		}
	}
}
