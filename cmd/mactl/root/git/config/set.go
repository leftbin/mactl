package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/leftbin/mactl/internal/cli/flag"
	"github.com/leftbin/mactl/internal/git/config"
)

var Set = &cobra.Command{
	Use:   "set",
	Short: "set config for host",
	Run:   setHandler,
}

func init() {
	Set.PersistentFlags().StringP(string(flag.Username), "", "", "git username. ex: swarupdonepudi")
	Set.PersistentFlags().StringP(string(flag.Email), "", "", "git email. ex: your-github-email@gmail.com")
	Set.PersistentFlags().StringP(string(flag.SshKey), "", "", "path for ssh key. ex: ~/.ssh/id_rsa")
	flag.MarkAllFlagsRequired(Set)
}

func setHandler(cmd *cobra.Command, args []string) {
	host, err := cmd.InheritedFlags().GetString(string(flag.Host))
	flag.HandleFlagErrAndVal(err, flag.Host, host)
	workspace, err := cmd.InheritedFlags().GetString(string(flag.Workspace))
	flag.HandleFlagErrAndVal(err, flag.Workspace, workspace)
	username, err := cmd.PersistentFlags().GetString(string(flag.Username))
	flag.HandleFlagErrAndVal(err, flag.Username, username)
	email, err := cmd.PersistentFlags().GetString(string(flag.Email))
	flag.HandleFlagErrAndVal(err, flag.Email, email)
	sshKeyPath, err := cmd.PersistentFlags().GetString(string(flag.SshKey))
	flag.HandleFlagErr(err, flag.SshKey)
	if err := config.Set(host, workspace, username, email, sshKeyPath); err != nil {
		log.Fatal(err)
	}
	log.Infof("success")
}
