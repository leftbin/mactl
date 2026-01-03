package step

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	log "github.com/sirupsen/logrus"
)

const (
	BrewPkg = "step"
)

func Setup() error {
	log.Info("installing step")
	if err :=
		brew.Install(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to install %s pkg using brew", BrewPkg)
	}
	log.Info("installed step")
	return nil
}

func Upgrade() error {
	log.Info("upgrading step")
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to upgrade %s pkg using brew", BrewPkg)
	}
	log.Info("upgraded step")
	return nil
}
