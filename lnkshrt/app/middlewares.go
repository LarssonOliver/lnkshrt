package app

import (
	"log"
	"net/http"
)

func LoggingMiddleware(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		n.ServeHTTP(w, r)
	})
}
