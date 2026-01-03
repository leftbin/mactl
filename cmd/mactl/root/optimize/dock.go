package optimize

import (
	"github.com/leftbin/mactl/internal/optimize/dock"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Dock = &cobra.Command{
	Use:   "dock",
	Short: "optimize dock",
	Run:   dockHandler,
}

func dockHandler(cmd *cobra.Command, args []string) {
	if err := dock.Optimize(); err != nil {
		log.Fatalf("%v", err)
	}
	log.Infof("success")
}
