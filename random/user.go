package random

import (
	"errors"

	"github.com/consensys/bpaas-e2e/constants"
	"github.com/consensys/bpaas-e2e/dto"
	"github.com/consensys/bpaas-e2e/util"
)

// RUser struct
type RUser struct {
	User     *dto.User
	Password string
}

// NewUser creates a new random company in db
func NewUser() (*RUser, error) {
	randomGenerator := New()
	token, err := util.GenerateToken(constants.TOKEN_TYPE_SIGNUP)
	if err != nil {
		return nil, err
	}

	route := "/v1/api/users"

	resp, err := util.APIClient().
		Post(route).
		JSON(map[string]string{
			"email":    randomGenerator.Email(),
			"password": randomGenerator.Password(0),
			"token":    token,
		}).
		Send()

	if err != nil {
		return nil, err
	}

	var response dto.UserCreateResponse
	err = resp.JSON(&response)
	if err != nil {
		return nil, err
	}

	if response.Error != "" {
		return nil, errors.New(response.Error)
	}
	return &RUser{
		User:     &response.Data,
		Password: randomGenerator.Password(0),
	}, nil

}

// Verify the user
func (r *RUser) Verify() error {
	resp, err := util.APIClient().
		Post("/v1/api/users/verify").
		JSON(map[string]string{
			"token": r.User.VerificationToken_,
		}).
		Send()

	if err != nil {
		return err
	}

	var response dto.UserCreateResponse
	err = resp.JSON(&response)
	if err != nil {
		return err
	}

	if response.Error != "" {
		return errors.New(response.Error)
	}

	return nil
}
