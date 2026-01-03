package bundle

import (
	"github.com/leftbin/mactl/internal/bundle"
	"github.com/spf13/cobra"
)

var List = &cobra.Command{
	Use:   "list",
	Short: "list available bundles",
	Run:   listHandler,
}

func listHandler(cmd *cobra.Command, args []string) {
	bundle.PrintList()
}
