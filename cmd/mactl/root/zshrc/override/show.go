package override

import (
	"fmt"
	"github.com/leftbin/mactl/internal/zshrc/override"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Show = &cobra.Command{
	Use:   "show",
	Short: fmt.Sprintf("show contents of ~/%s", override.FileName),
	Run:   showHandler,
}

func showHandler(cmd *cobra.Command, args []string) {
	if err := override.Show(); err != nil {
		if err == override.ErrNotFound {
			log.Infof("no override file %s", override.FileName)
			return
		}
		log.Fatalf("%v", err)
	}
}
