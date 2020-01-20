package util

import (
	"errors"

	"github.com/consensys/bpaas-e2e/dto"
	"github.com/consensys/bpaas-e2e/random"
)

// CreateCompany creates a random company with the default admin user or the owner of the jwt
func CreateCompany(jwt string) (*dto.Company, error) {
	client = AuthorizedAPIClient()
	if jwt != "" {
		client = AuthorizedAPIClientWith(jwt)
	}

	randomGenerator := random.New()

	resp, err := client.Post("/v1/api/companies").
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
