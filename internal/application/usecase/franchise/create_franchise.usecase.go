package usecase

import (
	"fmt"
	"strings"
	"sync"

	"github.com/viictormg/clubHub/internal/application/mappers"
	"github.com/viictormg/clubHub/internal/application/model"
	"github.com/viictormg/clubHub/internal/domain/dto"
	"github.com/viictormg/clubHub/internal/domain/entity"
)

const (
	numberOfRutines     = 3
	numeroPosibleErrors = 4
)

func (f *FranchiseUsecase) CreateFranchiseUsecase(franchise model.FranchiseCreateModel) (*dto.CreationDTO, error) {
	cleanURL := CleanURL(franchise.URL)

	var franchiseEntity entity.FranchiseEntity
	var wg sync.WaitGroup

	wg.Add(numberOfRutines)
	results := make(chan error, numeroPosibleErrors)

	go func(myFranchise *entity.FranchiseEntity, errCh chan<- error) {
		defer wg.Done()
		sslInfo, err := f.FranchiseAdapterHTTP.ExtractURLInfoAdapterHTTP(cleanURL)
		if err != nil {
			errCh <- err
			return
		}
		errCh <- mappers.MapInfoURLToFranchiseEntity(myFranchise, sslInfo)
	}(&franchiseEntity, results)

	go func(myFranchise *entity.FranchiseEntity, errCh chan<- error) {
		defer wg.Done()
		domainInfo, err := f.FranchiseAdapterHTTP.ExtractInfoDomainAdapterHTTP(cleanURL)
		if err != nil {
			errCh <- err
			return
		}
		mappers.MapInfoDomainToFranchiseEntity(myFranchise, domainInfo)

	}(&franchiseEntity, results)

	go func(myFranchise *entity.FranchiseEntity, errCh chan<- error) {
		defer wg.Done()
		urlImage, err := f.FranchiseAdapterHTTP.ExtractAssetsPageAdapterHTTP(cleanURL)
		if err != nil {
			errCh <- err
			return
		}
		myFranchise.LogoURL = *urlImage
	}(&franchiseEntity, results)

	wg.Wait()
	close(results)

	err := verifyErrorsRutines(results)
	if err != nil {
		return nil, err
	}

	mappers.MapInfoHardCodeToFranchiseEntity(&franchiseEntity)

	return f.FranchiseAdapterDB.CreateFranchiseAdapter(&franchiseEntity)
}

func CleanURL(url string) string {
	cleanURL := url

	replacements := map[string]string{
		"www.":     "",
		"http://":  "",
		"https://": "",
	}

	for oldString, newString := range replacements {
		cleanURL = strings.ReplaceAll(cleanURL, oldString, newString)
	}

	return cleanURL
}

func verifyErrorsRutines(jobs chan error) error {
	for job := range jobs {
		fmt.Println(job)
		if job != nil {
			return job
		}
	}
	return nil
}
