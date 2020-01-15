package random

import (
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

	var response dto.CompanyCreateResponse
	if resp.Ok {
		resp.JSON(&response)
		return &response.Data, nil
	}
	return nil, err
}
