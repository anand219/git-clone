package e2e

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/consensys/bpaas-e2e/constants"
	"github.com/consensys/bpaas-e2e/dto"
	"github.com/consensys/bpaas-e2e/util"
)

func TestCreateInvite(t *testing.T) {
	route := "/v1/api/invitations"
	t.Run("with empty body", func(t *testing.T) {
		util.AuthorizedAPIClient().
			Post(route).
			JSON(map[string]interface{}{}).
			Expect(t).
			Status(http.StatusBadRequest).
			Type(constants.RESPONSE_TYPE_JSON).
			JSON(&dto.APIResponse{Error: "role_id is required"}).
			Done()
	})

	t.Run("with only role_id", func(t *testing.T) {
		util.AuthorizedAPIClient().
			Post(route).
			JSON(map[string]interface{}{
				"role_id": "role_id",
			}).
			Expect(t).
			Status(http.StatusBadRequest).
			Type(constants.RESPONSE_TYPE_JSON).
			JSON(&dto.APIResponse{Error: "company_id is required"}).
			Done()
	})

	t.Run("with role_id and company_id", func(t *testing.T) {
		util.AuthorizedAPIClient().
			Post(route).
			JSON(map[string]interface{}{
				"role_id":    "role_id",
				"company_id": "company_id",
			}).
			Expect(t).
			Status(http.StatusBadRequest).
			Type(constants.RESPONSE_TYPE_JSON).
			JSON(&dto.APIResponse{Error: "invalid email"}).
			Done()
	})

	t.Run("with an invalid email", func(t *testing.T) {
		util.AuthorizedAPIClient().
			Post(route).
			JSON(map[string]interface{}{
				"role_id":        "role_id",
				"company_id":     "company_id",
				"receiver_email": "invalid_email",
			}).
			Expect(t).
			Status(http.StatusBadRequest).
			Type(constants.RESPONSE_TYPE_JSON).
			JSON(&dto.APIResponse{Error: "invalid email"}).
			Done()
	})

	t.Run("with valid role_id, company_id and receiver_email", func(t *testing.T) {
		companyRoles, err := util.GetCompanyRoles()
		if err != nil {
			t.Error(err)
		}
		companyOperatorRole, ok := companyRoles[constants.ROLE_COMPANY_OPERATOR]
		if !ok {
			t.Error("Company role not found")
		}

		company, err := util.CreateCompany("")
		if err != nil {
			t.Error(err)
		}

		user, _, err := util.CreateUser()
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("Here is the companyOperatorRole %+v\n", companyOperatorRole)
		fmt.Printf("Here is the company %+v\n", company)
		fmt.Printf("Here is the usere %+v\n", user)
		json := map[string]interface{}{
			"role_id":        companyOperatorRole.ID,
			"company_id":     company.ID,
			"receiver_email": user.Email,
		}

		fmt.Printf("%+v\n", json)

		util.AuthorizedAPIClient().
			Post(route).
			JSON(json).
			Expect(t).
			Status(http.StatusOK).
			Type(constants.RESPONSE_TYPE_JSON).
			JSON(&dto.APIResponse{Status: "inprogress"}).
			Done()
	})
}

func TestListInvite(t *testing.T) {
	const route = "/v1/api/invitations"

	var response dto.InviteListResponse

	util.AuthorizedAPIClient().
		Get(route).
		Expect(t).
		Status(http.StatusOK).
		Type(constants.RESPONSE_TYPE_JSON).
		AssertFunc(util.ParseJSON(&response)).
		Done()
	before := response.Data

	user, _, err := util.CreateUser()
	if err != nil {
		t.Error(err)
	}

	company, err := util.CreateCompany("")
	if err != nil {
		t.Error(err)
	}

	invite, err := util.CreateInvite(company, user, constants.ROLE_COMPANY_DEVELOPER, "")
	if err != nil {
		t.Error(err)
	}

	util.AuthorizedAPIClient().
		Get(route).
		Expect(t).
		Status(http.StatusOK).
		Type(constants.RESPONSE_TYPE_JSON).
		AssertFunc(util.ParseJSON(&response)).
		Done()

	after := response.Data

	const expectedDifferenceInCount = 1
	if len(after)-len(before) > expectedDifferenceInCount {
		t.Errorf("Expected an extra of %d inviteListResponse got %d ", expectedDifferenceInCount, len(after)-len(before))
	}

	lastInvite := after[len(after)-1]

	if lastInvite.ID != invite.ID {
		t.Errorf("Expected invite id to be %s got %s", invite.ID, lastInvite.ID)
	}
	if lastInvite.RoleID != invite.RoleID {
		t.Errorf("Expected role id to be %s got %s", invite.ID, lastInvite.ID)
	}
}
func TestAcceptInvite(t *testing.T) {
	const route = "/v1/api/invitations/accept"
	t.Run("without invite_id", func(t *testing.T) {
		util.AuthorizedAPIClient().
			Post(route).
			JSON(map[string]interface{}{}).
			Expect(t).
			Status(http.StatusBadRequest).
			Type(constants.RESPONSE_TYPE_JSON).
			JSON(&dto.APIResponse{Error: "invite_id is required"}).
			Done()
	})
	t.Run("with invite_id", func(t *testing.T) {
		user, _, err := util.CreateUser()
		if err != nil {
			t.Error(err)
		}

		company, err := util.CreateCompany("")
		if err != nil {
			t.Error(err)
		}

		invite, err := util.CreateInvite(company, user, constants.ROLE_COMPANY_DEVELOPER, "")
		if err != nil {
			t.Error(err)
		}

		util.AuthorizedAPIClient().
			Post(route).
			JSON(map[string]interface{}{
				"invite_id": invite.ID,
			}).
			Expect(t).
			Status(http.StatusOK).
			Type(constants.RESPONSE_TYPE_JSON).
			JSON(&dto.APIResponse{Status: "inprogress"}).
			Done()
	})
}
func TestRejectInvite(t *testing.T) {
	const route = "/v1/api/invitations/reject"
	t.Run("without invite_id", func(t *testing.T) {
		util.AuthorizedAPIClient().
			Post(route).
			JSON(map[string]interface{}{}).
			Expect(t).
			Status(http.StatusBadRequest).
			Type(constants.RESPONSE_TYPE_JSON).
			JSON(&dto.APIResponse{Error: "invite_id is required"}).
			Done()
	})
	t.Run("with invite_id", func(t *testing.T) {
		user, _, err := util.CreateUser()
		if err != nil {
			t.Error(err)
		}

		company, err := util.CreateCompany("")
		if err != nil {
			t.Error(err)
		}

		invite, err := util.CreateInvite(company, user, constants.ROLE_COMPANY_DEVELOPER, "")
		if err != nil {
			t.Error(err)
		}

		util.AuthorizedAPIClient().
			Post(route).
			JSON(map[string]interface{}{
				"invite_id": invite.ID,
			}).
			Expect(t).
			Status(http.StatusOK).
			JSON(&dto.APIResponse{Status: "inprogress"}).
			Type(constants.RESPONSE_TYPE_JSON).
			Done()
	})
}
