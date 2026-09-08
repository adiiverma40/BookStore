package routes

import "github.com/gorilla/mux"

var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controller.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookid}", controller.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookid}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookid}", controller.DeleteBook).Methods("DELETE")
	
	

	
}