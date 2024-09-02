package main

import (
	"log"
	"os"
	_ "time/tzdata"

	"github.com/TensoRaws/FinalRip/common/version"
	"github.com/TensoRaws/FinalRip/server/cmd"
)

func main() {
	app := cmd.NewApp()
	app.Name = "FinalRip"
	app.Usage = "FinalRip Aip Sever"
	app.Description = "a distributed video processing tool"
	app.Version = version.FINALRUP_VERSION

	err := app.Run(os.Args)
	if err != nil {
		log.Printf("Failed to run with %s: %v\\n", os.Args, err)
	}
}
