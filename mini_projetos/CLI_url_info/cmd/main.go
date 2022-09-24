package main

import (
	"cli_url_info/app"
	"log"
	"os"
)

func main() {
	appCli := app.Create()
	if err := appCli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
