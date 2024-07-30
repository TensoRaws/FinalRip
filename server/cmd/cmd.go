package cmd

import (
	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	app := cli.NewApp()
	app.EnableBashCompletion = true

	// 子命令集
	subCmdWithConfig := []*cli.Command{
		CmdWeb,
	}

	app.Commands = append(app.Commands, subCmdWithConfig...)
	return app
}
