package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/DanerSound/book-management_api/tree/main/pkg/utils"
	"github.com/DanerSound/book-management_api/tree/main/pkg/models"
)

 var NewBoook models.Book

 // getbooks
 func getBook(w http.ResponseWriter, r *http.Request){
	NewBoooks:= models.getAllBook()
	res, _ :=json.Marshal(NewBoooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
 }

 func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err:= strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	bookDetails,_:= models.GetBookById(ID)
	res, _:= json.Marshal(bookDetails)

	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook:=&models.Book{}
	utils.ParseBody(r, CreateBook)
	b:= CreateBook.CreateBook()
	res,_ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err:= strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while deleting")
	}
	book:= models.DeleteBook(ID)
	res, _:= json.Marshal(book)

	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var UpdateBook = &models.Book{}
	utils.Parse(r, UpdateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println(" error while deleting")
	}
	bookDetails, db:=models.GetBookById(ID)
	if UpdateBook.Name !=""{
		bbookDetails.Name = updateBook.Name
	}
	if updateBook.Author !=""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}
}