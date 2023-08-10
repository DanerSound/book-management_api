package models

import(
	"github.com/jinzhu/gorm"
	"github.com/DanerSound/book-management_api/tree/main/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json: "author"`
	Publication string `json:"publication"`
}

func init(){
	config.ConnectDB()
	db = config.getDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) createBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

 func getAllBook() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
 }

 func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db:= db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
 }

 func DeleteBook(ID int64) Book{
	var book Book
	db.Where("ID=?",ID).Delete(book)
	return book
 }

 // how to update, when the user will sent data 
 // to update a particolar record, you will find the record
 // delete it , and recreate the record with the new data
// GET DELETE POST 
