package _default

import (
	_default "github.com/leftbin/mactl/internal/zshrc/default"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Create = &cobra.Command{
	Use:     "create",
	Short:   "create default ~/.zshrc file",
	Run:     createHandler,
	Aliases: []string{"cre"},
}

func createHandler(cmd *cobra.Command, args []string) {
	if err := _default.Create(); err != nil {
		log.Fatalf("%v", err)
	}
}
