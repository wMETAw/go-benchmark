package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Users struct {
	Id   int64  `json:"id" xorm:"'id'"`
	Name string `json:"name" xorm:"'name'"`
	Age  int8   `json:"age" xorm:"'age'"`
}

var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:@/test")
	if err != nil {
		panic("error failed to connect DB")
	}

	user := new(Users)
	user.Name = "Sato"
	user.Age = 25
	engine.Insert(user)

	var u Users
	engine.Id(1).Get(&u)
	engine.Get(&u)

	fmt.Println(u)
}
