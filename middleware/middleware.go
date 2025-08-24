package middleware

import (
	"log"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, "executing middleware before!")
		next.ServeHTTP(w, r)
		log.Println(r.URL.Path, "executing middleware after!")
	})
}
