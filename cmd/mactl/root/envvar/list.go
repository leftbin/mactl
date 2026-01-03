package envvar

import (
	"github.com/leftbin/mactl/internal/cli/clierr"
	"github.com/leftbin/mactl/internal/envvar"
	"github.com/spf13/cobra"
)

var List = &cobra.Command{
	Use:   "list",
	Short: "list environment variables",
	Run:   listHandler,
}

func listHandler(cmd *cobra.Command, args []string) {
	clierr.HandleDefault(envvar.List())
}
