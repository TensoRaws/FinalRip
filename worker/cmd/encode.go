package cmd

import (
	"github.com/TensoRaws/FinalRip/module/config"
	"github.com/TensoRaws/FinalRip/module/db"
	"github.com/TensoRaws/FinalRip/module/log"
	"github.com/TensoRaws/FinalRip/module/oss"
	"github.com/TensoRaws/FinalRip/module/queue"
	"github.com/TensoRaws/FinalRip/worker/internal/encode"
	"github.com/urfave/cli/v2"
)

var CmdEnocdeWorker = &cli.Command{
	Name:        "encode",
	Usage:       "Start FinalRip Enocde Worker",
	Description: "Start FinalRip Enocde Worker",
	Action:      runEncodeWorker,
}

func runEncodeWorker(ctx *cli.Context) error {
	config.Init()
	log.Init()
	db.Init()
	oss.Init()
	queue.Init()
	encode.Start()
	return nil
}
