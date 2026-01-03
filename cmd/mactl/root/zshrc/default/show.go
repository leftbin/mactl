package _default

import (
	"github.com/leftbin/mactl/internal/zshrc/default"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Show = &cobra.Command{
	Use:   "show",
	Short: "show default zshrc",
	Run:   showHandler,
}

func showHandler(cmd *cobra.Command, args []string) {
	if err := _default.Show(); err != nil {
		log.Fatalf("%v", err)
	}
}
