package routes

import(
	"github.com/gorilla/mux"
	"github.com/DanerSound/book-management_api/blob/main/pkg/controllers"
)

var registerBookStoreRoutes = func (router *mux.Route)  {
	router.HandlerFunc("/book/",controllers.createBook).Methods("POST")
	router.HandleFunc("/book/",controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}",controllers.DeleteBook).Methods("DELETE")
}