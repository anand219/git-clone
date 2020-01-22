package e2e

import (
	"net/http"
	"testing"

	"github.com/consensys/bpaas-e2e/constants"
	"github.com/consensys/bpaas-e2e/util"
)

func TestSwagger(t *testing.T) {
	route := "/swagger/index.html"

	t.Run("unauthorized access", func(t *testing.T) {
		util.APIClient().
			Get(route).
			Expect(t).
			Status(http.StatusUnauthorized).
			Type(constants.RESPONSE_TYPE_JSON).
			JSON(map[string]interface{}{"message": "Unauthorized"}).
			Done()
	})

	t.Run("authorized access", func(t *testing.T) {
		util.APIClient().
			Get(route).
			SetHeader("Authorization", util.BasicAuth(constants.SWAGGER_USERNAME, constants.SWAGGER_PASSWORD)).
			Expect(t).
			Status(http.StatusOK).
			Type(constants.RESPONSE_TYPE_HTML).
			Done()
	})

}
