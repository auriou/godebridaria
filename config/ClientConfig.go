package config

import "github.com/auriou/godebridaria/models"

type ClientConfig struct {
	File   string
	Hosts  *models.DebridDomains
	Config *models.StoreConfig
}
