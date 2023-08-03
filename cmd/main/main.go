package main

import (
	"log"
	"net/http"
	"strconv"

	"bookstore/pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
	f := func(a int) string {
		return "" + strconv.Itoa(a)
	}
	f(1)
}
