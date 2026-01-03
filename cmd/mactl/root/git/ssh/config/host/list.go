package host

import (
	"github.com/leftbin/mactl/internal/git/ssh/config/host"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var List = &cobra.Command{
	Use:   "list",
	Short: "list git ssh config hosts",
	Run:   listHandler,
}

func listHandler(cmd *cobra.Command, args []string) {
	hosts, err := host.List()
	if err != nil {
		log.Fatalf("failed. err: %v", err)
	}
	host.PrintList(hosts)
}
