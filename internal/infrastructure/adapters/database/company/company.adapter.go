package database

import "gorm.io/gorm"

type CompanyAdapter struct {
	db *gorm.DB
}

const (
	tableCompany string = "company"
)

func NewCompanyAdapter(db *gorm.DB) *CompanyAdapter {
	return &CompanyAdapter{
		db: db,
	}
}
