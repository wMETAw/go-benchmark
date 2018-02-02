package main

import (
	"github.com/jmoiron/sqlx"
	"log"
	"testing"
)

func BenchmarkSqlxCreate(b *testing.B) {

	db := dbConnect()
	db.Exec("TRUNCATE TABLE users")
	b.ResetTimer()

	tx := db.MustBegin()
	for i := 0; i < b.N; i++ {
		tx.NamedExec("INSERT INTO users (name, age) VALUES (:name, :age)", &User{"Suzuki", 30})
	}
	tx.Commit()
}

func BenchmarkSqlxRead(b *testing.B) {
	db := dbConnect()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		user := User{}
		db.Get(&user, "SELECT name, age FROM users WHERE id = ?", 1)
	}
}

func BenchmarkSqlxUpdate(b *testing.B) {
	db := dbConnect()
	b.ResetTimer()

	tx := db.MustBegin()
	for i := 0; i < b.N; i++ {
		_, e := db.Exec("UPDATE users SET name = ? WHERE id = ?", "Tanaka", 1)
		if e != nil {
			tx.Rollback()
		}
	}
	tx.Commit()
}

func dbConnect() (db *sqlx.DB) {
	db, err := sqlx.Connect("mysql", "root:@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatalln(err)
	}
	return
}
