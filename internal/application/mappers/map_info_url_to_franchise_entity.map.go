package mappers

import (
	"github.com/viictormg/clubHub/internal/domain/dto"
	"github.com/viictormg/clubHub/internal/domain/entity"
)

func MapInfoURLToFranchiseEntity(franchise *entity.FranchiseEntity, sslInfo *dto.SSLInfoResultDTO) {
	var redirections []entity.RedirectTo

	for _, redirection := range sslInfo.Endpoints {
		jump := entity.RedirectTo{
			IPAddress:  redirection.IPAddress,
			ServerName: redirection.ServerName,
		}
		redirections = append(redirections, jump)
	}

	franchise.Protocol = sslInfo.Protocol
	franchise.Rerections = redirections
	franchise.NumberRedirections = len(redirections)
}
