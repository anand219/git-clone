package end_to_end

import (
	"testing"
)

func TestCompanyCreate(t *testing.T) {
	/*
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
	/*
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

		lastCompany := after[len(after)-1]
		if lastCompany.Name != company.Name {
			t.Errorf("Expected company name to be %s got %s", company.Name, lastCompany.Name)
		}*/
}

func TestCompanyListAll(t *testing.T) {
	/*
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

		lastCompany := after[len(after)-1]
		if lastCompany.Name != company.Name {
			t.Errorf("Expected company name to be %s got %s", company.Name, lastCompany.Name)
		}*/
}
