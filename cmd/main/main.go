package main

import (
"log"
"net/http"
"github.com/gorilla/mux"
_ "github.com/jinzhu/gorm/dialects/mysql"
"github.com/vadiraj/pro3/pkg/routes"
)

func main() {
	r:=mux.NewRouter()
    routes.RegisterBookStoreRoutes(r)
    http.Handle("/",r)
    log.Fatal(http.ListenAndServe(":3000",r))
}