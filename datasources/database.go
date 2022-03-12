package datasources

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	ctx := context.Background()

	var err error
	DB, err = sql.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/bookstore")
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	err = DB.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connected!")
}
