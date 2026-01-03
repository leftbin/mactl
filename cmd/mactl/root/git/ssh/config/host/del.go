package host

import (
	"github.com/leftbin/mactl/internal/cli/flag"
	sshcon "github.com/leftbin/mactl/internal/git/ssh/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Del = &cobra.Command{
	Use:   "del",
	Short: "del host from git ssh config hosts",
	Run:   delHandler,
}

func init() {
	Del.PersistentFlags().String(string(flag.Host), "", "hostname to delete. ex: github.company.com")
}

func delHandler(cmd *cobra.Command, args []string) {
	hostname, err := cmd.PersistentFlags().GetString(string(flag.Host))
	flag.HandleFlagErrAndVal(err, flag.Host, hostname)
	if err := sshcon.DelHost(hostname); err != nil {
		log.Fatalf("failed. err: %v", err)
	}
	log.Info("success")
}
