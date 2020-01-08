package end_to_end

import (
	"encoding/json"
	"fmt"
	"testing"
)

var (
	signupCode        string
	verificationToken string
	userID            string
	emailAddress      string
	bodyData          string
	jwt               string
	err               error
)

const (
	PASSWORD = "Password1!"
)

// test stores the HTTP testing client preconfigured
//var client = baloo.New("http://localhost:5000")

type RoleDataTransferType struct {
	ID    string
	Name  string
	Title string
}

type CompanyRoleDataTransferType struct {
	Company CompanyDataTransferType
	Role    RoleDataTransferType
}

type CompanyDataTransferType struct {
	ID                string
	Name              string
	Admin             *UserDataTransferType
	OrganizationCount uint
	UserCount         uint
}

type UserDataTransferType struct {
	ID                 string
	Name               string
	Email              string
	PhoneNumber        string
	CountryCode        string
	IsVerified         bool
	Status             string
	Gender             string
	VerificationToken_ string
}

const userSchema = `{
	"title": "Token",
	"type": "object",
	"properties": {
		"data": {
			"type": "object",
			"properties": {
				"Email": {
					"type": "string"
				}
			}
		}
	},
	"required": ["data"]
}`

type StringResponse struct {
	Data  string
	Error string
}

type UserResponse struct {
	Data  UserDataTransferType
	Error string
}

func TestUsers(t *testing.T) {
	MakeClient()

	t.Run("Create a signup token", func(t *testing.T) {

		Client.Post("/v1/api/tokens").
			JSON(map[string]string{"token_type": "SIGNUP"}).
			Expect(t).
			Status(200).
			Type("json").
			JSONSchema(userSchema).
			AssertFunc(GetBody).
			Done()

		tokenData, err := UnmarshalTokenData(BodyString)
		if err != nil {
			t.Error(err)
			return
		}

		signupCode = tokenData.Data.Code
	})

	t.Run("sign up with token", func(t *testing.T) {
		emailAddress = MakeEmailAddress()
		Client.Post("/v1/api/users").
			JSON(map[string]string{
				"email":    emailAddress,
				"password": PASSWORD,
				"token":    signupCode,
			}).
			Expect(t).
			Status(200).
			Type("json").
			JSONSchema(GeneralResponseSchema).
			AssertFunc(GetBody).
			Done()

		userData, err := unmarshalUserData(BodyString)
		if err != nil {
			t.Error(err)
			return
		}
		userID = userData.Data.ID
		verificationToken = userData.Data.VerificationToken_ //In TEST mode, the verification token is returned in the response instead of being sent in an email

	})

	t.Run("verify", func(t *testing.T) {
		Client.Post("/v1/api/users/verify").
			JSON(map[string]string{
				"token": verificationToken,
			}).
			Expect(t).
			Status(200).
			Done()

	})

	t.Run("sign in ", func(t *testing.T) {
		Client.Post("/v1/api/users/auth").
			JSON(map[string]string{
				"email":    emailAddress,
				"password": PASSWORD,
			}).
			Expect(t).
			Status(200).
			Type("json").
			AssertFunc(GetBody).
			Done()

		jwt, err = unmarshalStringData(BodyString)
		if err != nil {
			t.Error(err)
			return
		}
	})

	t.Run("get user", func(t *testing.T) {

		Client.Get("/v1/api/users/whoami").
			Param("id", fmt.Sprint(userID)).
			AddHeader("Authorization", fmt.Sprintf("Bearer %s", jwt)).
			Expect(t).
			Status(200).
			Type("json").
			JSONSchema(tokenSchema).
			AssertFunc(GetBody).
			Done()

	})

}

func unmarshalUserData(s string) (*UserResponse, error) {
	userData := UserResponse{
		Data: UserDataTransferType{},
	}

	err := json.Unmarshal([]byte(s), &userData)
	return &userData, err
}

func unmarshalStringData(s string) (string, error) {
	stringData := StringResponse{}

	err := json.Unmarshal([]byte(s), &stringData)
	return stringData.Data, err
}