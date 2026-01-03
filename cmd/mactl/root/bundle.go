package root

import (
	"github.com/leftbin/mactl/cmd/mactl/root/bundle"
	"github.com/spf13/cobra"
)

var Bundle = &cobra.Command{
	Use:   "bundle",
	Short: "bundle management",
}

func init() {
	Bundle.AddCommand(bundle.List, bundle.Add)
}
