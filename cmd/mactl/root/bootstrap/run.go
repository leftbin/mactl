package bootstrap

import (
	"github.com/leftbin/mactl/internal/bootstrap"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Run = &cobra.Command{
	Use:   "run",
	Short: "run bootstrap",
	Run:   bootstrapHandler,
}

func bootstrapHandler(cmd *cobra.Command, args []string) {
	if err := bootstrap.Run(); err != nil {
		log.Fatalf("%v", err)
	}
}
