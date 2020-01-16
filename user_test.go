package end_to_end

import (
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

		adminJwt, err = UnmarshalStringData(BodyString)
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
			JSONSchema(UserSchema).
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

		userData, err := UnmarshalUserData(BodyString)
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

		jwt, err = UnmarshalStringData(BodyString)
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

		userData, err := UnmarshalUserData(BodyString)
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
			JSONSchema(GeneralResponseSchema).
			AssertFunc(GetBody).
			Done()

		fmt.Printf("Body: %s\n", BodyString)

		userListData, err := UnmarshalUserListData(BodyString)
		if err != nil {
			t.Error(err)
			return
		}
		if len(userListData.Data) == 0 {
			t.Error("Empty array")
			return
		}

	})

}
