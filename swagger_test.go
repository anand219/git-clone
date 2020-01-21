package end_to_end

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
			Done()
	})

	t.Run("authorized access", func(t *testing.T) {
		request := util.APIClient().Get(route)
		request.Request.Context.Request.SetBasicAuth(constants.SWAGGER_USERNAME, constants.SWAGGER_PASSWORD)
		request.Expect(t).
			Status(http.StatusOK).
			Type(constants.RESPONSE_TYPE_HTML).
			Done()
	})

}
