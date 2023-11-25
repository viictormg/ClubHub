package database

import (
	"github.com/viictormg/clubHub/internal/domain/dto"
	"github.com/viictormg/clubHub/internal/domain/entity"
)

func (f *FranchiseAdapter) CreateFranchiseAdapter(newFranchise entity.FranchiseEntity) (*dto.CreationDTO, error) {
	err := f.db.Table(tableFranchise).
		Create(&newFranchise).Error

	if err != nil {
		return nil, err
	}

	return &dto.CreationDTO{ID: newFranchise.ID}, nil
}
