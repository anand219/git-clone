package dto

// Company struct
type Company struct {
	ID                string
	Name              string
	Admin             *User
	OrganizationCount uint
	UserCount         uint
}

// CompanyCreateResponse struct
type CompanyCreateResponse struct {
	APIResponse
	Data Company `json:"data"`
}

// CompanyListResponse struct
type CompanyListResponse struct {
	APIResponse
	Data []Company `json:"data"`
}

type CompanyListUserActionsResponse struct {
	APIResponse
	Data []string `json:"data"`
}
