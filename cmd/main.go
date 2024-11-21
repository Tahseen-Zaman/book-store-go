package main

import (
	"log"
	"net/http"

	"github.com/Tahseen-Zaman/book-store-go/pkg/models"
	"github.com/Tahseen-Zaman/book-store-go/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	models.Init()
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	log.Println("Server is running on localhost:9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
