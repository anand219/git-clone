package end_to_end

import (
	"net/http"
	"testing"

	constants "github.com/consensys/bpaas-e2e/constants"
	"github.com/consensys/bpaas-e2e/dto"
	"github.com/consensys/bpaas-e2e/random"
	"github.com/consensys/bpaas-e2e/util"
)

func TestPlatformUsers(t *testing.T) {
	var (
		err               error
		verificationToken string
	)

	const (
		PASSWORD = "Password1!"
	)

	route := "/v1/api/users/platform"
	randomGenerator := random.New()
	userEmailAddress := randomGenerator.Email()

	t.Run("Create a platform user", func(t *testing.T) {
		var response dto.PlatformUserCreateResponse

		util.AuthorizedAPIClient().
			Post(route).
			JSON(map[string]string{
				"email":            userEmailAddress,
				"platform_role_id": "1",
			}).
			Expect(t).
			Status(http.StatusOK).
			Type(constants.RESPONSE_TYPE_JSON).
			AssertFunc(util.ParseJSON(&response)).
			Done()

		//userID = response.Data.ID
		verificationToken := response.Data.VerificationToken_ //In TEST mode, the verification token is returned in the response instead of being sent in an email

		if verificationToken == "" {
			t.Error("No verification token")
		}

	})

	t.Run("Activate a platform user", func(t *testing.T) {
		var response dto.PlatformUserActivateResponse

		util.APIClient().
			Post(route).
			JSON(map[string]string{
				"token":    verificationToken,
				"password": PASSWORD,
			}).
			Expect(t).
			Status(http.StatusOK).
			Type(constants.RESPONSE_TYPE_JSON).
			AssertFunc(util.ParseJSON(&response)).
			Done()
	})

	t.Run("Sign in a platform user", func(t *testing.T) {
		_, err = util.Authenticate(userEmailAddress, PASSWORD)
		if err != nil {
			t.Error(err)
			return
		}
	})
}
