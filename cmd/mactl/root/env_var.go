package root

import (
	"github.com/leftbin/mactl/cmd/mactl/root/envvar"
	"github.com/spf13/cobra"
)

var EnvVar = &cobra.Command{
	Use:   "env-var",
	Short: "environment variables management",
}

func init() {
	EnvVar.AddCommand(
		envvar.List,
		envvar.Add,
	)
}
