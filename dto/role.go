package dto

// Role struct
type Role struct {
	ID    string
	Name  string
	Title string
}

// CompanyRole struct
type CompanyRole struct {
	Company Company
	Role    Role
}

type RoleListResponse struct {
	Data []*Role `json:"data"`
	APIResponse
}
