package cmd

import (
	"github.com/TensoRaws/FinalRip/module/config"
	"github.com/TensoRaws/FinalRip/module/db"
	"github.com/TensoRaws/FinalRip/module/log"
	"github.com/TensoRaws/FinalRip/module/oss"
	"github.com/TensoRaws/FinalRip/module/queue"
	"github.com/TensoRaws/FinalRip/worker/internal/cut"
	"github.com/urfave/cli/v2"
)

var CutWorker = &cli.Command{
	Name:        "cut",
	Usage:       "Start FinalRip Cut Worker",
	Description: "Start FinalRip Cut Worker",
	Action:      runCutWorker,
}

func runCutWorker(ctx *cli.Context) error {
	config.Init()
	log.Init()
	db.Init()
	oss.Init()
	queue.InitCutWorker()
	cut.Start()
	return nil
}
