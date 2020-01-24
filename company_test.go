package e2e

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCompanyCreate(t *testing.T) {

		randomGenerator := random.New()

		route := "/v1/api/companies"

		t.Run("without name", func(t *testing.T) {
			util.AuthorizedAPIClient().
				Post(route).
				JSON(map[string]string{}).
				Expect(t).
				Status(http.StatusBadRequest).
				Type(constants.RESPONSE_TYPE_JSON).
				JSON(dto.APIResponse{Error: "name is required"}).
				Done()
		})

		t.Run("with name", func(t *testing.T) {
			var response dto.CompanyCreateResponse

			util.AuthorizedAPIClient().
				Post(route).
				JSON(map[string]string{
					"name": randomGenerator.Company(),
				}).
				Expect(t).
				Status(http.StatusOK).
				Type(constants.RESPONSE_TYPE_JSON).
				AssertFunc(util.ParseJSON(&response)).
				Done()

			if response.Data.Admin.Email != constants.ADMIN_EMAIL {
				t.Errorf("Expected email to be %s got %s", constants.ADMIN_EMAIL, response.Data.Admin.Email)
			}

			if response.Data.Name != randomGenerator.Company() {
				t.Errorf("Expected company name to be %s got %s", randomGenerator.Company(), response.Data.Name)
			}
		})*/
}

func TestCompanyList(t *testing.T) {

		route := "/v1/api/companies"

		var response dto.CompanyListResponse

		util.AuthorizedAPIClient().
			Get(route).
			JSON(map[string]string{}).
			Expect(t).
			Status(http.StatusOK).
			Type(constants.RESPONSE_TYPE_JSON).
			AssertFunc(util.ParseJSON(&response)).
			Done()

		before := response.Data

		company, err := random.NewCompany()
		if err != nil {
			t.Error(err)
		}

		util.AuthorizedAPIClient().
			Get(route).
			JSON(map[string]string{}).
			Expect(t).
			Status(http.StatusOK).
			Type(constants.RESPONSE_TYPE_JSON).
			AssertFunc(util.ParseJSON(&response)).
			Done()

		after := response.Data
		const expectedDifferenceInCount = 1
		if len(after)-len(before) != expectedDifferenceInCount {
			t.Errorf("Expected an extra of %d got %d", expectedDifferenceInCount, len(after)-len(before))
		}

		if response.Data.Name != randomGenerator.Company() {
			t.Errorf("Expected company name to be %s got %s", randomGenerator.Company(), response.Data.Name)
		}
	})
}

func TestCompanyList(t *testing.T) {
	route := "/v1/api/companies"

	var response dto.CompanyListResponse

	util.AuthorizedAPIClient().
		Get(route).
		JSON(map[string]string{}).
		Expect(t).
		Status(http.StatusOK).
		Type(constants.RESPONSE_TYPE_JSON).
		AssertFunc(util.ParseJSON(&response)).
		Done()

	before := response.Data

	company, err := util.CreateCompany("")
	if err != nil {
		t.Error(err)
	}

	util.AuthorizedAPIClient().
		Get(route).
		JSON(map[string]string{}).
		Expect(t).
		Status(http.StatusOK).
		Type(constants.RESPONSE_TYPE_JSON).
		AssertFunc(util.ParseJSON(&response)).
		Done()

	after := response.Data
	const expectedDifferenceInCount = 1
	if len(after)-len(before) != expectedDifferenceInCount {
		t.Errorf("Expected an extra of %d got %d", expectedDifferenceInCount, len(after)-len(before))
	}

	lastCompany := after[len(after)-1]
	if lastCompany.Name != company.Name {
		t.Errorf("Expected company name to be %s got %s", company.Name, lastCompany.Name)
	}
}

func TestCompanyListAll(t *testing.T) {
	route := "/v1/api/companies/all"

	var response dto.CompanyListResponse

	util.AuthorizedAPIClient().
		Get(route).
		JSON(map[string]string{}).
		Expect(t).
		Status(http.StatusOK).
		Type(constants.RESPONSE_TYPE_JSON).
		AssertFunc(util.ParseJSON(&response)).
		Done()

	before := response.Data

	company, err := util.CreateCompany("")
	if err != nil {
		t.Error(err)
	}

	util.AuthorizedAPIClient().
		Get(route).
		JSON(map[string]string{}).
		Expect(t).
		Status(http.StatusOK).
		Type(constants.RESPONSE_TYPE_JSON).
		AssertFunc(util.ParseJSON(&response)).
		Done()

	after := response.Data
	const expectedDifferenceInCount = 1
	if len(after)-len(before) != expectedDifferenceInCount {
		t.Errorf("Expected an extra of %d got %d", expectedDifferenceInCount, len(after)-len(before))
	}

	lastCompany := after[len(after)-1]
	if lastCompany.Name != company.Name {
		t.Errorf("Expected company name to be %s got %s", company.Name, lastCompany.Name)
	}
}
func TestCompanyListUserActions(t *testing.T) {
	route := "/v1/api/companies/%s/actions"

	var response dto.CompanyListUserActionsResponse

	company, err := util.CreateCompany("")
	if err != nil {
		t.Error(err)
	}

	util.AuthorizedAPIClient().
		Get(fmt.Sprintf(route, company.ID)).
		Expect(t).
		Status(http.StatusOK).
		Type(constants.RESPONSE_TYPE_JSON).
		AssertFunc(util.ParseJSON(&response)).
		Done()

	if len(response.Data) == 0 {
		t.Error("Shouldn't be empty")
	}
}

func TestCompanyRemoveUser(t *testing.T) {
	route := "/v1/api/companies/users"

	t.Run("without user_id and company_id", func(t *testing.T) {
		util.AuthorizedAPIClient().
			Delete(route).
			Expect(t).
			Status(http.StatusBadRequest).
			Type(constants.RESPONSE_TYPE_JSON).
			Done()
	})

	t.Run("with only user_id", func(t *testing.T) {
		util.AuthorizedAPIClient().
			Delete(route).
			AddQuery("user_id", "invalid_user_id").
			Expect(t).
			Status(http.StatusBadRequest).
			Type(constants.RESPONSE_TYPE_JSON).
			Done()
	})

	t.Run("with user_id and company_id", func(t *testing.T) {
		util.AuthorizedAPIClient().
			Delete(route).
			AddQuery("user_id", "user_id").
			AddQuery("company_id", "company_id").
			Expect(t).
			Status(http.StatusOK).
			Type(constants.RESPONSE_TYPE_JSON).
			JSON(&dto.APIResponse{Status: "completed"}).
			Done()
	})
}
func TestCompanyUpdateUser(t *testing.T) {
	t.SkipNow() //TODO: fix no response from endpoint
	route := "/v1/api/companies/users"

	t.Run("without user_id and company_id", func(t *testing.T) {
		util.AuthorizedAPIClient().
			Put(route).
			JSON(map[string]interface{}{}).
			Expect(t).
			Status(http.StatusBadRequest).
			Type(constants.RESPONSE_TYPE_JSON).
			Done()
	})

	t.Run("with only user_id", func(t *testing.T) {
		util.AuthorizedAPIClient().
			Put(route).
			JSON(map[string]interface{}{
				"user_id": "invalid_user_id",
			}).
			Expect(t).
			Status(http.StatusBadRequest).
			Type(constants.RESPONSE_TYPE_JSON).
			Done()
	})
	t.Run("with user_id and company_id", func(t *testing.T) {
		util.AuthorizedAPIClient().
			Put(route).
			JSON(map[string]interface{}{
				"user_id":    "invalid_user_id",
				"company_id": "invalid_company_id",
			}).
			Expect(t).
			Status(http.StatusBadRequest).
			Type(constants.RESPONSE_TYPE_JSON).
			Done()
	})

	t.Run("with user_id, company_id and role_id", func(t *testing.T) {
		util.AuthorizedAPIClient().
			Put(route).
			JSON(map[string]interface{}{
				"user_id":    "invalid_user_id",
				"company_id": "invalid_company_id",
				"role_id":    "invalid_role_id",
			}).
			Expect(t).
			Status(http.StatusBadRequest).
			Type(constants.RESPONSE_TYPE_JSON).
			Done()
	})
}
