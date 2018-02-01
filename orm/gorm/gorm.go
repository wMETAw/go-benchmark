package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type User struct {
	gorm.Model
	Name string
	Age int8
}

func main() {
	db, err := gorm.Open("mysql", "root@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// migrate the schema
	db.AutoMigrate(&User{})

	// create
	db.Create(&User{Name: "Yamada", Age:20})

	var u User
	db.First(&u,1) // find product with id 1

	//db.Delete(&User{},4)

	fmt.Println(u)
}