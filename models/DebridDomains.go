package models

type DebridDomains struct {
	Hosts []struct {
		AltDomains []string `json:"altDomains"`
		Domain     string   `json:"domain"`
		Status     bool     `json:"status"`
	} `json:"hosts"`
	Redirectors []struct {
		Domain string `json:"domain"`
		Status bool   `json:"status"`
	} `json:"redirectors"`
	Success bool `json:"success"`
}
