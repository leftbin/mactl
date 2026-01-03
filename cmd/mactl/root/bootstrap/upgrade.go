package bootstrap

import (
	"github.com/leftbin/mactl/internal/bootstrap"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Upgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "upgrade bootstrap",
	Run:   bootstrapUpgradeHandler,
}

func bootstrapUpgradeHandler(cmd *cobra.Command, args []string) {
	if err := bootstrap.Upgrade(); err != nil {
		log.Fatalf("%v", err)
	}
}
