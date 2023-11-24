package mappers

import (
	"fmt"
	"strings"

	"github.com/viictormg/clubHub/internal/domain/entity"
)

const (
	containCreationDateMessage               string = "Creation Date:"
	containExpirationDateMessage             string = "Registrar Registration Expiration Date:"
	containRegistrantOrganizationDateMessage string = "Registrant Organization:"
	containEmailContant                      string = "Registrar Abuse Contact Email:"
	containTechOrganization                  string = "Tech Organization: "
)

func MapInfoDomainToFranchiseEntity(franchise *entity.FranchiseEntity, infoDomain *string) error {

	lines := strings.Split(*infoDomain, "\n")

	for _, line := range lines {
		fmt.Println(line)
		if strings.Contains(line, containCreationDateMessage) {
			franchise.DomainCreation = strings.TrimSpace(strings.ReplaceAll(line, containCreationDateMessage, ""))
			break
		}

	}
	for _, line := range lines {
		if strings.Contains(line, containExpirationDateMessage) {
			franchise.DomainExpiration = strings.TrimSpace(strings.ReplaceAll(line, containExpirationDateMessage, ""))
			break
		}

	}
	for _, line := range lines {
		if strings.Contains(line, containRegistrantOrganizationDateMessage) {
			franchise.DomainOwnerName = strings.TrimSpace(strings.ReplaceAll(line, containRegistrantOrganizationDateMessage, ""))
			break
		}
	}

	for _, line := range lines {
		if strings.Contains(line, containEmailContant) {
			franchise.DomainContactEmail = strings.TrimSpace(strings.ReplaceAll(line, containEmailContant, ""))
			break
		}
	}
	for _, line := range lines {
		if strings.Contains(line, containTechOrganization) {
			franchise.Name = strings.TrimSpace(strings.ReplaceAll(line, containTechOrganization, ""))
			break
		}
	}

	return nil
}
