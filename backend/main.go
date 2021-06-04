package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/girikuncoro/belajar-docker-pemula/backend/db"
	"github.com/girikuncoro/belajar-docker-pemula/backend/handler"
	_ "github.com/lib/pq"
)

func main() {
	var postgres *db.Postgres
	var err error

	postgres, err = db.ConnectPostgres()
	if err != nil {
		panic(err)
	}
	if postgres == nil {
		panic("postgres is unreachable")
	}

	mux := handler.InitRoutes(postgres)
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
