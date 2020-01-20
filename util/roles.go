package util

import (
	"errors"

	"github.com/consensys/bpaas-e2e/dto"
)

var (
	platformRoles = map[string]*dto.Role{}
	companyRoles  = map[string]*dto.Role{}
)

// GetPlatformRoles returns the list of all platform roles
func GetPlatformRoles() (map[string]*dto.Role, error) {
	var err error
	if len(platformRoles) == 0 {
		platformRoles, err = getRolesFrom("/v1/api/roles/platform")
		return platformRoles, err
	}
	return platformRoles, nil

}

// GetCompanyRoles returns the list of all company roles
func GetCompanyRoles() (map[string]*dto.Role, error) {
	var err error
	if len(companyRoles) == 0 {
		companyRoles, err = getRolesFrom("/v1/api/roles/company")
		return companyRoles, err
	}
	return companyRoles, nil
}

func getRolesFrom(route string) (map[string]*dto.Role, error) {
	resp, err := AuthorizedAPIClient().
		Get(route).
		Send()

	if err != nil {
		return nil, err
	}

	var response dto.RoleListResponse
	err = resp.JSON(&response)
	if err != nil {
		return nil, err
	}

	if response.Error != "" {
		return nil, errors.New(response.Error)
	}

	roleMap := make(map[string]*dto.Role)
	for _, role := range response.Data {
		roleMap[role.Name] = role
	}

	return roleMap, nil
}
