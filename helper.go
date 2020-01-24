package e2e

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"

	baloo "gopkg.in/h2non/baloo.v3"
)

var (
	BodyString string
	doOnce     sync.Once
	Client     *baloo.Client
)

//TODO Import these
type RoleDataTransferType struct {
	ID    string
	Name  string
	Title string
}

type CompanyRoleDataTransferType struct {
	Company CompanyDataTransferType
	Role    RoleDataTransferType
}

type CompanyDataTransferType struct {
	ID                string
	Name              string
	Admin             *UserDataTransferType
	OrganizationCount uint
	UserCount         uint
}

type UserDataTransferType struct {
	ID                 string
	Name               string
	Email              string
	PhoneNumber        string
	CountryCode        string
	IsVerified         bool
	Status             string
	Gender             string
	VerificationToken_ string
}

var UserSchema = fmt.Sprintf(`{
	"title": "Token",
	"type": "object",
	"properties": {
		"data": {
			"type": "object",
			"properties": {
				"Email": {
					"type": "string"
				}
			}
		}
	},
	"required": ["data"],
	"definitions": %s
}`, DefinitionsSchema)

func MakeEmailAddress() string {
	return fmt.Sprintf("%d.%d@tempmail.com", Seed().Intn(10000), Seed().Intn(10000))
}

func Seed() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GetBody(res *http.Response, req *http.Request) error {
	bodyBytes, err := ioutil.ReadAll(res.Body)
	BodyString = string(bodyBytes)
	return err
}

func MakeClient() {
	doOnce.Do(func() {
		Client = baloo.New("http://localhost:5000")
	})
}

func UnmarshalUserData(s string) (*UserResponse, error) {
	userData := UserResponse{
		Data: UserDataTransferType{},
	}

	err := json.Unmarshal([]byte(s), &userData)
	return &userData, err
}

func UnmarshalUserListData(s string) (*UserListResponse, error) {
	users := []*UserDataTransferType{}
	userListData := UserListResponse{
		Data: users,
	}

	err := json.Unmarshal([]byte(s), &userListData)
	return &userListData, err
}

func UnmarshalStringData(s string) (string, error) {
	stringData := StringResponse{}

	err := json.Unmarshal([]byte(s), &stringData)
	return stringData.Data, err
}

type StringResponse struct {
	Data  string
	Error string
}

type UserResponse struct {
	Data  UserDataTransferType
	Error string
}

type UserListResponse struct {
	Data  []*UserDataTransferType
	Error string
}

const DefinitionsSchema = `{
	

}`

const GeneralResponseSchema = `{
	"title": "Response",
	"type": "object",
	"properties": {
		"data": {
			"oneOf": [
				{
					"type": "object"
				},
				{
					"type": "array"
				},
				{
					"type": "string"
				}
			]
			
		}
	},
	"required": ["data"]
}`

func Init() {
	MakeClient()
}
