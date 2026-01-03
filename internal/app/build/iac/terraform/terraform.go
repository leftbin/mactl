package terraform

import (
	"github.com/leftbin/go-util/pkg/file"
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

const (
	BrewPkg         = "terraform"
	BrewPkgTfSwitch = "warrensbox/tap/tfswitch"
	RcFile          = `
plugin_cache_dir   = "$HOME/.terraform.d/plugin-cache"
disable_checkpoint = true
`
)

func Setup() error {
	log.Info("installing tfswitch")
	err := brew.Install(BrewPkgTfSwitch)
	if err != nil {
		return errors.Wrap(err, "failed to install tfswitch")
	}
	log.Info("installed tfswitch")
	log.Info("installing terraform")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to install terraform")
	}
	log.Info("installed terraform")
	log.Info("ensuring terraform plugin cache")
	if err := setupPluginCache(); err != nil {
		return errors.Wrapf(err, "failed to ensure plugin cache")
	}
	log.Info("ensured terraform plugin cache")
	return nil
}

func setupPluginCache() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get home dir")
	}
	tfrcFile := filepath.Join(homeDir, ".terraformrc")
	if err := os.WriteFile(tfrcFile, []byte(RcFile), os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to write %s file", tfrcFile)
	}
	pluginCacheDir := filepath.Join(homeDir, ".terraform.d", "plugin-cache")
	if !file.IsDirExists(pluginCacheDir) {
		if err := os.MkdirAll(pluginCacheDir, os.ModePerm); err != nil {
			return errors.Wrapf(err, "failed to ensure %s dir", pluginCacheDir)
		}
	}
	return nil
}

func Upgrade() error {
	log.Info("upgrading terraform")
	err := brew.Upgrade(BrewPkg)
	if err != nil {
		return errors.Wrap(err, "failed to upgrade terraform")
	}
	log.Info("upgraded terraform")
	return nil
}
