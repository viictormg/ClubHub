package database

import (
	"github.com/viictormg/clubHub/internal/domain/dto"
)

const (
	tableFranchise = "franchise"
)

func (f *FranchiseAdapter) GetFranchiseByID(id int) (*dto.Franchise, error) {
	var franchise dto.Franchise

	err := f.db.Table(tableFranchise).
		Where("id = ?", id).
		Find(&franchise).Error

	if err != nil {
		return nil, err
	}

	return &franchise, nil
}
