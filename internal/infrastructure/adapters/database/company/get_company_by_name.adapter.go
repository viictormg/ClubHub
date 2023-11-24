package database

import "github.com/viictormg/clubHub/internal/domain/entity"

func (c *CompanyAdapter) GetCompanyByNameAdapter(name string) (*entity.CompanyEntity, error) {
	var company entity.CompanyEntity

	err := c.db.Table(tableCompany).
		Where("name = ?", name).
		Find(&company).Error

	if err != nil {
		return nil, err
	}

	return &company, nil
}
