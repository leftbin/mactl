package key

import (
	"github.com/leftbin/mactl/internal/cli/flag"
	"github.com/leftbin/mactl/internal/git/ssh/key"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Use = &cobra.Command{
	Use:   "use",
	Short: "change ssh key to be used for selected host and workspace",
	Run:   useHandler,
}

func useHandler(cmd *cobra.Command, args []string) {
	host, err := cmd.InheritedFlags().GetString(string(flag.Host))
	flag.HandleFlagErrAndVal(err, flag.Host, host)
	workspace, err := cmd.InheritedFlags().GetString(string(flag.Workspace))
	flag.HandleFlagErr(err, flag.Host)
	if err := key.Use(host, workspace); err != nil {
		log.Fatalf("%v", err)
	}
	log.Info("success")
}
