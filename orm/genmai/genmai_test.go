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
	defer db.Close()

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
	defer db.Close()

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
	defer db.Close()

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

		var u []Users
		if err := db.Select(&u); err != nil {
			panic(err)
		}
		o := u[0]
		o.Name = "nico"
		if _, err := db.Update(&o); err != nil {
			panic(err)
		}
	}
}
