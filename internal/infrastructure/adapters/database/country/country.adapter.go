package database

import "gorm.io/gorm"

type CountryAdapterDB struct {
	db *gorm.DB
}

const (
	tableCountry string = "country"
)

func NewCountryAdapterDB(db *gorm.DB) *CountryAdapterDB {
	return &CountryAdapterDB{
		db: db,
	}
}
