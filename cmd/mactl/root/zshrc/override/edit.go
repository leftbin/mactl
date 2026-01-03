package override

import (
	"fmt"
	"github.com/leftbin/mactl/internal/zshrc/override"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Edit = &cobra.Command{
	Use:   "edit",
	Short: fmt.Sprintf("edit ~/%s file", override.FileName),
	Run:   editHandler,
}

func editHandler(cmd *cobra.Command, args []string) {
	if err := override.Edit(); err != nil {
		log.Fatalf("%v", err)
	}
}
