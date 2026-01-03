package zshrc

import (
	"github.com/leftbin/mactl/internal/zshrc/generate"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Generated = &cobra.Command{
	Use:   "generated",
	Short: "open generated zshrc",
	Run:   generatedHandler,
}

func generatedHandler(cmd *cobra.Command, args []string) {
	if err := generate.Show(); err != nil {
		log.Fatalf("%v", err)
	}
}
