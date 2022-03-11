package main

import (
	"github.com/octavio-luna/bookstore/datasources"
)

func main() {
	db, err := datasources.ConnectDB()
}
