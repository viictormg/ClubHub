package adapters

import (
	"encoding/json"
	"net/http"

	"github.com/viictormg/clubHub/internal/domain/dto"
)

const (
	URL_GET_INFO_SITE string = "https://api.ssllabs.com/api/v3/analyze/"
)

func (f *FranchiseAdapterHTTP) ExtractURLInfoAdapterHTTP(domain string) (*dto.SSLInfoResultDTO, error) {
	var infoResultDTO dto.SSLInfoResultDTO

	request, _ := http.NewRequest(http.MethodGet, URL_GET_INFO_SITE, http.NoBody)

	q := request.URL.Query()
	q.Add("host", domain)
	request.URL.RawQuery = q.Encode()

	response, err := f.httpClient.Do(request)

	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 400 {
		return nil, err
	}

	err = json.NewDecoder(response.Body).Decode(&infoResultDTO)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return &infoResultDTO, nil
}
