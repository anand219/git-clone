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
