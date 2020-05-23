package models

type DataPinResponse struct {
	Status string       `json:"status"`
	Data   *PinResponse `json:"Data"`
}

type PinResponse struct {
	Pin       string `json:"pin"`
	Check     string `json:"check"`
	ExpiredIn int    `json:"expired_in"`
	UserURL   string `json:"user_url"`
	BaseURL   string `json:"base_url"`
	CheckURL  string `json:"check_url"`
}
