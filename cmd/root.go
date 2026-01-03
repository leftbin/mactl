package mactl

import (
	"fmt"
	"github.com/leftbin/mactl/cmd/mactl/root"
	"github.com/leftbin/mactl/internal/cli/flag"
	intos "github.com/leftbin/mactl/internal/lib/os"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var debug bool

var rootCmd = &cobra.Command{
	Use:   "mactl",
	Short: "tool to manage macbook",
	Long:  `mactl simplifies macbook management`,
}

const HomebrewAppleSiliconBinPath = "/opt/homebrew/bin"

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, string(flag.Debug), false, "set log level to debug")
	rootCmd.DisableSuggestions = true
	cobra.OnInitialize(func() {
		if debug {
			log.SetLevel(log.DebugLevel)
			log.Debug("running in debug mode")
		}
		if intos.IsArmArch() {
			pathEnvVal := os.Getenv("PATH")
			if err := os.Setenv("PATH", fmt.Sprintf("%s:%s", pathEnvVal, HomebrewAppleSiliconBinPath)); err != nil {
				log.Fatalf("failed to set PATH env. err: %v", err)
			}
		}
	})
    rootCmd.AddCommand(
        root.App,
        root.Bootstrap,
        root.Bundle,
        root.EnvVar,
        root.Git,
        // Keyboard shortcuts moved to karayaml; command removed from mactl
        root.MergeFiles,
        root.Optimize,
        root.Version,
        root.Zshrc,
    )
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
