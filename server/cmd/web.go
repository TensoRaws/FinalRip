package cmd

import (
	"github.com/TensoRaws/FinalRip/module/config"
	"github.com/TensoRaws/FinalRip/module/db"
	"github.com/TensoRaws/FinalRip/module/log"
	"github.com/TensoRaws/FinalRip/module/oss"
	"github.com/TensoRaws/FinalRip/module/queue"
	"github.com/TensoRaws/FinalRip/server/internal/router"
	"github.com/urfave/cli/v2"
)

// CmdWeb api 子命令
var CmdWeb = &cli.Command{
	Name:        "server",
	Usage:       "Start FinalRip Api Server",
	Description: "Start FinalRip Api Server",
	Action:      runWeb,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Value:   "3000",
			Usage:   "Temporary port number to prevent conflict",
		},
	},
}

func runWeb(ctx *cli.Context) error {
	config.Init()
	log.Init()
	db.Init()
	oss.Init()
	queue.Init()
	router.Init()
	return nil
}
