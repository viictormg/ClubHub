package ports

import (
	"github.com/viictormg/clubHub/internal/application/model"
	"github.com/viictormg/clubHub/internal/domain/dto"
)

type IFranchiseUsecase interface {
	CreateFranchiseUsecase(model.FranchiseCreateModel) (*dto.CreationDTO, error)
	GetFranchiseByIDUsecase(id int) (*dto.FranchiseDTO, error)
	GetFranchiseByNameUsecase(name string) (*dto.FranchiseDTO, error)
	GetFranchiseParamUsecase(key, param string) (*[]dto.FranchiseDTO, error)
}
