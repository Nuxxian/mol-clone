package main

import (
	"net/http"
	"os"
)

func main() {
	router := http.NewServeMux()

	server := http.Server{Addr: ":" + os.Getenv("PORT"), Handler: router}
    server.ListenAndServe()
	return
}
