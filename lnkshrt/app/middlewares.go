package app

import (
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

func LoggingMiddleware(n http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("panic: %v\nstacktrace: %v\n", err, debug.Stack())
			}
		}()

		start := time.Now()
		n.ServeHTTP(w, r)
		log.Println(
			// w.Header().Get("Status"),
			r.Method,
			r.URL.EscapedPath(),
			time.Since(start),
		)
	}

	return http.HandlerFunc(fn)
}
