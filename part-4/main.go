package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-4/db"

	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-4/router"
)

func backup() {
	for {
		time.Sleep(12 * time.Hour)
		fmt.Println(db.Backup())
	}
}

func main() {
	router := router.NewRoute()

	go backup()

	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}
