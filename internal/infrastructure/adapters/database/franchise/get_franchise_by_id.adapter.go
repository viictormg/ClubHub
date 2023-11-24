package database

import (
	"github.com/viictormg/clubHub/internal/domain/entity"
)

func (f *FranchiseAdapter) GetFranchiseByIDAdapter(id int) (*entity.FranchiseEntity, error) {
	var franchise entity.FranchiseEntity

	err := f.db.Table(tableFranchise).
		Where("id = ?", id).
		Find(&franchise).Error

	if err != nil {
		return nil, err
	}

	return &franchise, nil
}
