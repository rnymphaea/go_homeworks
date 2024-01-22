package main

import (
	"log"
	"net/http"
	"simple_server/pkg/api"
)

func main() {
	api := api.New("localhost:8090", http.NewServeMux())
	api.FillEndpoints()
	log.Fatal(api.ListenAndServe())
}
