package end_to_end

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/consensys/bpaas-e2e/constants"
	"github.com/consensys/bpaas-e2e/dto"
	"github.com/consensys/bpaas-e2e/util"
)

func TestPing(t *testing.T) {
	route := "/v1/api/ping/%s"

	t.Run("with invalid service", func(t *testing.T) {
		util.APIClient().
			Get(fmt.Sprintf(route, "invalid_service")).
			Expect(t).
			Status(http.StatusBadRequest).
			Type(constants.RESPONSE_TYPE_JSON).
			JSON(&dto.APIResponse{Error: "invalid service"}).
			Done()
	})
	t.Run("with valid service", func(t *testing.T) {
		util.APIClient().
			Get(fmt.Sprintf(route, "user")).
			Expect(t).
			Status(http.StatusOK).
			Type(constants.RESPONSE_TYPE_JSON).
			JSON(&dto.APIResponse{Data: "pong"}).
			Done()
	})
}
