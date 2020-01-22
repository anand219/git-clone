package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/consensys/bpaas-e2e/constants"
	"gopkg.in/h2non/baloo.v3"
)

var (
	client       *baloo.Client
	doOnceClient sync.Once

	authorizedClient       *baloo.Client
	doOnceAuthorizedClient sync.Once
)

// APIClient initialises and returns the api client (Singleton)
func APIClient() *baloo.Client {
	doOnceClient.Do(func() {
		client = baloo.New(constants.API_URL)
	})
	return client
}

// AuthorizedAPIClient initializes and returns an authorized api client (Singleton)
func AuthorizedAPIClient() *baloo.Client {
	doOnceAuthorizedClient.Do(func() {
		jwt, err := Authenticate(constants.ADMIN_EMAIL, constants.ADMIN_PASSWORD)
		if err != nil {
			log.Fatalln(err)
		}
		authorizedClient = baloo.New(constants.API_URL).UseRequest(AuthMiddleware(jwt))
	})
	return authorizedClient
}

// AuthorizedAPIClientWith initializes with the given jwt and returns an authorized api client
func AuthorizedAPIClientWith(jwt string) *baloo.Client {
	return baloo.New(constants.API_URL).
		UseRequest(AuthMiddleware(jwt))
}

// AuthorizedAPIClientFor initializes  and returns an authorized api client
func AuthorizedAPIClientFor(username string, password string) *baloo.Client {
	jwt, err := Authenticate(username, password)
	if err != nil {
		log.Fatalln(err)
	}
	return baloo.New(constants.API_URL).
		UseRequest(AuthMiddleware(jwt))
}

// ExtractBodyAsString extracts the response body as string
func ExtractBodyAsString(target *string) func(res *http.Response, req *http.Request) error {
	return func(res *http.Response, req *http.Request) error {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		bodyString := string(bodyBytes)
		*target = bodyString
		return nil
	}
}

// ParseJSON parse the json response body
func ParseJSON(target interface{}) func(res *http.Response, req *http.Request) error {
	return func(res *http.Response, req *http.Request) error {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return json.Unmarshal(bodyBytes, target)
	}
}
