package usecase

import "github.com/viictormg/clubHub/internal/domain/dto"

func (f *FranchiseUsecase) GetFranchiseByIDUsecase(id int) (*dto.Franchise, error) {
	return f.FranchiseAdapterDB.GetFranchiseByID(id)
}
