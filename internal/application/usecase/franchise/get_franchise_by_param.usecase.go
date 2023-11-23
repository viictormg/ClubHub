package usecase

import "github.com/viictormg/clubHub/internal/domain/dto"

func (f *FranchiseUsecase) GetFranchiseParamUsecase(key, value string) (*[]dto.Franchise, error) {
	return f.FranchiseAdapterDB.GetFranchisesByParamAdapter(key, value)
}
