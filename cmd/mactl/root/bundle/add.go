package bundle

import (
	"github.com/leftbin/mactl/internal/bundle"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Add = &cobra.Command{
	Use:   "add",
	Short: "add bundle",
	Run:   addHandler,
}

func addHandler(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Info("bundle name required")
		return
	}
	bundleName := bundle.Get(args[0])
	if bundleName.IsInvalid() {
		log.Errorf("bundle %s not found", args[0])
		return
	}
	if err := bundle.Add(bundleName); err != nil {
		log.Fatalf("%v", err)
	}
	log.Info("success")
}
