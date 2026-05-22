package main

import (
	_ "embed"

	"github.com/luboszk/act/cmd"
	"github.com/luboszk/act/pkg/common"
)

//go:embed VERSION
var version string

func main() {
	ctx, cancel := common.CreateGracefulJobCancellationContext()
	defer cancel()

	// run the command
	cmd.Execute(ctx, version)
}
