package main

import (
	"database/sql"

	"github.com/ffelipelimao/ports-adapters-architeture/adapter/db"
	"github.com/ffelipelimao/ports-adapters-architeture/application"
)

func main() {
	DB, _ := sql.Open("sqlite3", "db.sqlite")
	productDBAdapter := db.NewProductDB(DB)
	productService := application.NewProductService(productDBAdapter)
	productService.Create("product example", 30)
}
