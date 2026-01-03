package zshrc

import (
	"github.com/leftbin/mactl/cmd/mactl/root/zshrc/default"
	"github.com/spf13/cobra"
)

var Default = &cobra.Command{
	Use:   "default",
	Short: "manage default zshrc",
}

func init() {
	Default.AddCommand(_default.Create, _default.Show)
}
