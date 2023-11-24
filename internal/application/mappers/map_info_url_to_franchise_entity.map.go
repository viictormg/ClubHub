package mappers

import (
	"encoding/json"

	"github.com/viictormg/clubHub/internal/domain/dto"
	"github.com/viictormg/clubHub/internal/domain/entity"
)

func MapInfoURLToFranchiseEntity(franchise *entity.FranchiseEntity, sslInfo *dto.SSLInfoResultDTO) error {
	var redirections []entity.RedirectTo

	for _, redirection := range sslInfo.Endpoints {
		jump := entity.RedirectTo{
			IPAddress:  redirection.IPAddress,
			ServerName: redirection.ServerName,
		}
		redirections = append(redirections, jump)
	}
	jsonData, err := json.Marshal(redirections)
	if err != nil {
		return err
	}

	franchise.Protocol = sslInfo.Protocol
	franchise.Redirections = string(jsonData)
	franchise.NumberRedirections = len(redirections)
	franchise.URL = sslInfo.Host

	return nil
}
