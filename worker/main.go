package main

import (
	"log"
	"os"

	"github.com/TensoRaws/FinalRip/common/version"
	"github.com/TensoRaws/FinalRip/worker/cmd"
)

func main() {
	app := cmd.NewApp()
	app.Name = "FinalRip Enocde Worker"
	app.Usage = "FinalRip Enocde Worker is a worker that encode video files"
	app.Description = "FinalRjson Enocde Worker is a worker that encode video files"
	app.Version = version.FINALRUP_VERSION

	err := app.Run(os.Args)
	if err != nil {
		log.Printf("Failed to run with %s: %v\\n", os.Args, err)
	}
}
