package end_to_end

import (
	"encoding/json"
	"testing"
)

var (
	tokenCode string
)

type TokensDTO struct {
	Code string
}

const tokenSchema = `{
	"title": "Token",
	"type": "object",
	"properties": {
		"data": {
			"type": "object",
			"properties": {
				"Code": {
					"type": "string"
				}
			}
		}
	},
	"required": ["data"]
}`

type TokenResponse struct {
	Data  TokensDTO
	Error string
}

func TestTokens(t *testing.T) {

	MakeClient()
	t.Run("Create a signup token", func(t *testing.T) {

		Client.Post("/v1/api/tokens").
			JSON(map[string]string{"token_type": "SIGNUP"}).
			Expect(t).
			Status(200).
			Type("json").
			JSONSchema(tokenSchema).
			AssertFunc(GetBody).
			Done()

		bodyData, err := UnmarshalTokenData(BodyString)
		if err != nil {
			t.Error(err)
			return
		}

		tokenCode = bodyData.Data.Code
	})

	t.Run("get token", func(t *testing.T) {

		Client.Get("/v1/api/tokens").
			AddQuery("token_code", tokenCode).
			Expect(t).
			Status(200).
			Type("json").
			JSONSchema(tokenSchema).
			AssertFunc(GetBody).
			Done()

		bodyData := TokenResponse{
			Data: TokensDTO{},
		}

		err := json.Unmarshal([]byte(BodyString), &bodyData)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("sign up with token", func(t *testing.T) {
		emailAddress := MakeEmailAddress()
		Client.Post("/v1/api/users").
			JSON(map[string]string{
				"email":    emailAddress,
				"password": "Password1!",
				"token":    tokenCode,
			}).
			Expect(t).
			Status(200).
			Type("json").
			JSONSchema(GeneralResponseSchema).
			AssertFunc(GetBody).
			Done()
	})
}

func UnmarshalTokenData(s string) (*TokenResponse, error) {
	tokenData := TokenResponse{
		Data: TokensDTO{},
	}

	err := json.Unmarshal([]byte(s), &tokenData)
	return &tokenData, err
}
