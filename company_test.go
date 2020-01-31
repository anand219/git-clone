package e2e

import (
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/consensys/bpaas-e2e/constants"
	"github.com/consensys/bpaas-e2e/dto"
	"github.com/consensys/bpaas-e2e/random"
	"github.com/consensys/bpaas-e2e/util"
	"gopkg.in/h2non/baloo.v3"
)

var (
	_companyAdmin         *dto.User
	_companyAdminPassword string
	_companyClient        *baloo.Client
	_companyAdminJWT      string
	_err                  error
)

func init() {
	_companyAdmin, _companyAdminPassword, _err = util.CreateUser()
	if _err != nil {
		log.Fatalln(_err)
	}

	_err = util.Verify(_companyAdmin)
	if _err != nil {
		log.Fatalln(_err)

	}

	_companyAdminJWT, _err = util.Authenticate(_companyAdmin.Email, _companyAdminPassword)
	if _err != nil {
		log.Fatalln(_err)
	}

	_companyClient = util.AuthorizedAPIClientWith(_companyAdminJWT)

}
func TestCompanyCreate(t *testing.T) {
	randomGenerator := random.New()

	route := "/v1/api/companies"

	t.Run("without name", func(t *testing.T) {
		_companyClient.Post(route).
			JSON(map[string]string{}).
			Expect(t).
			Status(http.StatusBadRequest).
			Type(constants.RESPONSE_TYPE_JSON).
			JSON(dto.APIResponse{Error: "name is required"}).
			Done()
	})

	t.Run("with name", func(t *testing.T) {
		var response dto.CompanyCreateResponse

		companyName := randomGenerator.Company()

		_companyClient.Post(route).
			JSON(map[string]string{
				"name": companyName,
			}).
			Expect(t).
			Status(http.StatusOK).
			Type(constants.RESPONSE_TYPE_JSON).
			AssertFunc(util.ParseJSON(&response)).
			Done()

		if response.Data.Admin.Email != _companyAdmin.Email {
			t.Errorf("Expected email to be %s got %s", _companyAdmin.Email, response.Data.Admin.Email)
		}

		if response.Data.Name != companyName {
			t.Errorf("Expected company name to be %s got %s", companyName, response.Data.Name)
		}
	})
}

func TestCompanyList(t *testing.T) {
	route := "/v1/api/companies"

	var response dto.CompanyListResponse

	_companyClient.Get(route).
		JSON(map[string]string{}).
		Expect(t).
		Status(http.StatusOK).
		Type(constants.RESPONSE_TYPE_JSON).
		AssertFunc(util.ParseJSON(&response)).
		Done()

	before := response.Data

	company, err := util.CreateCompany(_companyAdminJWT)
	if err != nil {
		t.Error(err)
	}

	_companyClient.Get(route).
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

	util.AuthorizedAPIClient().Get(route).
		JSON(map[string]string{}).
		Expect(t).
		Status(http.StatusOK).
		Type(constants.RESPONSE_TYPE_JSON).
		AssertFunc(util.ParseJSON(&response)).
		Done()

	before := response.Data

	company, err := util.CreateCompany(_companyAdminJWT)
	if err != nil {
		t.Error(err)
	}

	util.AuthorizedAPIClient().Get(route).
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
	t.SkipNow() //TODO: fix
	route := "/v1/api/companies/%s/actions"

	var response dto.CompanyListUserActionsResponse

	company, err := util.CreateCompany(_companyAdminJWT)
	if err != nil {
		t.Error(err)
	}

	_companyClient.Get(fmt.Sprintf(route, company.ID)).
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
	t.SkipNow() //TODO: fix
	route := "/v1/api/companies/users"

	t.Run("without user_id and company_id", func(t *testing.T) {
		_companyClient.Delete(route).
			Expect(t).
			Status(http.StatusBadRequest).
			Type(constants.RESPONSE_TYPE_JSON).
			Done()
	})

	t.Run("with only user_id", func(t *testing.T) {
		_companyClient.Delete(route).
			AddQuery("user_id", "invalid_user_id").
			Expect(t).
			Status(http.StatusBadRequest).
			Type(constants.RESPONSE_TYPE_JSON).
			Done()
	})

	t.Run("with user_id and company_id", func(t *testing.T) {
		_companyClient.Delete(route).
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
		_companyClient.Put(route).
			JSON(map[string]interface{}{}).
			Expect(t).
			Status(http.StatusBadRequest).
			Type(constants.RESPONSE_TYPE_JSON).
			Done()
	})

	t.Run("with only user_id", func(t *testing.T) {
		_companyClient.Put(route).
			JSON(map[string]interface{}{
				"user_id": "invalid_user_id",
			}).
			Expect(t).
			Status(http.StatusBadRequest).
			Type(constants.RESPONSE_TYPE_JSON).
			Done()
	})
	t.Run("with user_id and company_id", func(t *testing.T) {
		_companyClient.Put(route).
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
		_companyClient.Put(route).
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
