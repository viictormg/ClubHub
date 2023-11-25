package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/viictormg/clubHub/internal/domain/dto"
	"github.com/viictormg/clubHub/mocks"
	"github.com/viictormg/clubHub/mocks/stubs"
)

func TestCreateFranchiseUsecase(t *testing.T) {
	t.Run("error calling adaper ExtractURLInfoAdapterHTTP", func(t *testing.T) {
		franchiseModel := stubs.GetFranchiseModelStub()
		franchiseAdapterHTTPMock := new(mocks.IFranchiseAdapterHTTP)
		franchiseAdapterDBMock := new(mocks.IFranchiseAdapterDB)
		countryCodeAdapterDBMock := new(mocks.ICountryCodeAdapterDB)
		sslInfoDTOStub := stubs.GetSslInfoDTOStub()
		domainDTOStub := stubs.GetDomainDTOStub()
		assetsStub := stubs.GetAssetsStub()

		url := CleanURL(franchiseModel.URL)

		franchiseAdapterHTTPMock.On("ExtractURLInfoAdapterHTTP", url).Return(&sslInfoDTOStub, errors.New(""))
		franchiseAdapterHTTPMock.On("ExtractInfoDomainAdapterHTTP", url).Return(&domainDTOStub, nil)
		franchiseAdapterHTTPMock.On("ExtractAssetsPageAdapterHTTP", url).Return(&assetsStub, nil)

		franchiseUsecase := NewFranchiseUsecase(franchiseAdapterHTTPMock, franchiseAdapterDBMock, countryCodeAdapterDBMock)

		response, err := franchiseUsecase.CreateFranchiseUsecase(franchiseModel)

		assert.EqualError(t, err, "")
		assert.Nil(t, response)
	})

	t.Run("error calling adaper ExtractInfoDomainAdapterHTTP", func(t *testing.T) {
		franchiseModel := stubs.GetFranchiseModelStub()
		franchiseAdapterHTTPMock := new(mocks.IFranchiseAdapterHTTP)
		franchiseAdapterDBMock := new(mocks.IFranchiseAdapterDB)
		countryCodeAdapterDBMock := new(mocks.ICountryCodeAdapterDB)
		sslInfoDTOStub := stubs.GetSslInfoDTOStub()
		domainDTOStub := stubs.GetDomainDTOStub()
		assetsStub := stubs.GetAssetsStub()

		url := CleanURL(franchiseModel.URL)

		franchiseAdapterHTTPMock.On("ExtractURLInfoAdapterHTTP", url).Return(&sslInfoDTOStub, nil)
		franchiseAdapterHTTPMock.On("ExtractInfoDomainAdapterHTTP", url).Return(&domainDTOStub, errors.New(""))
		franchiseAdapterHTTPMock.On("ExtractAssetsPageAdapterHTTP", url).Return(&assetsStub, nil)

		franchiseUsecase := NewFranchiseUsecase(franchiseAdapterHTTPMock, franchiseAdapterDBMock, countryCodeAdapterDBMock)

		response, err := franchiseUsecase.CreateFranchiseUsecase(franchiseModel)

		assert.EqualError(t, err, "")
		assert.Nil(t, response)
	})
	t.Run("error calling adaper ExtractAssetsPageAdapterHTTP", func(t *testing.T) {
		franchiseModel := stubs.GetFranchiseModelStub()
		franchiseAdapterHTTPMock := new(mocks.IFranchiseAdapterHTTP)
		franchiseAdapterDBMock := new(mocks.IFranchiseAdapterDB)
		countryCodeAdapterDBMock := new(mocks.ICountryCodeAdapterDB)
		sslInfoDTOStub := stubs.GetSslInfoDTOStub()
		domainDTOStub := stubs.GetDomainDTOStub()
		assetsStub := stubs.GetAssetsStub()

		url := CleanURL(franchiseModel.URL)

		franchiseAdapterHTTPMock.On("ExtractURLInfoAdapterHTTP", url).Return(&sslInfoDTOStub, nil)
		franchiseAdapterHTTPMock.On("ExtractInfoDomainAdapterHTTP", url).Return(&domainDTOStub, nil)
		franchiseAdapterHTTPMock.On("ExtractAssetsPageAdapterHTTP", url).Return(&assetsStub, errors.New(""))

		franchiseUsecase := NewFranchiseUsecase(franchiseAdapterHTTPMock, franchiseAdapterDBMock, countryCodeAdapterDBMock)

		response, err := franchiseUsecase.CreateFranchiseUsecase(franchiseModel)

		assert.EqualError(t, err, "")
		assert.Nil(t, response)
	})
	t.Run("create franchise successful", func(t *testing.T) {
		franchiseModel := stubs.GetFranchiseModelStub()
		franchiseAdapterHTTPMock := new(mocks.IFranchiseAdapterHTTP)
		franchiseAdapterDBMock := new(mocks.IFranchiseAdapterDB)
		countryCodeAdapterDBMock := new(mocks.ICountryCodeAdapterDB)
		sslInfoDTOStub := stubs.GetSslInfoDTOStub()
		domainDTOStub := stubs.GetDomainDTOStub()
		assetsStub := stubs.GetAssetsStub()
		franchiseEntityStub := stubs.GetFranchiseEntity()

		url := CleanURL(franchiseModel.URL)

		franchiseAdapterHTTPMock.On("ExtractURLInfoAdapterHTTP", url).Return(&sslInfoDTOStub, nil)
		franchiseAdapterHTTPMock.On("ExtractInfoDomainAdapterHTTP", url).Return(&domainDTOStub, nil)
		franchiseAdapterHTTPMock.On("ExtractAssetsPageAdapterHTTP", url).Return(&assetsStub, nil)
		franchiseAdapterDBMock.On("CreateFranchiseAdapter", franchiseEntityStub).Return(&dto.CreationDTO{ID: 1}, nil)

		franchiseUsecase := NewFranchiseUsecase(franchiseAdapterHTTPMock, franchiseAdapterDBMock, countryCodeAdapterDBMock)

		response, err := franchiseUsecase.CreateFranchiseUsecase(franchiseModel)

		assert.Equal(t, dto.CreationDTO{ID: 1}, *response)
		assert.Nil(t, err)
	})
}
