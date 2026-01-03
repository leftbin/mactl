package root

import (
	"github.com/spf13/cobra"
	"github.com/leftbin/mactl/cmd/mactl/root/app"
)

var App = &cobra.Command{
	Use:   "app",
	Short: "app management",
}

func init() {
	App.AddCommand(app.List, app.Add, app.Upgrade)
}
