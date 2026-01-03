package battenberg

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	"github.com/leftbin/mactl/internal/installer/pip"
	log "github.com/sirupsen/logrus"
)

const (
	PipPkgPyGit2     = "pygit2"
	PipPkgBattenberg = "battenberg"
	BrewPkgLibSsh2   = "libssh2"
)

func Setup() error {
	log.Infof("installing battenberg")
	if err := pip.UnInstall(PipPkgPyGit2); err != nil {
		return errors.Wrapf(err, "failed to uninstall %s pkg using pip", PipPkgPyGit2)
	}
	if err := brew.Install(BrewPkgLibSsh2); err != nil {
		return errors.Wrapf(err, "failed to install %s pkg using brew", BrewPkgLibSsh2)
	}
	if err := pip.UpgradePip(); err != nil {
		return errors.Wrap(err, "failed to upgrade pip")
	}
	if err := pip.Install(PipPkgPyGit2); err != nil {
		return errors.Wrapf(err, "failed to install %s pkg using pip", PipPkgPyGit2)
	}
	if err := pip.Install(PipPkgBattenberg); err != nil {
		return errors.Wrapf(err, "failed to install %s pkg using pip", PipPkgBattenberg)
	}
	log.Infof("installed battenberg")
	return nil
}
