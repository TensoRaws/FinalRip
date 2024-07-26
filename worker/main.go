package main

import (
	"log"
	"os"

	"github.com/TensoRaws/FinalRip/worker/cmd"
)

const version = "v0.0.1"

func main() {
	app := cmd.NewApp()
	app.Name = "FinalRip Enocde Worker"
	app.Usage = "FinalRip Enocde Worker is a worker that encode video files"
	app.Description = "FinalRjson Enocde Worker is a worker that encode video files"
	app.Version = version

	err := app.Run(os.Args)
	if err != nil {
		log.Printf("Failed to run with %s: %v\\n", os.Args, err)
	}
}
