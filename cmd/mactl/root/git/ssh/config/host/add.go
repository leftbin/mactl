package host

import (
	"github.com/leftbin/mactl/internal/cli/flag"
	sshcon "github.com/leftbin/mactl/internal/git/ssh/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Add = &cobra.Command{
	Use:   "add",
	Short: "add host to git ssh config hosts",
	Run:   addHandler,
}

func init() {
	Add.PersistentFlags().String(string(flag.Host), "", "hostname to add. ex: github.company.com")
}

func addHandler(cmd *cobra.Command, args []string) {
	hostname, err := cmd.PersistentFlags().GetString(string(flag.Host))
	flag.HandleFlagErrAndVal(err, flag.Host, hostname)
	if err := sshcon.AddHost(hostname); err != nil {
		log.Fatalf("failed. err: %v", err)
	}
	log.Info("success")
}
