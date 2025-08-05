package main

import (
	"context"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type Wallet struct {
	gorm.Model
	DebitLimit    float32
	CreditLimit   float32
	CurrentAmount float32
	Name          string
	Description   string
}

type Recurrent struct {
	gorm.Model
	StartDate   time.Time
	EndDate     time.Time
	Amount      float32
	Name        string
	Description string
	PayByDay    uint8
}

type Category struct {
	gorm.Model
	ParentID    uint
	Name        string
	Description string
}

type Transaction struct {
	gorm.Model
	TransactionDate time.Time
	Amount          float32
	WalletID        uint
	RecurrentID     uint
	TransactionID   uint
	CategoryID      uint
	Description     string
	IsOutgoing      bool
	IsTransaction   bool
	IsRecurrent     bool
}

func prepare() (db *gorm.DB, ctx context.Context, err error) {
	db, err = gorm.Open(sqlite.Open("accountant.db"), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	ctx = context.Background()
	db.AutoMigrate(&Wallet{})
	db.AutoMigrate(&Recurrent{})
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Transaction{})

	return db, ctx, nil
}
