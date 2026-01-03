package ssh

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/leftbin/mactl/internal/git/ssh"
)

var Init = &cobra.Command{
	Use:   "init",
	Short: "initialize ssh",
	Run:   initHandler,
}

func initHandler(cmd *cobra.Command, args []string) {
	if err := ssh.Init(); err != nil {
		log.Fatalf("failed. err: %v", err)
	}
	log.Info("success")
}
