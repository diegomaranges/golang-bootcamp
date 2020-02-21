package main

import (
	"log"
	"net/http"

	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-4/router"
)

func main() {
	router := router.NewRute()

	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}
