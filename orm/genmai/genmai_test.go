package main

import (
	"github.com/naoina/genmai"
	"testing"
)

func BenchmarkGenmaiCreate(b *testing.B) {

	db, err := genmai.New(&genmai.MySQLDialect{}, "root:@tcp(localhost:3306)/test")
	if err != nil {
		panic(err)
	}

	db.DB().Query("TRUNCATE TABLE users")

	b.ResetTimer()

	// Transaction
	defer func() {
		if err := recover(); err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
	}()
	if err := db.Begin(); err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		_, e := db.Insert(&Users{Name: "alice", Age: 25})
		if e != nil {
			panic(e)
		}
	}
}

func BenchmarkGenmaiRead(b *testing.B) {

	db, err := genmai.New(&genmai.MySQLDialect{}, "root:@tcp(localhost:3306)/test")
	if err != nil {
		panic(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var user []Users
		if err := db.Select(&user, db.Where("id", "=", 1)); err != nil {
			panic(err)
		}
	}
}

func BenchmarkGenmaiUpdate(b *testing.B) {

	db, err := genmai.New(&genmai.MySQLDialect{}, "root:@tcp(localhost:3306)/test")
	if err != nil {
		panic(err)
	}

	b.ResetTimer()

	// Transaction
	defer func() {
		if err := recover(); err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
	}()
	if err := db.Begin(); err != nil {
		panic(err)
	}

	var users []Users
	if err := db.Select(&users, db.Where("id", "=", 1)); err != nil {
		panic(err)
	}
	user := users[0]
	user.Name = "mini"

	for i := 0; i < b.N; i++ {
		if _, err := db.Update(&user); err != nil {
			panic(err)
		}
	}
}
