package key

import (
	"github.com/leftbin/mactl/internal/cli/flag"
	"github.com/leftbin/mactl/internal/git/ssh/key"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cre = &cobra.Command{
	Use:   "cre",
	Short: "cre ssh key",
	Run:   creHandler,
}

func creHandler(cmd *cobra.Command, args []string) {
	host, err := cmd.InheritedFlags().GetString(string(flag.Host))
	flag.HandleFlagErrAndVal(err, flag.Host, host)
	workspace, err := cmd.InheritedFlags().GetString(string(flag.Workspace))
	flag.HandleFlagErr(err, flag.Workspace)
	keyPath, err := key.Cre(host, workspace)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("key location: %s", keyPath)
}
