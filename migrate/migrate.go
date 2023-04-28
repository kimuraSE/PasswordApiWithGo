package main

import (
	"PasswordApiWithGo/db"
	"PasswordApiWithGo/model"
	"fmt"
)

func main() {

	dbConn := db.NewDB()
	defer fmt.Println("Migrate Successfully")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Title{}, &model.Password{})
}
