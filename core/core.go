package core

import (
	"net/http"
	"strings"

	"github.com/auriou/godebridaria/alldebrid"
	"github.com/auriou/godebridaria/aria2"
	"github.com/auriou/godebridaria/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type ClientCore struct {
	Alldebrid *alldebrid.ClientAlldebrid
	Aria2     *aria2.ClientAria2
	Config    *config.ClientConfig
}

func New() *ClientCore {
	client := &ClientCore{}
	client.Config = config.New()
	client.Alldebrid = alldebrid.New(client.Config)
	client.Alldebrid.CheckActivate()
	client.Aria2 = aria2.New(client.Config)
	return client
}

func (c *ClientCore) Debrid(link string) {
	c.Alldebrid.PrintDebrid(link)
}

func (c *ClientCore) Download(link string) []string {
	urls := c.Alldebrid.Debrid(link)
	c.Aria2.AddUrl(urls)
	return urls
}

func (c *ClientCore) StartApi() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/download", c.HandlerDownload)
	e.GET("/unlock", c.HandlerFormDownload)

	// Start server
	e.Logger.Fatal(e.Start(c.Config.GetAddress())) // ":1234"
}

func (client *ClientCore) HandlerDownload(c echo.Context) error {
	url := c.QueryParam("url")
	token := c.QueryParam("token")
	if token == client.Config.GetToken() {
		urls := client.Download(url)
		return c.HTML(http.StatusOK, strings.Join(urls, ", <br/> "))
	} else {
		return c.HTML(http.StatusUnauthorized, "token error")
	}
}

func (client *ClientCore) HandlerFormDownload(c echo.Context) error {
	form := `<html><body><form action="./download" method="get" > <div> <label for="token">Token &nbsp;:</label> <input type="token" name="token" id="token" required style="width:200px;" > </div><div> <label for="name">Link &nbsp;&nbsp; :</label> <input type="text" name="url" id="url" required style="width:600px;"> </div><div> <input type="submit" value="Download"> </div></form></body></html>`
	return c.HTML(http.StatusOK, form)
}
