package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type FranchiseCreateModel struct {
	URL string `json:"URL"`
}

func (f *FranchiseCreateModel) Validate() error {
	return validation.ValidateStruct(f,
		validation.Field(&f.URL, validation.Required),
		validation.Field(&f.URL, is.URL),
	)
}
