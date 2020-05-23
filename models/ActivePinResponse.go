package models

type DataActivePinResponse struct {
	Status string             `json:"status"`
	Data   *ActivePinResponse `json:"data"`
}

type ActivePinResponse struct {
	Apikey    string `json:"apikey"`
	Activated bool   `json:"activated"`
	ExpiresIn int    `json:"expires_in"`
}
