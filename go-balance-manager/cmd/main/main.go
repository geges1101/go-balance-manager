package main

import (
	"log"
	"net/http"

	"github.com/geges1101/go-balance-manager/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBalanceRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3306", r))
}
