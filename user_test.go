package end_to_end

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/consensys/bpaas-e2e/constants"
	"github.com/consensys/bpaas-e2e/dto"
	"github.com/consensys/bpaas-e2e/random"
	"github.com/consensys/bpaas-e2e/util"
)

// test stores the HTTP testing client preconfigured
//var client = baloo.New("http://localhost:5000")

func TestUsers(t *testing.T) {

	var (
		signupCode        string
		verificationToken string
		userEmailAddress  string
		err               error
	)

	const (
		PASSWORD = "Password1!"
	)

	randomGenerator := random.New()
	route := "/v1/api/users"

	t.Run("Sign in as admin", func(t *testing.T) {
		_, err := util.Authenticate(constants.ADMIN_EMAIL, constants.ADMIN_PASSWORD)
		if err != nil {
			t.Error(err)
			return
		}
	})

	//TODO: Make this a util function
	t.Run("Create a signup token", func(t *testing.T) {

		var response dto.TokenCreateResponse

		util.AuthorizedAPIClient().
			Post("/v1/api/tokens").
			JSON(map[string]string{
				"token_type": "SIGNUP",
			}).
			Expect(t).
			Status(http.StatusOK).
			Type(constants.RESPONSE_TYPE_JSON).
			AssertFunc(util.ParseJSON(&response)).
			Done()

		if err != nil {
			t.Error(err)
			return
		}

		signupCode = response.Data.Code
	})

	t.Run("sign up with token", func(t *testing.T) {

		var response dto.UserCreateResponse
		userEmailAddress = randomGenerator.Email()
		fmt.Printf("Signing up with token %s\n", signupCode)
		util.APIClient().
			Post(route).
			JSON(map[string]string{
				"email":    userEmailAddress,
				"password": PASSWORD,
				"token":    signupCode,
			}).
			Expect(t).
			Status(http.StatusOK).
			Type(constants.RESPONSE_TYPE_JSON).
			AssertFunc(util.ParseJSON(&response)).
			Done()

		//userID = response.Data.ID
		verificationToken = response.Data.VerificationToken_ //In TEST mode, the verification token is returned in the response instead of being sent in an email
	})

	t.Run("verify", func(t *testing.T) {

		var response dto.UserVerifyResponse

		util.APIClient().
			Post(fmt.Sprintf("%s/verify", route)).
			JSON(map[string]string{
				"token": verificationToken,
			}).
			Expect(t).
			Status(http.StatusOK).
			Type(constants.RESPONSE_TYPE_JSON).
			AssertFunc(util.ParseJSON(&response)).
			Done()
	})

	t.Run("sign in", func(t *testing.T) {
		_, err = util.Authenticate(userEmailAddress, PASSWORD)
		if err != nil {
			t.Error(err)
			return
		}
	})

	t.Run("get user", func(t *testing.T) {

		var response dto.UserGetResponse

		fmt.Printf("Getting user %s %s\n", userEmailAddress, PASSWORD)
		util.AuthorizedAPIClientFor(userEmailAddress, PASSWORD).
			Get(fmt.Sprintf("%s/whoami", route)).
			//Param("id", fmt.Sprint(userID)).
			Expect(t).
			Status(http.StatusOK).
			Type(constants.RESPONSE_TYPE_JSON).
			AssertFunc(util.ParseJSON(&response)).
			Done()

		if response.Data.Email != userEmailAddress {
			t.Error("Wrong email address")
			return
		}
	})

	t.Run("list users", func(t *testing.T) {

		var response dto.UserListResponse
		util.AuthorizedAPIClient().
			Get(fmt.Sprintf("%s/all", route)).
			Expect(t).
			Status(http.StatusOK).
			Type(constants.RESPONSE_TYPE_JSON).
			AssertFunc(util.ParseJSON(&response)).
			Done()

		if len(response.Data) == 0 {
			t.Error("Empty array")
			return
		}

	})
}
