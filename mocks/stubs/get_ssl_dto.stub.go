package stubs

import "github.com/viictormg/clubHub/internal/domain/dto"

func GetSslInfoDTOStub() dto.SSLInfoResultDTO {
	return dto.SSLInfoResultDTO{
		Host: "marriot.com",
		Port: 1,
		Endpoints: []dto.EndpointInfoDTO{
			dto.EndpointInfoDTO{
				IPAddress: "1",
			},
		},
	}
}
