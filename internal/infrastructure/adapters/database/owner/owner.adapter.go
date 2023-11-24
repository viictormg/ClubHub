package database

import "gorm.io/gorm"

type OwnerAdapter struct {
	db *gorm.DB
}

func NewOwnerAdapter(db *gorm.DB) *OwnerAdapter {
	return &OwnerAdapter{
		db: db,
	}
}
