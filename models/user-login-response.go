package models

// UserLoginResponse build the User login response model
type UserLoginResponse struct {
	Token string `json:"token,omitempty"`
}
