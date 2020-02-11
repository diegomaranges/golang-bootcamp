package main

import (
	"log"
	"net/http"

	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-4/routes"
)

func main() {
	router := new(routes.Router)
	router.NewRouter()

	server := http.ListenAndServe(":8080", router.PointerToRouter())
	log.Fatal(server)
}
