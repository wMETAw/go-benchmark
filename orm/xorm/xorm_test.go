package main

import (
	"github.com/go-xorm/xorm"
	"testing"
)

func BenchmarkXormCreate(b *testing.B) {

	engine := dbConnect()
	engine.Query("TRUNCATE TABLE users")

	b.ResetTimer()

	session := engine.NewSession()
	defer session.Close()
	session.Begin()

	var err error
	for i := 0; i < b.N; i++ {
		_, err = session.Insert(Users{Name: "kimura", Age: 25})
		if err != nil {
			session.Rollback()
			return
		}
	}

	err = session.Commit()
	if err != nil {
		return
	}
}

func BenchmarkXormRead(b *testing.B) {

	engine := dbConnect()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var user = Users{Id: 1}
		engine.Get(&user)
	}
}

func BenchmarkXormUpdate(b *testing.B) {

	engine := dbConnect()
	b.ResetTimer()

	session := engine.NewSession()
	defer session.Close()
	session.Begin()

	var err error
	for i := 0; i < b.N; i++ {

		user := Users{Name: "Tanaka"}
		_, err = session.Where("id = ?", 1).Update(&user)
		if err != nil {
			session.Rollback()
			return
		}
	}

	err = session.Commit()
	if err != nil {
		return
	}
}

func dbConnect() (engine *xorm.Engine) {

	engine, err := xorm.NewEngine("mysql", "root:@/test")
	if err != nil {
		panic("error failed to connect DB")
	}
	return
}
