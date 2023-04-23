package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sugartr3e/go-seeder"
	"log"
)

func main() {
	dsn := "sample:password@tcp(localhost:3306)/sample?charset=utf8mb4&parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	s := seeder.NewSeeder(db)
	err = s.Table("test").Insert(map[string]any{
		"name": "sample",
		"age":  20,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("success")
}
