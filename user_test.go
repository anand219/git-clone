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

const (
	USERS_ROUTE = "/v1/api/users"
)

// test stores the HTTP testing client preconfigured
//var client = baloo.New("http://localhost:5000")
func GetUser(t *testing.T, httpStatus int, email string, password string) (response dto.UserGetResponse) {
	route := fmt.Sprintf("%s/whoami", USERS_ROUTE)
	util.AuthorizedAPIClientFor(email, password).
		Get(route).
		Expect(t).
		Status(httpStatus).
		Type(constants.RESPONSE_TYPE_JSON).
		AssertFunc(util.ParseJSON(&response)).
		Done()

	return response
}

func SuspendUser(t *testing.T, userID string) (response dto.UserSuspendResponse) {
	route := fmt.Sprintf("%s/suspend", USERS_ROUTE)
	util.AuthorizedAPIClient().
		Post(route).
		JSON(map[string]string{
			"user_id": userID,
		}).
		Expect(t).
		Status(http.StatusOK).
		Type(constants.RESPONSE_TYPE_JSON).
		AssertFunc(util.ParseJSON(&response)).
		Done()

	return response
}

func TestUsers(t *testing.T) {

	var (
		signupCode        string
		verificationToken string
		userEmailAddress  string
		err               error
	)

	const (
		PASSWORD     = "Password1!"
		NEW_PASSWORD = "Password2!"
		NEW_NAME     = "Test Name"
	)

	randomGenerator := random.New()
	//route := "/v1/api/users"

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
		util.APIClient().
			Post(USERS_ROUTE).
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
			Post(fmt.Sprintf("%s/verify", USERS_ROUTE)).
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

		response = GetUser(t, http.StatusOK, userEmailAddress, PASSWORD)

		if response.Data.Email != userEmailAddress {
			t.Error("Wrong email address")
			return
		}
		userID = response.Data.ID
	})

	t.Run("list users", func(t *testing.T) {

		var response dto.UserListResponse
		util.AuthorizedAPIClient().
			Get(fmt.Sprintf("%s/all", USERS_ROUTE)).
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

	t.Run("update password", func(t *testing.T) {

		var response dto.APIResponse

		util.AuthorizedAPIClientFor(userEmailAddress, PASSWORD).
			Put(fmt.Sprintf("%s/password", USERS_ROUTE)).
			JSON(map[string]string{
				"current_password": PASSWORD,
				"new_password":     NEW_PASSWORD,
			}).
			Expect(t).
			Status(http.StatusOK).
			Type(constants.RESPONSE_TYPE_JSON).
			AssertFunc(util.ParseJSON(&response)).
			Done()

		_, err = util.Authenticate(userEmailAddress, NEW_PASSWORD)
		if err != nil {
			t.Error(err)
			return
		}
	})

	t.Run("reset password", func(t *testing.T) {

		var response dto.UserPasswordResetResponse
		util.AuthorizedAPIClientFor(userEmailAddress, NEW_PASSWORD).
			Post(fmt.Sprintf("%s/password/reset", USERS_ROUTE)).
			JSON(map[string]string{
				"email": userEmailAddress,
			}).
			Expect(t).
			Status(http.StatusOK).
			Type(constants.RESPONSE_TYPE_JSON).
			AssertFunc(util.ParseJSON(&response)).
			Done()

		if response.Data.VerificationToken_ == "" {
			t.Error("No verfication token. Ensure user microservice is run with EXECUTION_MODE=test")
		} else {
			verificationToken = response.Data.VerificationToken_
		}
	})

	t.Run("confirm reset password", func(t *testing.T) {

		var response dto.UserPasswordResetResponse
		util.APIClient().
			Post(fmt.Sprintf("%s/password/reset/confirm", USERS_ROUTE)).
			JSON(map[string]string{
				"token":    verificationToken,
				"password": PASSWORD,
			}).
			Expect(t).
			Status(http.StatusOK).
			Type(constants.RESPONSE_TYPE_JSON).
			AssertFunc(util.ParseJSON(&response)).
			Done()

		if response.Data.VerificationToken_ != "" {
			t.Error("No verfication token. Ensure user microservice is run with EXECUTION_MODE=test")
		}

		_, err = util.Authenticate(userEmailAddress, PASSWORD)
		if err != nil {
			t.Error(err)
			return
		}
	})

	t.Run("update profile", func(t *testing.T) {

		var response dto.UserProfileUpdateResponse
		util.AuthorizedAPIClientFor(userEmailAddress, PASSWORD).
			Put(fmt.Sprintf("%s/profile", USERS_ROUTE)).
			JSON(map[string]string{
				"gender":       "male",
				"name":         NEW_NAME,
				"country_code": "123",
				"phone_number": randomGenerator.PhoneNumber(2),
			}).
			Expect(t).
			Status(http.StatusOK).
			Type(constants.RESPONSE_TYPE_JSON).
			AssertFunc(util.ParseJSON(&response)).
			Done()

		profileResponse := GetUser(t, http.StatusOK, userEmailAddress, PASSWORD)
		if profileResponse.Data.Name != NEW_NAME {
			t.Errorf("Name did not change: '%s'", profileResponse.Data.Name)
		}
	})

	t.Run("suspend user", func(t *testing.T) {
		SuspendUser(t, userID)
		//GetUser(t, http.StatusForbidden, userEmailAddress, PASSWORD)
	})

}
