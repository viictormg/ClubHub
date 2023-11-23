package adapters

import (
	"github.com/likexian/whois"
)

func (f *FranchiseAdapterHTTP) ExtractInfoDomainAdapterHTTP(url string) (*string, error) {
	respose, err := whois.Whois(url)

	return &respose, err
}
