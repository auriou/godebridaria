package models

type PinResponse struct {
	Success   bool   `json:"success"`
	Pin       string `json:"pin"`
	ExpiredIn int    `json:"expired_in"`
	UserURL   string `json:"user_url"`
	BaseURL   string `json:"base_url"`
	CheckURL  string `json:"check_url"`
}
