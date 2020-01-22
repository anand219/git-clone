package util

import (
	"github.com/consensys/bpaas-e2e/dto"
)

// GenerateToken of the given type
func GenerateToken(tokenType string) (string, error) {
	resp, err := AuthorizedAPIClient().
		Post("/v1/api/tokens").
		JSON(map[string]string{
			"token_type": tokenType,
		}).
		Send()

	if resp.Ok {
		tokenResponse := dto.TokenCreateResponse{}
		resp.JSON(&tokenResponse)
		return tokenResponse.Data.Code, nil
	}
	return "", err
}
