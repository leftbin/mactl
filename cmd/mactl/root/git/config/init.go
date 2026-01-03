package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/leftbin/mactl/internal/cli/flag"
	"github.com/leftbin/mactl/internal/git/config/inititialize"
)

var Init = &cobra.Command{
	Use:   "init",
	Short: "initialize git config for selected host and workspace",
	Run:   initHandler,
}

func init() {
	Init.PersistentFlags().StringP(string(flag.Username), "", "", "git username. ex: leftbin-john")
	Init.PersistentFlags().StringP(string(flag.Email), "", "", "git email. ex: john@leftbin.com")
	flag.MarkAllFlagsRequired(Init)
}

func initHandler(cmd *cobra.Command, args []string) {
	host, err := cmd.InheritedFlags().GetString(string(flag.Host))
	flag.HandleFlagErrAndVal(err, flag.Host, host)
	workspace, err := cmd.InheritedFlags().GetString(string(flag.Workspace))
	flag.HandleFlagErr(err, flag.Workspace)
	username, err := cmd.PersistentFlags().GetString(string(flag.Username))
	flag.HandleFlagErrAndVal(err, flag.Username, username)
	email, err := cmd.PersistentFlags().GetString(string(flag.Email))
	flag.HandleFlagErrAndVal(err, flag.Email, email)
	if err := inititialize.Do(host, workspace, username, email); err != nil {
		log.Fatalf("%v", err)
	}
	log.Info("success")
}
