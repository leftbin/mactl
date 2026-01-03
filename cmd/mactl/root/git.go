package root

import (
	"github.com/leftbin/mactl/cmd/mactl/root/git"
	"github.com/spf13/cobra"
)

var Git = &cobra.Command{
	Use:   "git",
	Short: "git management",
}

func init() {
	Git.AddCommand(git.Config, git.Ssh)
}
