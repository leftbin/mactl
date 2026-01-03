package root

import (
	"github.com/spf13/cobra"
	"github.com/leftbin/mactl/cmd/mactl/root/optimize"
)

var Optimize = &cobra.Command{
	Use:   "optimize",
	Short: "optimize mac features",
}

func init() {
	Optimize.AddCommand(optimize.Dock)
}
