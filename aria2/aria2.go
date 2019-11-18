package aria2

import (
	"context"
	"time"

	"github.com/auriou/godebridaria/config"
	"github.com/zyxar/argo/rpc"
)

type ClientAria2 struct {
	Url           string
	Secret        string
	ContextConfig *config.ClientConfig
}

func New(cfg *config.ClientConfig) *ClientAria2 {
	client := &ClientAria2{Url: cfg.GetAria2Url(), Secret: cfg.GetAria2Secret()}
	client.ContextConfig = cfg
	return client
}

func (c *ClientAria2) AddUrl(urls []string) error {
	rpcc, err := rpc.New(context.Background(), c.Url, c.Secret, time.Second*5, nil)
	if err != nil {
		return err
	}
	defer rpcc.Close()

	for _, url := range urls {
		rpcc.AddURI(url)
	}
	return nil
}
