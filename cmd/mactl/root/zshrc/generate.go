package zshrc

import (
	"github.com/leftbin/mactl/internal/cli/flag"
	"github.com/leftbin/mactl/internal/zshrc/generate"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Generate = &cobra.Command{
	Use:   "generate",
	Short: "generate zshrc",
	Run:   generateHandler,
}

func init() {
	Generate.PersistentFlags().BoolP(string(flag.Show), "", false, "show generated zshrc?")
}

func generateHandler(cmd *cobra.Command, args []string) {
	isShow, err := cmd.PersistentFlags().GetBool(string(flag.Show))
	flag.HandleFlagErr(err, flag.Show)
	if err := generate.Create(); err != nil {
		log.Fatalf("%v", err)
	}
	if !isShow {
		return
	}
	if err := generate.Show(); err != nil {
		log.Fatalf("%v", err)
	}
}
