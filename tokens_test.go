package end_to_end

import (
	"net/http"
	"testing"

	"github.com/consensys/bpaas-e2e/constants"
	"github.com/consensys/bpaas-e2e/dto"
	"github.com/consensys/bpaas-e2e/util"
)

func TestTokenCreate(t *testing.T) {
	const route = "/v1/api/tokens"

	t.Run("with empty body", func(t *testing.T) {
		util.AuthorizedAPIClient().
			Post(route).
			JSON(map[string]interface{}{}).
			Expect(t).
			Status(http.StatusBadRequest).
			Type(constants.RESPONSE_TYPE_JSON).
			JSON(&dto.APIResponse{Error: "token type is required"}).
			Done()
	})

	t.Run("with token_type", func(t *testing.T) {
		var response dto.TokenCreateResponse
		util.AuthorizedAPIClient().
			Post(route).
			JSON(map[string]interface{}{
				"token_type": constants.TOKEN_TYPE_SIGNUP,
			}).
			Expect(t).
			Status(http.StatusOK).
			Type(constants.RESPONSE_TYPE_JSON).
			AssertFunc(util.ParseJSON(&response)).
			Done()

		if response.Data.Code == "" {
			t.Error("Should generate token")
		}
		if response.Data.User != nil {
			t.Error("Should be nil")
		}
		if response.Data.Type != constants.TOKEN_TYPE_SIGNUP {
			t.Errorf("Expected token type to be %s got %s", constants.TOKEN_TYPE_SIGNUP, response.Data.Type)
		}
	})

	t.Run("with token_type and user_id", func(t *testing.T) {
		user, _, err := util.CreateUser()
		if err != nil {
			t.Error(err)
		}
		var response dto.TokenCreateResponse
		util.AuthorizedAPIClient().
			Post(route).
			JSON(map[string]interface{}{
				"token_type": constants.TOKEN_TYPE_SIGNUP,
				"user_id":    user.ID,
			}).
			Expect(t).
			Type(constants.RESPONSE_TYPE_JSON).
			AssertFunc(util.ParseJSON(&response)).
			Done()

		if response.Data.Code == "" {
			t.Error("Should generate token")
		}
		if response.Data.User == nil {
			t.Error("Should return user")
		}
		if response.Data.Type != constants.TOKEN_TYPE_SIGNUP {
			t.Errorf("Expected token type to be %s got %s", constants.TOKEN_TYPE_SIGNUP, response.Data.Type)
		}
		if response.Data.User.ID != user.ID {
			t.Errorf("Expected user id to be %s got %s", user.ID, response.Data.User.ID)
		}
	})
}

func TestTokenGet(t *testing.T) {
	const route = "/v1/api/tokens"

	t.Run("with empty body", func(t *testing.T) {
		util.AuthorizedAPIClient().
			Get(route).
			JSON(map[string]interface{}{}).
			Expect(t).
			Status(http.StatusBadRequest).
			Type(constants.RESPONSE_TYPE_JSON).
			JSON(&dto.APIResponse{Error: "token_code or token_id is required"}).
			Done()
	})

	t.Run("with token_code", func(t *testing.T) {
		token, err := util.GenerateToken(constants.TOKEN_TYPE_SIGNUP)
		if err != nil {
			t.Error(err)
		}
		util.AuthorizedAPIClient().
			Get(route).
			AddQuery("token_code", token.Code).
			Expect(t).
			Status(http.StatusOK).
			Type(constants.RESPONSE_TYPE_JSON).
			JSON(&dto.APIResponse{Data: token}).
			Done()
	})

	t.Run("with token_id", func(t *testing.T) {
		token, err := util.GenerateToken(constants.TOKEN_TYPE_SIGNUP)
		if err != nil {
			t.Error(err)
		}
		util.AuthorizedAPIClient().
			Get(route).
			AddQuery("token_id", token.ID).
			Expect(t).
			Status(http.StatusOK).
			Type(constants.RESPONSE_TYPE_JSON).
			JSON(&dto.APIResponse{Data: token}).
			Done()
	})
}
