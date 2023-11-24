package database

import "github.com/viictormg/clubHub/internal/domain/entity"

func (f *FranchiseAdapter) GetFranchiseByNameAdapter(name string) (*entity.FranchiseEntity, error) {
	var franchise entity.FranchiseEntity

	err := f.db.Table(tableFranchise).
		Where("name = ?", name).
		Find(&franchise).Error

	if err != nil {
		return nil, err
	}

	return &franchise, nil
}
