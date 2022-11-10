package product

import (
	"context"
	"database/sql"
	"log"

	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/database-fundamentals/products-api/internal/domain"
)

// Diseñar interfaz “Repository” en la que exista un método GetByName() que reciba por parámetro un string y retorne un objeto del tipo Product.
// Implementar el método de forma que con el string recibido lo use para buscar en la DB por el campo “name”.

type Repository interface {
	GetByName(ctx context.Context, name string) (domain.Product, error)
	Store(ctx context.Context, product domain.Product) (domain.Product, error)
	Exists(ctx context.Context, id int) bool
	GetOne(id int) domain.Product
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

const (
	GET_PRODUCT_BY_NAME = "SELECT id, name, type, count, price FROM products WHERE name=?;"
	GET_PRODUCT_BY_ID   = "SELECT id, name, type, count, price FROM products WHERE id=?;"
	SAVE_PRODUCT        = "INSERT INTO products (name, type, count, price) VALUES (?, ?, ?, ?)"
	EXIST_PRODUCT       = "SELECT products.id  FROM products WHERE products.id=?;"
)

func (r *repository) Exists(ctx context.Context, id int) bool {
	rows := r.db.QueryRow(EXIST_PRODUCT, id)
	err := rows.Scan(id) // If no row matches the query, Scan returns an error
	return err == nil
}

func (r *repository) GetByName(ctx context.Context, name string) (domain.Product, error) {
	row := r.db.QueryRow(GET_PRODUCT_BY_NAME, name)
	var product domain.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r *repository) Store(ctx context.Context, product domain.Product) (domain.Product, error) {
	stmt, err := r.db.Prepare(SAVE_PRODUCT) // Create a prepared statement for later queries or executions
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close() // Close the statement preventing memory leaks

	// Execute the prepared statement with the given arguments
	result, err := stmt.Exec(product.Name, product.Type, product.Count, product.Price)
	if err != nil {
		return domain.Product{}, err
	}
	insertedId, err := result.LastInsertId()
	if err != nil {
		return domain.Product{}, err
	}
	product.ID = int(insertedId)

	return product, nil
}

func (r *repository) GetOne(id int) domain.Product {
	var product domain.Product
	rows, err := r.db.Query(GET_PRODUCT_BY_ID, id)
	if err != nil {
		log.Println(err.Error())
		return domain.Product{}
	}
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Println(err.Error())
			return domain.Product{}
		}

	}
	return product
}
