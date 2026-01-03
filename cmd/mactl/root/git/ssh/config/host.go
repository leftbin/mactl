package config

import (
	"github.com/leftbin/mactl/cmd/mactl/root/git/ssh/config/host"
	"github.com/spf13/cobra"
)

var Host = &cobra.Command{
	Use:   "host",
	Short: "git ssh host management",
}

func init() {
	Host.AddCommand(host.List, host.Add, host.Del)
}
