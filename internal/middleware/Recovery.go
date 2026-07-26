package middleware

import (
	"insurance-api/internal/config"
	"insurance-api/internal/logger"
	"log"
	"net/http"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {

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

			if err := recover(); err != nil {

				appLogger.Info("PANIC RECOVERED: %v", err)

				http.Error(
					w,
					"Internal Server Error",
					http.StatusInternalServerError,
				)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
