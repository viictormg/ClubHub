package usecase

import (
	"errors"
	"fmt"
	"strings"

	"github.com/viictormg/clubHub/internal/application/model"
	"github.com/viictormg/clubHub/internal/domain/dto"
)

func (f *FranchiseUsecase) CreateFranchiseUsecase(franchise model.FranchiseCreateModel) (*dto.CreationDTO, error) {
	cleanURL := strings.TrimPrefix(franchise.URL, "www.")

	response, err := f.FranchiseAdapter.ExtractURLInfoAdapter(cleanURL)

	fmt.Println(response, err)

	return &dto.CreationDTO{ID: "ddd"}, errors.New("d")
}
