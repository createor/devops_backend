package main

import (
	"net/http"

	route "server/views"
)

func main() {
	s := &http.Server{
		Addr:    ":1521",
		Handler: route.InitRouter(),
	}

	s.ListenAndServe()
}
