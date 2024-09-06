package main

import (
	"log"
	"os"

	"github.com/TensoRaws/FinalRip/common/version"
	"github.com/TensoRaws/FinalRip/worker/cmd"
)

func main() {
	app := cmd.NewApp()
	app.Name = "FinalRip Worker"
	app.Usage = "Cut, Encode, and Merge videos"
	app.Description = "FinalRip Worker is a tool to cut, encode, and merge videos"
	app.Version = version.FINALRIP_VERSION

	err := app.Run(os.Args)
	if err != nil {
		log.Printf("Failed to run with %s: %v\\n", os.Args, err)
	}
}
