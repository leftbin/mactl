package root

import (
	"github.com/leftbin/mactl/cmd/mactl/root/zshrc"
	"github.com/spf13/cobra"
)

var Zshrc = &cobra.Command{
	Use:   "zshrc",
	Short: "zshrc management",
}

func init() {
	Zshrc.AddCommand(zshrc.Default, zshrc.Override, zshrc.Generate, zshrc.Generated)
}
