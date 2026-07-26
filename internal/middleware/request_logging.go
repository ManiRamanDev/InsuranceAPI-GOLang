package middleware

import (
	"insurance-api/internal/config"
	"insurance-api/internal/logger"
	"log"
	"net/http"
	"time"
)

func RequestLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Load Configurations
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

		start := time.Now()

		duration := time.Since(start)

		appLogger.Info("Method= " + r.Method + ", UrlPath= " + r.URL.Path + ", Duration= " + duration.String())

		next.ServeHTTP(w, r)
	})
}
