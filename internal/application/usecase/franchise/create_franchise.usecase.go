package usecase

import (
	"github.com/viictormg/clubHub/internal/application/model"
	"github.com/viictormg/clubHub/internal/domain/dto"
)

func (f *FranchiseUsecase) CreateFranchiseUsecase(franchise model.FranchiseCreateModel) (*dto.CreationDTO, error) {
	return &dto.CreationDTO{ID: "ddd"}, nil
}
