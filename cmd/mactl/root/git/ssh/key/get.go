package key

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/leftbin/mactl/internal/cli/flag"
	"github.com/leftbin/mactl/internal/git/ssh/key"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Get = &cobra.Command{
	Use:   "get",
	Short: "get ssh public key",
	Run:   getHandler,
}

func init() {
	Get.PersistentFlags().BoolP(string(flag.Copy), "", false, "copy public key to clipboard?")
}

func getHandler(cmd *cobra.Command, args []string) {
	host, err := cmd.InheritedFlags().GetString(string(flag.Host))
	flag.HandleFlagErrAndVal(err, flag.Host, host)
	workspace, err := cmd.InheritedFlags().GetString(string(flag.Workspace))
	flag.HandleFlagErr(err, flag.Host)
	isCopyToClipBoard, err := cmd.PersistentFlags().GetBool(string(flag.Copy))
	flag.HandleFlagErr(err, flag.Copy)
	pubKey, err := key.Get(host, workspace)
	if err != nil {
		log.Fatal(err)
	}
	if !isCopyToClipBoard {
		fmt.Print(pubKey)
		return
	}
	if err := clipboard.WriteAll(pubKey); err != nil {
		log.Fatal("failed to copy pub key to clipboard")
	}
	key.HandlePubKeyUpdOnScm(host)
}
