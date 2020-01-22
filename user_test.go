package e2e

import (
	"encoding/json"
	"fmt"
	"testing"
)

var (
	signupCode               string
	verificationToken        string
	userID                   string
	userEmailAddress         string
	platformUserEmailAddress string
	bodyData                 string
	jwt                      string
	adminJwt                 string
	err                      error
)

const (
	PASSWORD = "Password1!"
)

// test stores the HTTP testing client preconfigured
//var client = baloo.New("http://localhost:5000")

//TODO Import these
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

type UserListResponse struct {
	Data  []*UserDataTransferType
	Error string
}

func TestUsers(t *testing.T) {
	MakeClient()

	t.Run("Sign in as admin", func(t *testing.T) {
		Client.Post("/v1/api/users/auth").
			JSON(map[string]string{
				"email":    "admin@example.com",
				"password": "adminsecret",
			}).
			Expect(t).
			Status(200).
			Type("json").
			AssertFunc(GetBody).
			Done()

		adminJwt, err = unmarshalStringData(BodyString)
		if err != nil {
			t.Error(err)
			return
		}
	})

	t.Run("Create a platform user", func(t *testing.T) {
		platformUserEmailAddress = MakeEmailAddress()
		Client.Post("/v1/api/users/platform").
			AddHeader("Authorization", fmt.Sprintf("Bearer %s", adminJwt)).
			JSON(map[string]string{"email": platformUserEmailAddress, "platform_role_id": "1"}).
			Expect(t).
			Status(200).
			Type("json").
			JSONSchema(userSchema).
			AssertFunc(GetBody).
			Done()
		fmt.Printf("Response %s\n", BodyString)
		userData, err := unmarshalUserData(BodyString)
		verificationToken = userData.Data.VerificationToken_ //In TEST mode, the verification token is returned in the response instead of being sent in an email

		if err != nil {
			t.Error(err)
			return
		}

	})

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
		userEmailAddress = MakeEmailAddress()
		Client.Post("/v1/api/users").
			JSON(map[string]string{
				"email":    userEmailAddress,
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
				"email":    userEmailAddress,
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

		userData, err := unmarshalUserData(BodyString)
		if err != nil {
			t.Error(err)
			return
		}
		if userData.Data.Email != userEmailAddress {
			t.Error("Wrong email address")
			return
		}

	})

	t.Run("list users", func(t *testing.T) {

		Client.Get("/v1/api/users/all").
			Param("id", fmt.Sprint(userID)).
			AddHeader("Authorization", fmt.Sprintf("Bearer %s", adminJwt)).
			Expect(t).
			Status(200).
			Type("json").
			//JSONSchema(GeneralResponseSchema).
			AssertFunc(GetBody).
			Done()

		fmt.Printf("Body: %s\n", BodyString)

		/*userListData, err := unmarshalUserListData(BodyString)
		if err != nil {
			t.Error(err)
			return
		}
		if len(userListData.Data) == 0 {
			t.Error("Empty array")
			return
		}*/

	})

}

func unmarshalUserData(s string) (*UserResponse, error) {
	userData := UserResponse{
		Data: UserDataTransferType{},
	}

	err := json.Unmarshal([]byte(s), &userData)
	return &userData, err
}

func unmarshalUserListData(s string) (*UserListResponse, error) {
	users := []*UserDataTransferType{}
	userListData := UserListResponse{
		Data: users,
	}

	err := json.Unmarshal([]byte(s), &userListData)
	return &userListData, err
}

func unmarshalStringData(s string) (string, error) {
	stringData := StringResponse{}

	err := json.Unmarshal([]byte(s), &stringData)
	return stringData.Data, err
}

func UnmarshalTokenData(s string) (*TokenResponse, error) {
	tokenData := TokenResponse{
		Data: TokensDTO{},
	}

	err := json.Unmarshal([]byte(s), &tokenData)
	return &tokenData, err
}

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
