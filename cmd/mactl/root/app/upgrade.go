package app

import (
	"github.com/leftbin/mactl/internal/app"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Upgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "upgrade",
	Run:   upgradeHandler,
}

func upgradeHandler(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Info("app name required")
		return
	}
	appName := app.Get(args[0])
	if appName.IsInvalid() {
		log.Errorf("app %s not found", args[0])
		return
	}
	if err := app.Add(appName); err != nil {
		log.Fatalf("%v", err)
	}
	log.Info("success")
}
