package main

//go:binary-only-package

import (
	"context"
	"log"

	"github.com/albert-widi/sqlimporter"
	_ "github.com/lib/pq"
	// "github.com/tokopedia/logistic/pkg/testutil/sqlimporter"
)

func main() {
	dsn := "postgres://logistic:logistic@localhost:5432?sslmode=disable"
	// dsn := "user=exampleapp password=exampleapp host=127.0.0.1:5432 dbname=exampleapp sslmode=disable"
	// db, err := sqlx.Open("postgres", dsn)
	db, drop, err := sqlimporter.CreateRandomDB("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = sqlimporter.ImportSchemaFromFiles(context.TODO(), db, "../files")
	if err != nil {
		log.Fatal("Failed to import ", err.Error())
	}
	defer func() {
		err := drop()
		if err != nil {
			log.Printf("Failed to drop. Error: %s", err.Error())
		}
	}()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("finished")
}
