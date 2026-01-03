package root

import (
	"github.com/spf13/cobra"
	"github.com/leftbin/mactl/cmd/mactl/root/bootstrap"
)

var Bootstrap = &cobra.Command{
	Use:   "bootstrap",
	Short: "bootstrap a new mac",
}

func init() {
	Bootstrap.AddCommand(bootstrap.Checklist, bootstrap.Run, bootstrap.Upgrade)
}
