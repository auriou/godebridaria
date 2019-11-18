package models

type RedirectorLinks struct {
	Success bool     `json:"success"`
	Folder  bool     `json:"folder"`
	Links   []string `json:"links"`
}
