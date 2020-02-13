package main

import (
	"log"
	"net/http"

	"github.corp.globant.com/diego-maranges/golang-bootcamp/part-4/routes"
)

func main() {
	router := routes.CreateNewRouter()
	router.NewRouter()

	server := http.ListenAndServe(":8080", router.PointerToRouter())
	log.Fatal(server)
}
