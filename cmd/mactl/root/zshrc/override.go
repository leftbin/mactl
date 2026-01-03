package zshrc

import (
	"github.com/leftbin/mactl/cmd/mactl/root/zshrc/override"
	"github.com/spf13/cobra"
)

var Override = &cobra.Command{
	Use:   "override",
	Short: "manage zshrc override",
}

func init() {
	Override.AddCommand(override.Edit, override.Show)
}
