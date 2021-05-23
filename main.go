package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"post-example/handler"
	"time"
)

func Middleware(r *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			next.ServeHTTP(w, req)
		})
	}
}

func main() {
	starship := handler.NewStarship()
	r := mux.NewRouter()
	r.HandleFunc("/starships", starship.StarshipsHandler)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(Middleware(r))

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
