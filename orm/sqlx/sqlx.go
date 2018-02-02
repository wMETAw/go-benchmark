package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE users (
    id integer,
    created_at date,
    updated_at date,
    deleted_at date,
    name text,
    age integer
);`

type User struct {
	Name string `db:"name"`
	Age  int8   `db:"age"`
}

func main() {
	db, err := sqlx.Connect("mysql", "root:@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatalln(err)
	}

	//db.MustExec(schema)

	tx := db.MustBegin()

	// INSERT
	tx.MustExec("INSERT INTO users (name, age) VALUES (?, ?)", "Yamada", 25)

	// 構造体を指定してINSERT
	tx.NamedExec("INSERT INTO users (name, age) VALUES (:name, :age)", &User{"Suzuki", 30})

	// Get
	// => {Yamada 25}
	user := User{}
	db.Get(&user, "SELECT name, age FROM users WHERE id = ?", 1)

	// Update
	_, e := db.Exec("UPDATE users SET name = ? WHERE id = ?", "Tanaka", 1)
	if e != nil {
		tx.Rollback()
	}

	tx.Commit()
}
