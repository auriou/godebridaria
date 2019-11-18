package alldebrid

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/auriou/godebridaria/config"
	"github.com/auriou/godebridaria/models"
)

func New(cfg *config.ClientConfig) *ClientAlldebrid {
	client := &ClientAlldebrid{
		Base:          "https://api.alldebrid.com",
		Agent:         "JDownloader",
		HTTPClient:    http.DefaultClient,
		ContextConfig: cfg,
	}
	return client
}

func (c *ClientAlldebrid) GetReq(path string, queries Queries) *http.Request {
	url := path
	if !strings.HasPrefix(url, "http") {
		url = c.Base + path
	}
	req, _ := http.NewRequest("GET", url, nil)
	if queries != nil {
		q := req.URL.Query()
		q.Set("agent", c.Agent)
		for k, v := range queries {
			q.Add(k, v)
		}
		if c.ContextConfig.IsActivated() {
			q.Set("token", c.ContextConfig.GetToken())
		}
		req.URL.RawQuery = q.Encode()
	}
	return req
}

func (c *ClientAlldebrid) GetActivate() (*models.PinResponse, error) {
	res, err := c.HTTPClient.Do(c.GetReq("/pin/get", Queries{}))
	pin := &models.PinResponse{}
	err = json.NewDecoder(res.Body).Decode(pin)
	if err != nil {
		return nil, err
	}
	if !pin.Success {
		return nil, errors.New("error : get pin request")
	}
	c.ContextConfig.SavePin(pin)
	return pin, nil
}

func (c *ClientAlldebrid) SetActivate() (*models.ActivePinResponse, error) {
	res, err := c.HTTPClient.Do(c.GetReq(c.ContextConfig.GetCheckURL(), nil))
	pin := &models.ActivePinResponse{}
	err = json.NewDecoder(res.Body).Decode(pin)
	if err != nil {
		return nil, err
	}
	if !pin.Success {
		return nil, errors.New("error : activate pin request")
	}
	c.ContextConfig.SaveActivePin(pin)
	return pin, nil
}

func (c *ClientAlldebrid) Activate() error {
	res, err := c.GetActivate()
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	fmt.Println("Activate pin into your browser", res.UserURL)
	for i := 0; i < 23; i++ {
		time.Sleep(5 * time.Second)
		active, err := c.SetActivate()
		if err == nil && active.Activated {
			fmt.Println("Success Activate : token saved")
			return nil
		}
	}
	return errors.New("error : activate abandoned")
}

func (c *ClientAlldebrid) CheckActivate() error {
	if !c.ContextConfig.IsActivated() {
		return c.Activate()
	}
	return nil
}

func (c *ClientAlldebrid) GetUser() (*models.User, error) {
	res, err := c.HTTPClient.Do(c.GetReq("/user/login", Queries{}))
	user := &models.User{}
	err = json.NewDecoder(res.Body).Decode(user)
	if err != nil {
		return nil, err
	}
	if !user.Success {
		return nil, errors.New("error : user get")
	}
	return user, nil
}

func (c *ClientAlldebrid) PrintUser() {
	user, err := c.GetUser()
	if err != nil {
		fmt.Println(err)
	}
	json, _ := json.MarshalIndent(user, "", "\t")
	fmt.Println("User ", string(json))
}

func (c *ClientAlldebrid) Redirector(link string) (*models.RedirectorLinks, error) {
	res, err := c.HTTPClient.Do(c.GetReq("/link/redirector", Queries{"link": link}))
	redirector := &models.RedirectorLinks{}
	err = json.NewDecoder(res.Body).Decode(redirector)
	if err != nil {
		return nil, err
	}
	if !redirector.Success {
		return nil, errors.New("error : user get")
	}
	return redirector, nil
}

func (c *ClientAlldebrid) Unlock(link string) (*models.DownloadLink, error) {
	res, err := c.HTTPClient.Do(c.GetReq("/link/unlock", Queries{"link": link}))
	unlock := &models.DownloadLink{}
	err = json.NewDecoder(res.Body).Decode(unlock)
	if err != nil {
		return nil, err
	}
	if !unlock.Success {
		return nil, errors.New("error : user get")
	}
	return unlock, nil
}

func (c *ClientAlldebrid) PrintDebrid(link string) {
	redirect, _ := c.Redirector(link)
	for _, value := range redirect.Links {
		link, _ := c.Unlock(value)
		fmt.Println("URL :", value)
		fmt.Println("   + UNLOCK : ", link.Infos.Link)
	}
}

func (c *ClientAlldebrid) Debrid(link string) []string {
	tab := make([]string, 0)
	redirect, _ := c.Redirector(link)
	for _, value := range redirect.Links {
		link, _ := c.Unlock(value)
		tab = append(tab, link.Infos.Link)
	}
	return tab
}
