package ports

import "github.com/viictormg/clubHub/internal/domain/entity"

type IOwnerAdapter interface {
	GetOwnerByCompanyIDAdapter(companyID int) (*entity.OwnerEntity, error)
}
