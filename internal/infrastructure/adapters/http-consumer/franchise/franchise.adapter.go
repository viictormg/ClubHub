package adapters

import httpClient "github.com/viictormg/clubHub/internal/infrastructure/adapters/http-consumer"

type FranchiseAdapter struct {
	httpClient httpClient.IHTTPPortClient
}

func NewFranchiseAdapter(httpClient httpClient.IHTTPPortClient) *FranchiseAdapter {
	return &FranchiseAdapter{
		httpClient: httpClient,
	}
}
