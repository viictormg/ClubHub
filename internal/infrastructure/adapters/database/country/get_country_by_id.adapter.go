package database

import "github.com/viictormg/clubHub/internal/domain/entity"

func (f *CountryAdapterDB) GetCountryByIDAdapter(id int) (*entity.CountryEntity, error) {
	var country entity.CountryEntity

	err := f.db.Table(tableCountry).
		Where("id = ?", id).
		Find(&country).Error

	if err != nil {
		return nil, err
	}

	return &country, nil
}
