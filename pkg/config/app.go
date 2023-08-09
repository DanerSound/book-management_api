package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func ConnectDB(){
	d, err := gorm.Open("testdb","root:password")
	if err != nil{
		panic(err)
	}
	db = d
}

func getDB() *gorm.DB{
	return db
}