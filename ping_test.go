package end_to_end

import "testing"

import "github.com/consensys/bpaas-e2e/util"

import "fmt"

import "github.com/consensys/bpaas-e2e/dto"

import "net/http"

func TestPing(t *testing.T) {
	route := "/v1/api/ping/%s"

	t.Run("with invalid service", func(t *testing.T) {
		util.APIClient().
			Get(fmt.Sprintf(route, "invalid_service")).
			Expect(t).
			Status(http.StatusBadRequest).
			JSON(&dto.APIResponse{Error: "invalid service"}).
			Done()
	})
	t.Run("with valid service", func(t *testing.T) {
		util.APIClient().
			Get(fmt.Sprintf(route, "user")).
			Expect(t).
			Status(http.StatusOK).
			JSON(&dto.APIResponse{Data: "pong"}).
			Done()
	})
}
