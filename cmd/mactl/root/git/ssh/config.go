package ssh

import (
	"github.com/spf13/cobra"
	"github.com/leftbin/mactl/cmd/mactl/root/git/ssh/config"
)

var Config = &cobra.Command{
	Use:   "config",
	Short: "git ssh config management",
}

func init() {
	Config.AddCommand(config.Init, config.Host)
}
