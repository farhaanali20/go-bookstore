package main

import (
	"log"
	"net/http"

	routes "github.com/farhaan/go-bookstore/pkg/routers"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoute(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9010",r))
}