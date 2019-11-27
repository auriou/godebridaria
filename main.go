package main

import (
	"os"

	flags "github.com/jessevdk/go-flags"
)

func main() {
	//client := core.New()
	//client.StartApi()
	//client.Debrid("https://ed-protect.org/olxkRDTH")
	//client.Debrid("https://uptobox.com/58caxhmejtgy")

	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}
