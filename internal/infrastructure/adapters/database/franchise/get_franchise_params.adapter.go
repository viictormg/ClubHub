package database

import (
	"github.com/viictormg/clubHub/internal/domain/dto"
)

func (f *FranchiseAdapter) GetFranchisesByParamAdapter(key, value string) (*[]dto.Franchise, error) {
	var franchises []dto.Franchise

	err := f.db.Table(tableFranchise).
		Where(key+" ILIKE ?", "%"+value+"%").
		Find(&franchises).Error

	if err != nil {
		return nil, err
	}

	return &franchises, nil
}
