package dto

// Invite struct
type Invite struct {
	ID         string
	Sender     User
	Receiver   User
	Company    Company
	Status     string
	IsOutgoing bool
	RoleID     string
}

// InviteListResponse struct
type InviteListResponse struct {
	Data []Invite `json:"data"`
	APIResponse
}
