package main

import (
	"testing"
	"github.com/jinzhu/gorm"
)

func BenchmarkGinGormCreate(b *testing.B){

	db := dbConnect()
	db.Exec("TRUNCATE TABLE users;")
	b.ResetTimer()

	for i:=0;i<b.N ;i++  {
		tx := db.Begin()

		// create
		if err := tx.Create(&User{Name:"Sato",Age:30}).Error; err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}
}

func BenchmarkGinGormRead(b *testing.B){

	db := dbConnect()
	b.ResetTimer()

	for i:=0;i<b.N ;i++  {
		db.First(&User{},1) // find User with id 1
	}
}


func BenchmarkGinGormUpdate(b *testing.B){

	db := dbConnect()
	b.ResetTimer()

	for i:=0;i<b.N ;i++  {
		tx := db.Begin()

		if err := tx.Model(&User{}).Where("id = ?", 1).Update("name","Suzuki").Error; err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}
}

func dbConnect() (db *gorm.DB){
	db, err := gorm.Open("mysql", "root@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	return
}
