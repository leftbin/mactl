package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/leftbin/mactl/internal/cli/flag"
	"github.com/leftbin/mactl/internal/git/config"
)

var Del = &cobra.Command{
	Use:   "del",
	Short: "delete config for host",
	Run:   delHandler,
}

func delHandler(cmd *cobra.Command, args []string) {
	host, err := cmd.InheritedFlags().GetString(string(flag.Host))
	flag.HandleFlagErrAndVal(err, flag.Host, host)
	workspace, err := cmd.InheritedFlags().GetString(string(flag.Workspace))
	flag.HandleFlagErr(err, flag.Host)
	if err := config.Del(host, workspace); err != nil {
		log.Fatal(err)
	}
	log.Info("success")
}
