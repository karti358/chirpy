package main

import (
	"net/http"
	"fmt"
	"github.com/karti358/chirpy/router"
)

func main() {

	var handler *http.ServeMux = http.NewServeMux()
	handler.Handle(
		"/api/",
		http.StripPrefix(
			"/api",
			router.APIHandler,
		),
	)

	var server *http.Server = &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server Stopped")
	}

}