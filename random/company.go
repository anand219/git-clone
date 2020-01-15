package random

import (
	"errors"

	"github.com/consensys/bpaas-e2e/dto"
	"github.com/consensys/bpaas-e2e/util"
)

// NewCompany creates a new random company in db
func NewCompany() (*dto.Company, error) {
	randomGenerator := New()

	route := "/v1/api/companies"

	resp, err := util.AuthorizedAPIClient().
		Post(route).
		JSON(map[string]string{
			"name": randomGenerator.Company(),
		}).
		Send()

	if err != nil {
		return nil, err
	}

	var response dto.CompanyCreateResponse
	err = resp.JSON(&response)
	if err != nil {
		return nil, err
	}

	if response.Error != "" {
		return nil, errors.New(response.Error)
	}

	return &response.Data, nil

}
