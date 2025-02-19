package cmd

import "coolpilot/pkg/core"

func Execute() {
	options := core.ParseOptions()
	runner, _ := core.NewRunner(options)
	runner.Run(nil)
}
