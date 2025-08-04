package main

import (
	"context"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	TransactionDate time.Time
	Description     string
	Amount          float64
}

func prepare() (db *gorm.DB, ctx context.Context, err error) {
	db, err = gorm.Open(sqlite.Open("accountant.db"), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	ctx = context.Background()
	db.AutoMigrate(&Transaction{})

	return db, ctx, nil
}
