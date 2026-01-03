package git

import (
	"github.com/leftbin/mactl/cmd/mactl/root/git/config"
	"github.com/leftbin/mactl/internal/cli/flag"
	"github.com/spf13/cobra"
)

var Config = &cobra.Command{
	Use:   "config",
	Short: "config management",
}

func init() {
	Config.PersistentFlags().StringP(string(flag.Host), "", "gitlab.com", "git host. ex: github.com")
	Config.PersistentFlags().StringP(string(flag.Workspace), "", "leftbin", "git workspace. this can be user profile name or org/group name. ex: swarupdonepudi or some-github-org-name")
	Config.AddCommand(config.Get, config.Set, config.Del, config.Use, config.Init)
}
