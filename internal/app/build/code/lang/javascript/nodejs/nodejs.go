package nodejs

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	"github.com/leftbin/mactl/internal/lib/shell"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	BrewPkgNvm         = "nvm"
	BrewPkgCorepack    = "corepack"
	DefaultYarnVersion = "v3.2.3"
	YarnrcYmlFilename  = ".yarnrc.yml"
	YarnrcYmlContents  = `
nodeLinker: node-modules
enableGlobalCache: true
npmScopes:
  "buf":
    npmAlwaysAuth: true
    npmRegistryServer: https://npm.buf.build
    npmAuthToken: "${BUF_TOKEN:-dummy-buf-token}"
`
)

// Setup ensures that all components needed for nodejs development are installed and configured
// installs nvm, yarn and configures yarn
func Setup() error {
	log.Info("installing nvm")
	if err := brew.Install(BrewPkgNvm); err != nil {
		return errors.Wrapf(err, "failed to install nvm")
	}
	log.Info("installed nvm")
	log.Info("ensure yarn")
	if err := ensureYarn(); err != nil {
		return errors.Wrapf(err, "failed to ensure yarn")
	}
	log.Info("ensured yarn")
	return nil
}

func Upgrade() error {
	log.Info("upgrading nvm")
	err := brew.Upgrade(BrewPkgNvm)
	if err != nil {
		return errors.Wrapf(err, "failed to upgrade nvm")
	}
	log.Info("upgraded nvm")
	return nil
}

// ensureYarn ensures yarn is installed and configured
// * installs corepack
// * installs yarn using corepack
// * configures yarn to work with buf.build npm registry
func ensureYarn() error {
	log.Info("installing corepack")
	if err := brew.Install(BrewPkgCorepack); err != nil {
		return errors.Wrapf(err, "failed to install corepack")
	}
	log.Info("installed corepack")
	log.Info("installing yarn")
	//corepack enable
	if err := shell.RunCmd(exec.Command("corepack", "enable")); err != nil {
		return errors.Wrapf(err, "failed to enable corepack")
	}
	///corepack prepare yarn@v3.2.3 --activate
	if err := shell.RunCmd(exec.Command("corepack", "prepare", fmt.Sprintf("yarn@%s", DefaultYarnVersion), "--activate")); err != nil {
		return errors.Wrapf(err, "failed to install %s version of yarn", DefaultYarnVersion)
	}
	log.Info("installed yarn")
	log.Info("configuring yarn")
	if err := ensureYarnrcYmlConfigFile(); err != nil {
		return errors.Wrap(err, "failed to configure")
	}
	log.Info("configured yarn")
	return nil
}

// ensureYarnrcYmlConfigFile ensures .yarnrc.yml file is created in the user home directory.
func ensureYarnrcYmlConfigFile() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrapf(err, "failed to get user home directory")
	}
	yarnrcYmlFilePath := filepath.Join(homeDir, YarnrcYmlFilename)
	if err := os.WriteFile(yarnrcYmlFilePath, []byte(YarnrcYmlContents), os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to write %s file", yarnrcYmlFilePath)
	}
	return nil
}
