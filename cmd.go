package main

import (
	"fmt"

	"github.com/auriou/godebridaria/config"
	"github.com/auriou/godebridaria/core"
	flags "github.com/jessevdk/go-flags"
)

type Options struct {
	Aria2Url    func(string) `short:"u" long:"aria2-url" description:"url aria2 server"`
	Aria2Secret func(string) `short:"s" long:"aria2-secret" description:"secret aria2 server"`
	Port        func(string) `short:"a" long:"address" description:"address of http service debrid >  :8080"`
	Daemon      func()       `short:"h" long:"http" description:"start http server"`
	Download    func(string) `short:"d" long:"download" description:"start download on aria2 one [url]"`
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
}
