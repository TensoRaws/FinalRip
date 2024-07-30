package cmd

import (
	"github.com/TensoRaws/FinalRip/module/config"
	"github.com/TensoRaws/FinalRip/module/db"
	"github.com/TensoRaws/FinalRip/module/log"
	"github.com/TensoRaws/FinalRip/module/oss"
	"github.com/TensoRaws/FinalRip/module/queue"
	"github.com/TensoRaws/FinalRip/worker/internal/merge"
	"github.com/urfave/cli/v2"
)

var MergeWorker = &cli.Command{
	Name:        "merge",
	Usage:       "Start FinalRip Merge Worker",
	Description: "Start FinalRip Merge Worker",
	Action:      runMergeWorker,
}

func runMergeWorker(ctx *cli.Context) error {
	config.Init()
	log.Init()
	db.Init()
	oss.Init()
	queue.InitMergeWorker()
	merge.Start()
	return nil
}
