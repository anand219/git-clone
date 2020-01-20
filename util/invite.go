package util

import (
	"errors"

	"github.com/consensys/bpaas-e2e/dto"
)

// CreateInvite creates an invite for the given user
func CreateInvite(company *dto.Company, user *dto.User, role string, jwt string) (*dto.Invite, error) {
	companyRoles, err := GetCompanyRoles()
	if err != nil {
		return nil, err
	}

	companyOperatorRole, ok := companyRoles[role]
	if !ok {
		return nil, errors.New("Company role not found")
	}

	client := AuthorizedAPIClient()
	if jwt != "" {
		client = AuthorizedAPIClientWith(jwt)
	}
	resp, err := client.
		Post("/v1/api/invitations").
		JSON(map[string]interface{}{
			"role_id":        companyOperatorRole.ID,
			"company_id":     company.ID,
			"receiver_email": user.Email,
		}).
		Send()

	if err != nil {
		return nil, err
	}
	var response dto.APIResponse

	err = resp.JSON(&response)
	if err != nil {
		return nil, err
	}

	if response.Error != "" {
		return nil, errors.New(response.Error)
	}

	// Retreive the invite (since there is no dedicated api for getting the invite details falling back to list inviteListResponse)
	resp, err = client.
		Get("/v1/api/invitations").
		Send()

	if err != nil {
		return nil, err
	}

	var inviteListResponse dto.InviteListResponse

	err = resp.JSON(&inviteListResponse)
	if err != nil {
		return nil, err
	}

	if inviteListResponse.Error != "" {
		return nil, errors.New(inviteListResponse.Error)
	}

	if len(inviteListResponse.Data) > 0 {
		return &inviteListResponse.Data[len(inviteListResponse.Data)-1], nil
	}
	return nil, errors.New("Invite not found")
}
