package util

import (
	"fmt"

	"github.com/consensys/bpaas-e2e/dto"
	"gopkg.in/h2non/gentleman.v2/context"
)

// Authenticate authenticates a user with username and password returns a JWT
func Authenticate(email string, password string) (string, error) {
	resp, err := APIClient().
		Post("/v1/api/users/auth").
		JSON(map[string]string{
			"email":    email,
			"password": password,
		}).
		Send()

	if resp.Ok {
		tokenResponse := dto.APIResponse{}
		resp.JSON(&tokenResponse)
		if jwt, ok := tokenResponse.Data.(string); ok {
			return jwt, nil
		}
	}

	return "", err
}

// AuthMiddleware attaches the Authorization header for the request
func AuthMiddleware(jwt string) func(ctx *context.Context, h context.Handler) {
	return func(ctx *context.Context, h context.Handler) {
		ctx.Request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwt))
		h.Next(ctx)
	}
}
