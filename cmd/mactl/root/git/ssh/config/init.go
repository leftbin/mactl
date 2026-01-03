package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	sshcon "github.com/leftbin/mactl/internal/git/ssh/config"
)

var Init = &cobra.Command{
	Use:   "init",
	Short: "initialize ssh config",
	Run:   initHandler,
}

func initHandler(cmd *cobra.Command, args []string) {
	if err := sshcon.Init(); err != nil {
		log.Fatalf("failed. err: %v", err)
	}
	log.Info("success")
}
