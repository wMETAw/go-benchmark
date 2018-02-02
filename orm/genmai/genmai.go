package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/naoina/genmai"
	"time"
)

// define a table schema.
type Users struct {
	Id         int64 `db:"pk"`
	Created_at *time.Time
	Updated_at *time.Time
	Deleted_at *time.Time
	Name       string `default:"me"`
	Age        int8
}

func main() {
	db, err := genmai.New(&genmai.MySQLDialect{}, "root:@tcp(localhost:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Insert
	obj := &Users{Name: "alice", Age: 25}
	n, err := db.Insert(obj)
	if err != nil {
		panic(err)
	}
	fmt.Printf("inserted rows: %d\n", n)

	// select
	var user []Users
	if err := db.Select(&user, db.Where("id", "=", 1)); err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", user)

	// update
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
