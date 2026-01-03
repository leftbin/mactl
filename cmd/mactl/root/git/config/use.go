package config

import (
	"github.com/leftbin/mactl/internal/cli/flag"
	"github.com/leftbin/mactl/internal/git/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Use = &cobra.Command{
	Use:   "use",
	Short: "set default config to use for the host",
	Run:   useHandler,
}

func useHandler(cmd *cobra.Command, args []string) {
	host, err := cmd.InheritedFlags().GetString(string(flag.Host))
	flag.HandleFlagErrAndVal(err, flag.Host, host)
	workspace, err := cmd.InheritedFlags().GetString(string(flag.Workspace))
	flag.HandleFlagErrAndVal(err, flag.Workspace, workspace)
	if err := config.Use(host, workspace); err != nil {
		log.Fatal(err)
	}
	log.Infof("success")
}
