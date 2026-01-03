package ssh

import (
	"github.com/spf13/cobra"
	"github.com/leftbin/mactl/cmd/mactl/root/git/config"
	"github.com/leftbin/mactl/cmd/mactl/root/git/ssh/key"
	"github.com/leftbin/mactl/internal/cli/flag"
)

var Key = &cobra.Command{
	Use:   "key",
	Short: "ssh key management",
}

func init() {
	Key.PersistentFlags().StringP(string(flag.Host), "", "gitlab.com", "git host. ex: github.com")
	Key.PersistentFlags().StringP(string(flag.Workspace), "", "leftbin", "git workspace. this can be user profile name or org/group name. ex: swarupdonepudi or some-github-org-name")
	flag.MarkFlagsRequired(Key, flag.Host)
	Key.AddCommand(key.Cre, key.Get, key.Use, config.Init)
}
