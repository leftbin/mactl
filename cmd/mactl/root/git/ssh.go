package git

import (
	"github.com/spf13/cobra"
	"github.com/leftbin/mactl/cmd/mactl/root/git/ssh"
	"github.com/leftbin/mactl/internal/cli/flag"
)

var Ssh = &cobra.Command{
	Use:   "ssh",
	Short: "git ssh management",
}

func init() {
	Ssh.PersistentFlags().StringP(string(flag.Host), "", "gitlab.com", "git host. ex: github.com")
	Ssh.PersistentFlags().StringP(string(flag.Workspace), "", "plantoncode", "git workspace. this can be user profile name or org/group name. ex: swarupdonepudi or some-github-org-name")
	Ssh.AddCommand(ssh.Init, ssh.Key, ssh.Config)
}
