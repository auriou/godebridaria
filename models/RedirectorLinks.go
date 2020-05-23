package models

type RedirectorLinks struct {
	Status string `json:"status"`
	Data   struct {
		Links []string `json:"links"`
	} `json:"data"`
}
