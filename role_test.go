package end_to_end

import (
	"net/http"
	"testing"

	"github.com/consensys/bpaas-e2e/constants"
	"github.com/consensys/bpaas-e2e/dto"
	"github.com/consensys/bpaas-e2e/util"
)

func TestListCompanyRoles(t *testing.T) {
	route := "/v1/api/roles/company"

	response := dto.RoleListResponse{}
	util.AuthorizedAPIClient().
		Get(route).
		JSON(map[string]string{}).
		Expect(t).
		Status(http.StatusOK).
		Type(constants.RESPONSE_TYPE_JSON).
		AssertFunc(util.ParseJSON(&response)).
		Done()

	if response.Error != "" {
		t.Errorf("Error should be empty got `%s`", response.Error)
	}

	const expected = 3
	if len(response.Data) != expected {
		t.Errorf("Expected %d got %d", expected, len(response.Data))
	}
}
func TestListAssignableCompanyRoles(t *testing.T) {
	route := "/v1/api/roles/company/assignable"

	response := dto.RoleListResponse{}
	util.AuthorizedAPIClient().
		Get(route).
		JSON(map[string]string{}).
		Expect(t).
		Status(http.StatusOK).
		Type(constants.RESPONSE_TYPE_JSON).
		AssertFunc(util.ParseJSON(&response)).
		Done()

	if response.Error != "" {
		t.Errorf("Error should be empty got `%s`", response.Error)
	}

	const expected = 2
	if len(response.Data) != expected {
		t.Errorf("Expected %d got %d", expected, len(response.Data))
	}

	for _, role := range response.Data {
		if role.Name == constants.ROLE_COMPANY_ADMIN {
			t.Errorf("%s should be in assignable company roles", constants.ROLE_COMPANY_ADMIN)
		}
	}
}
func TestListPlatformRoles(t *testing.T) {
	route := "/v1/api/roles/platform"

	response := dto.RoleListResponse{}
	util.AuthorizedAPIClient().
		Get(route).
		JSON(map[string]string{}).
		Expect(t).
		Status(http.StatusOK).
		Type(constants.RESPONSE_TYPE_JSON).
		AssertFunc(util.ParseJSON(&response)).
		Done()

	if response.Error != "" {
		t.Errorf("Error should be empty got `%s`", response.Error)
	}

	const expected = 3
	if len(response.Data) != expected {
		t.Errorf("Expected %d got %d", expected, len(response.Data))
	}
}
func TestListAssignablePlatformRoles(t *testing.T) {
	route := "/v1/api/roles/platform/assignable"

	response := dto.RoleListResponse{}
	util.AuthorizedAPIClient().
		Get(route).
		JSON(map[string]string{}).
		Expect(t).
		Status(http.StatusOK).
		Type(constants.RESPONSE_TYPE_JSON).
		AssertFunc(util.ParseJSON(&response)).
		Done()

	if response.Error != "" {
		t.Errorf("Error should be empty got `%s`", response.Error)
	}
	const expected = 2
	if len(response.Data) != expected {
		t.Errorf("Expected %d got %d", expected, len(response.Data))
	}

	for _, role := range response.Data {
		if role.Name == constants.ROLE_PLATFORM_ADMIN {
			t.Errorf("%s should be in assignable platform roles", constants.ROLE_PLATFORM_ADMIN)
		}
	}
}
