package main

import (
	"fmt"
	"os"

	"github.com/auriou/godebridaria/config"
	"github.com/auriou/godebridaria/core"
	flags "github.com/jessevdk/go-flags"
)

type Options struct {
	Aria2Url    func(string) `short:"u" long:"aria2-url" description:"url aria2 server"`
	Aria2Secret func(string) `short:"s" long:"aria2-secret" description:"secret aria2 server"`
	Port        func(string) `short:"a" long:"address" description:"address of http service debrid >  :8080"`
	Daemon      func()       `short:"p" long:"http" description:"start http server"`
	GetPin      func()       `short:"c" long:"getpin" description:"configure apikey for alldebrid"`
	Download    func(string) `short:"d" long:"download" description:"start download on aria2 one [url]"`
	Debrid      func(string) `short:"b" long:"debrid" description:"debrid url [url]"`
}

var options Options

var parser = flags.NewParser(&options, flags.Default)

func init() {
	options.Aria2Url = func(aria2Url string) {
		client := config.New()
		client.SaveAria2Url(aria2Url)
		fmt.Println("Saved")
	}
	options.Aria2Secret = func(aria2Secret string) {
		client := config.New()
		client.SaveAria2Secret(aria2Secret)
		fmt.Println("Saved")
	}
	options.Port = func(address string) {
		client := config.New()
		client.SaveAddress(address)
		fmt.Println("Saved")
	}
	options.Daemon = func() {
		client := core.New()
		client.StartApi()
	}
	options.Download = func(url string) {
		client := core.New()
		client.Download(url)
	}
	options.Debrid = func(url string) {
		client := core.New()
		client.Debrid(url)
	}
	options.GetPin = func() {
		client := core.New()
		client.Activate()
	}
}

func main() {
	/* TESTS
		client := core.New()
		client.Debrid("******")
	*/

	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

}
