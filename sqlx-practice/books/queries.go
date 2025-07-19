package main

import "github.com/jmoiron/sqlx"

func getBooksPriceGreaterThan(db *sqlx.DB, price float32) ([]Book, error) {
	var books []Book
	err := db.Select(&books, "SELECT * FROM gorm.books where price > ?", price)
	return books, err
}
