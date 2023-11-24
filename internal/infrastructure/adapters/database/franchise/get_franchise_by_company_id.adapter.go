package database

import (
	"github.com/viictormg/clubHub/internal/domain/entity"
)

func (f *FranchiseAdapter) GetFranchisesByCompanyID(companyID int) (*[]entity.FranchiseEntity, error) {
	var franchises []entity.FranchiseEntity

	err := f.db.Table(tableFranchise).
		Where("company_id = ?", companyID).
		Find(&franchises).Error

	if err != nil {
		return nil, err
	}

	return &franchises, nil
}
