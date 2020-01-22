package dto

// User struct
type User struct {
	ID                 string
	Name               string
	Email              string
	PlatformRole       Role
	CompanyRoles       []*CompanyRole
	PhoneNumber        string
	CountryCode        string
	IsVerified         bool
	Status             string
	Gender             string
	VerificationToken_ string
}

// UserListResponse struct
type UserListResponse struct {
	APIResponse
	Data []User `json:"data"`
}

// UserCreateResponse struct
type UserCreateResponse struct {
	APIResponse
	Data User `json:"data"`
}

type UserVerifyResponse struct {
	APIResponse
	Data User `json:"data"`
}

type UserGetResponse struct {
	APIResponse
	Data User `json:"data"`
}
