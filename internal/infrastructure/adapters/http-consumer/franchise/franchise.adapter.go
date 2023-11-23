package adapters

import httpClient "github.com/viictormg/clubHub/internal/infrastructure/adapters/http-consumer"

type FranchiseAdapterHTTP struct {
	httpClient httpClient.IHTTPPortClient
}

func NewFranchiseAdapterHTTP(httpClient httpClient.IHTTPPortClient) *FranchiseAdapterHTTP {
	return &FranchiseAdapterHTTP{
		httpClient: httpClient,
	}
}
