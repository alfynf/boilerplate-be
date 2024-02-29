package main

import (
	m "boilerplate/middlewares"
	"boilerplate/routes"
	"net/http"
	"time"
)

func main() {
	r := routes.New()
	m.LogMiddleware(r)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
