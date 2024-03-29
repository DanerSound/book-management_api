package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/DanerSound/book-management_api/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.registerBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))// create the server
	
}

