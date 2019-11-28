package core

import (
	"net/http"
	"os"
	"strings"

	"github.com/auriou/godebridaria/alldebrid"
	"github.com/auriou/godebridaria/aria2"
	"github.com/auriou/godebridaria/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var header = "<!doctype html><html lang='en'><head> <meta charset='utf-8'> <meta name='viewport' content='width=device-width, initial-scale=1, shrink-to-fit=no'> <title>Download</title> <link rel='stylesheet' href='https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css' integrity='sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm' crossorigin='anonymous'> <link href='style.css' rel='stylesheet'></head><body class='text-center'> <form class='form-download' action='unlock' role='form' method='GET' class='form-horizontal'> <ul class='list-group'> <h5>"
var footer = " </h5> </ul> <div class='form-group'> <button class='btn btn-lg btn-primary btn-block' type='submit'>Download other files</button> </div></form></body></body></html>"

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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (c *ClientCore) StartApi() {
	f, err := os.OpenFile("debrid.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkErr(err)
	defer f.Close()

	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","ip":"${remote_ip}",` +
			`"method":"${method}","uri":"${uri}","status":${status},"error":"${error}"}` + "\n",
		Output: f}))

	e.Static("/", "assets")
	// Routes
	e.POST("/download", c.HandlerDownloadPost)

	// Start server
	e.Logger.Fatal(e.Start(c.Config.GetAddress())) // ":1234"
}

func (client *ClientCore) HandlerDownloadPost(c echo.Context) error {
	token := c.FormValue("token")
	urlsForm := c.FormValue("urls")
	if token == client.Config.GetToken() {
		result := header
		askurls := strings.Split(strings.Replace(urlsForm, "\r\n", "\n", -1), "\n")
		for _, askurl := range askurls {
			urls := client.Download(askurl)
			for _, url := range urls {
				result += "<li class='list-group-item d-flex justify-content-between align-items-center'>" + url +
					"<span class='badge badge-success badge-pill'>OK</span></li>"
			}
		}
		result += footer

		return c.HTML(http.StatusOK, result)
	} else {
		return c.HTML(http.StatusUnauthorized, "token error")
	}
}
