package stubs

import "github.com/viictormg/clubHub/internal/application/model"

func GetFranchiseModelStub() model.FranchiseCreateModel {
	return model.FranchiseCreateModel{
		URL: "www.marriot.com",
	}
}
func GetFranchiseModelBadStub() model.FranchiseCreateModel {
	return model.FranchiseCreateModel{
		URL: "bad_url",
	}
}
