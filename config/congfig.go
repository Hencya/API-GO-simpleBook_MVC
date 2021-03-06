package config

import (
	Models "Book_CRUD/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB(){
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "root",
		"DB_Name": "rent_book",
	}

	connectingString := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Name"],
		)

	var err error
	DB,err = gorm.Open(mysql.Open(connectingString),&gorm.Config{})
	if err != nil{
		panic(err)
	}

	initialMigration()
}

func initialMigration(){
	DB.AutoMigrate(&Models.Books{})
}