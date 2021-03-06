package conf

import (
	"fmt"

	"github.com/Williberto14/prueba-tecnica/model"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func ConfDatabase() *gorm.DB {

	db, err = gorm.Open("sqlite3", "./database/products.db")

	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&model.Product{})

	return db
}
