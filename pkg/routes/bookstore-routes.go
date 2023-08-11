package routes

import(
	"github.com/gorilla/mux"
	"github.com/DanerSound/book-management_api/pkg/controllers"
)

var registerBookStoreRoutes = func (router *mux.Route)  {
	router.HandlerFunc("/book/",controllers.createBook).Methods("POST")
	router.HandlerFunc("/book/",controllers.GetBook).Methods("GET")
	router.HandlerFunc("/book/{bookId}",controllers.GetBookById).Methods("GET")
	router.HandlerFunc("/book/{bookId}",controllers.UpdateBook).Methods("PUT")
	router.HandlerFunc("/book/{bookId}",controllers.DeleteBook).Methods("DELETE")
}