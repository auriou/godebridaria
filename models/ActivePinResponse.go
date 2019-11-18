package models

type ActivePinResponse struct {
	Success   bool   `json:"success"`
	Token     string `json:"token"`
	Activated bool   `json:"activated"`
	ExpiresIn int    `json:"expires_in"`
}
