package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/luisc2023/http-handler-golang/routes"
)

func main() {
	fmt.Println("Hi Folks!")

	router := routes.MainRoutes()

	server := &http.Server{
		Addr:           "localhost:8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error server connection:", err)
	}
}
