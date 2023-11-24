package database

import (
	"github.com/viictormg/clubHub/internal/domain/entity"
)

func (f *FranchiseAdapter) GetFranchisesByParamAdapter(key, value string) (*[]entity.FranchiseEntity, error) {
	var franchises []entity.FranchiseEntity

	err := f.db.Table(tableFranchise).
		Where(key+" ILIKE ?", "%"+value+"%").
		Find(&franchises).Error

	if err != nil {
		return nil, err
	}

	return &franchises, nil
}
