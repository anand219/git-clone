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

var (
	userID string
)

const (
	route    = "/v1/api/users/platform"
	PASSWORD = "Password1!"
)

func CreatePlatformUser(t *testing.T, emailAddress string) (response dto.PlatformUserCreateResponse) {

	util.AuthorizedAPIClient().
		Post(route).
		JSON(map[string]string{
			"email":            emailAddress,
			"platform_role_id": "1",
		}).
		Expect(t).
		Status(http.StatusOK).
		Type(constants.RESPONSE_TYPE_JSON).
		AssertFunc(util.ParseJSON(&response)).
		Done()

	return response
}

func ActivatePlatformUser(t *testing.T, token string) (response dto.PlatformUserActivateResponse) {

	util.APIClient().
		Post(fmt.Sprintf("%s/activate", route)).
		JSON(map[string]string{
			"token":    token,
			"password": PASSWORD,
		}).
		Expect(t).
		Status(http.StatusOK).
		Type(constants.RESPONSE_TYPE_JSON).
		AssertFunc(util.ParseJSON(&response)).
		Done()

	return response
}

func CancelPlatformUser(t *testing.T, userID string) (response dto.PlatformUserCancelResponse) {
	util.AuthorizedAPIClient().
		Post(fmt.Sprintf("%s/cancel", route)).
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

func TestPlatformUserCreate(t *testing.T) {
	var (
		err               error
		verificationToken string
	)

	const (
		PASSWORD = "Password1!"
	)
	_ = err

	randomGenerator := random.New()
	userEmailAddress := randomGenerator.Email()

	t.Run("Create a platform user", func(t *testing.T) {

		response := CreatePlatformUser(t, userEmailAddress)

		verificationToken = response.Data.VerificationToken_ //In TEST mode, the verification token is returned in the response instead of being sent in an email
		if verificationToken == "" {
			t.Error("No verification token")
		}
	})

	t.Run("Activate a platform user with wrong code", func(t *testing.T) {
		var response dto.PlatformUserActivateResponse
		response = ActivatePlatformUser(t, "WRONG")
		if response.Error == "" {
			t.Error("Allowed activation with wrong code")
		}
	})

	/*t.Run("Sign in a platform user before activation", func(t *testing.T) {
		var response dto.UserGetResponse
		response = GetUser(t, http.StatusUnauthorized, userEmailAddress, PASSWORD)

		if response.Error == "" {
			t.Error("Allowed access before activation")
		}
	})*/

	t.Run("Activate a platform user with correct code", func(t *testing.T) {
		var response dto.PlatformUserActivateResponse
		response = ActivatePlatformUser(t, verificationToken)
		userID = response.Data.ID
		if response.Error != "" {
			t.Error(response.Error)
		}
	})

	t.Run("Sign in a platform user", func(t *testing.T) {
		_, err = util.Authenticate(userEmailAddress, PASSWORD)
		if err != nil {
			t.Error(err)
			return
		}
	})

	t.Run("Cancel a platform user", func(t *testing.T) {
		response := CancelPlatformUser(t, userID)
		if response.Error != "" {
			t.Error(response.Error)
		}
	})

	/*t.Run("Sign in a cancelled platform user", func(t *testing.T) {
		_, err = util.Authenticate(userEmailAddress, PASSWORD)
		if err == nil {
			t.Error("Signed in a cancelled user")
			return
		}
	})*/

}
