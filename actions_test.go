package e2e

import (
	"testing"
)

const (
	ACTIONS_ROUTE = "/v1/api/users/actions"
)

func TestRoles(t *testing.T) {
	/*
		var (
			err error
		)

		//TODO: Make this a util function
		t.Run("Get actions for admin", func(t *testing.T) {

			var response dto.ActionsResponse

			util.AuthorizedAPIClient().
				Get(ACTIONS_ROUTE).
				Expect(t).
				Status(http.StatusOK).
				Type(constants.RESPONSE_TYPE_JSON).
				AssertFunc(util.ParseJSON(&response)).
				Done()

			if err != nil {
				t.Error(err)
				return
			}
		})
	*/
}
