package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adiiverma40/BookStore/pkg/routes"
	"github.com/gorilla/mux"
)



func main(){
	fmt.Println("Welcome to the book store!")

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:9010", r))


}
