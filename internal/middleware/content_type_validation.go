package middleware

import "net/http"

func ContentTypeValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Only validate methods that normally have a request body
		switch r.Method {
		case http.MethodPost, http.MethodPut, http.MethodPatch:
			if r.Header.Get("Content-Type") != "application/json" {
				http.Error(w, "Content-Type must be application/json please fix it and retry", http.StatusUnsupportedMediaType)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
