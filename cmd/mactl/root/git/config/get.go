package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/leftbin/mactl/internal/cli/flag"
	"github.com/leftbin/mactl/internal/git/config"
)

var Get = &cobra.Command{
	Use:   "get",
	Short: "get config for host",
	Run:   getHandler,
}

func getHandler(cmd *cobra.Command, args []string) {
	host, err := cmd.InheritedFlags().GetString(string(flag.Host))
	flag.HandleFlagErrAndVal(err, flag.Host, host)
	workspace, err := cmd.InheritedFlags().GetString(string(flag.Workspace))
	flag.HandleFlagErr(err, flag.Host)
	c, err := config.Get(host, workspace)
	if err != nil {
		log.Fatal(err)
	}
	config.Print(c)
}
