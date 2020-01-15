package dto

// Token struct
type Token struct {
	ID       string
	Code     string
	User     *User
	Type     string
	Used     bool
	DateUsed string
}

// TokenCreateResponse struct
type TokenCreateResponse struct {
	Data Token `json:"data"`
	APIResponse
}
