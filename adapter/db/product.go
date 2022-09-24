package db

import (
	"database/sql"

	"github.com/ffelipelimao/ports-adapters-architecture/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (p *ProductDB) Get(ID string) (application.IProduct, error) {
	var product application.Product

	stmt, err := p.db.Prepare("select id, name, price, status from products where id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(ID).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDB) Save(product application.IProduct) (application.IProduct, error) {
	var rows int
	p.db.QueryRow("select id from products where id=?", product.GetID()).Scan(&rows)
	if rows == 0 {
		_, err := p.create(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	}
	return product, nil
}

func (p *ProductDB) create(product application.IProduct) (application.IProduct, error) {
	stmt, err := p.db.Prepare(`insert into products(id, name, price, status) values (?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetPrice(),
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDB) update(product application.IProduct) (application.IProduct, error) {
	_, err := p.db.Exec(`update products set name=?, price=?, status=? where id = ?`, product.GetName(),
		product.GetPrice(),
		product.GetPrice(),
		product.GetID())

	if err != nil {
		return nil, err
	}

	return product, nil

}
