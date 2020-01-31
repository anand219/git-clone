package dto

// PlatformUserCreateResponse struct
type PlatformUserCreateResponse struct {
	APIResponse
	Data User `json:"data"`
}

type PlatformUserActivateResponse struct {
	APIResponse
	Data User `json:"data"`
}

type PlatformUserCancelResponse struct {
	APIResponse
	UserID string `json:"user_id"`
}
