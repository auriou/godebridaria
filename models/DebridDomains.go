package models

type DebridDomains struct {
	Status string `json:"status"`
	Data   struct {
		Hosts       []string `json:"hosts"`
		Redirectors []string `json:"redirectors"`
	} `json:"data"`
}
