package flux

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	log "github.com/sirupsen/logrus"
)

const (
	BrewCaskPkg = "flux"
)

func Setup() error {
	log.Info("installing flux")
	if err := brew.InstallCask(BrewCaskPkg); err != nil {
		return errors.Wrapf(err, "failed to install %s pkg using brew", BrewCaskPkg)
	}
	log.Info("installed flux")
	return nil
}

func Upgrade() error {
	log.Info("upgrading flux")
	if err := brew.UpgradeCask(BrewCaskPkg); err != nil {
		return errors.Wrapf(err, "failed to upgrade %s pkg using brew", BrewCaskPkg)
	}
	log.Info("upgraded flux")
	return nil
}
