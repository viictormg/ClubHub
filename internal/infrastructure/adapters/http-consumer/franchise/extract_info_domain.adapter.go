package adapters

import (
	"github.com/likexian/whois"
)

func (f *FranchiseAdapter) ExtractInfoDomainAdapter(url string) (*string, error) {
	respose, err := whois.Whois(url)

	return &respose, err
}
