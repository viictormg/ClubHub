package database

import "gorm.io/gorm"

const (
	tableFranchise = "franchise"
)

type FranchiseAdapter struct {
	db *gorm.DB
}

func NewFranchiseAdapter(db *gorm.DB) *FranchiseAdapter {
	return &FranchiseAdapter{
		db: db,
	}
}
