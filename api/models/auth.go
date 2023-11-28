package models

type SignInRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInResponse struct {
	AccessToken string `json:"access_token"`
	ExpiredAt   string `json:"expired_at"`
}
