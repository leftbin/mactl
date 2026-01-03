package app

import (
	"github.com/spf13/cobra"
	"github.com/leftbin/mactl/internal/app"
)

var List = &cobra.Command{
	Use:   "list",
	Short: "list apps",
	Run:   listHandler,
}

func listHandler(cmd *cobra.Command, args []string) {
	app.PrintList()
}
