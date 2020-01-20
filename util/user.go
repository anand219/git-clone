package util

import (
	"errors"

	"github.com/consensys/bpaas-e2e/constants"
	"github.com/consensys/bpaas-e2e/dto"
	"github.com/consensys/bpaas-e2e/random"
)

var (
	errorInvalidUser = errors.New("Invalid user")
)

// CreateUser creates a new random user in db
func CreateUser() (*dto.User, string, error) {
	randomGenerator := random.New()
	token, err := GenerateToken(constants.TOKEN_TYPE_SIGNUP)
	if err != nil {
		return nil, "", err
	}

	route := "/v1/api/users"

	res, err := APIClient().
		Post(route).
		JSON(map[string]string{
			"email":    randomGenerator.Email(),
			"password": randomGenerator.Password(0),
			"token":    token,
		}).
		Send()

	if err != nil {
		return nil, "", err
	}

	var response dto.UserCreateResponse
	err = res.JSON(&response)
	if err != nil {
		return nil, "", err
	}

	if response.Error != "" {
		return nil, "", errors.New(response.Error)
	}

	return &response.Data, randomGenerator.Password(0), nil
}

// Verify the user
func Verify(user *dto.User) error {
	if user == nil {
		return errorInvalidUser
	}
	resp, err := APIClient().
		Post("/v1/api/users/verify").
		JSON(map[string]string{
			"token": user.VerificationToken_,
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
