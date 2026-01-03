package bootstrap

import (
	"github.com/spf13/cobra"
	"github.com/leftbin/mactl/internal/bootstrap"
)

var Checklist = &cobra.Command{
	Use:   "checklist",
	Short: "see checklist for bootstrap",
	Run:   checklistHandler,
}

func checklistHandler(cmd *cobra.Command, args []string) {
	bootstrap.Checklist()
}
