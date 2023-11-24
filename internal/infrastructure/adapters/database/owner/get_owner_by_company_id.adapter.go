package database

import "github.com/viictormg/clubHub/internal/domain/entity"

const (
	tableOwner string = "owner"
)

func (o *OwnerAdapter) GetOwnerByCompanyIDAdapter(companyID int) (*entity.OwnerEntity, error) {
	var owner entity.OwnerEntity
	err := o.db.Table(tableOwner).
		Joins("join company c on owner.id = c.owner_id").
		Where("c.id = ?", companyID).
		Find(&owner).Error

	if err != nil {
		return nil, err
	}

	return &owner, nil
}
