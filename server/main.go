package main

import (
	"log"
	"os"

	"github.com/TensoRaws/FinalRip/server/cmd"
)

const version = "v0.0.1"

func main() {
	app := cmd.NewApp()
	app.Name = "FinalRip"
	app.Usage = "FinalRip Aip Sever"
	app.Description = "a distributed video processing tool"
	app.Version = version

	err := app.Run(os.Args)
	if err != nil {
		log.Printf("Failed to run with %s: %v\\n", os.Args, err)
	}
}
