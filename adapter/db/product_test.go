package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/ffelipelimao/ports-adapters-architeture/adapter/db"
	"github.com/ffelipelimao/ports-adapters-architeture/application"
	"github.com/stretchr/testify/assert"
)

var DB *sql.DB

func setup() {
	DB, _ = sql.Open("sqlite3", ":memory:")
	createTable(DB)
	createProduct(DB)
}

func createTable(db *sql.DB) {
	table := `create table products (
		id string,
		name string,
		price float,
		status string
	);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values ("abc","Products Test",0,"disable")`

	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDB_Get(t *testing.T) {
	setup()
	defer DB.Close()

	productDB := db.NewProductDB(DB)
	product, err := productDB.Get("abc")

	assert.Nil(t, err)
	assert.Equal(t, "Products Test", product.GetName())
	assert.Equal(t, "disable", product.GetStatus())
}

func TestProductDB_Save(t *testing.T) {
	setup()
	defer DB.Close()

	productDB := db.NewProductDB(DB)
	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25

	productResult, err := productDB.Save(product)
	assert.Nil(t, err)
	assert.Equal(t, product.Name, productResult.GetName())
	assert.Equal(t, product.Status, productResult.GetStatus())

	product.Status = "enable"
	productResult, err = productDB.Save(product)
	assert.Nil(t, err)
	assert.Equal(t, product.Status, productResult.GetStatus())
}
