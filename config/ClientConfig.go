package config

import "github.com/auriou/godebridaria/models"

type ClientConfig struct {
	File   string
	Config *models.StoreConfig
}
