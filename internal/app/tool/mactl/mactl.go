package mactl

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	log "github.com/sirupsen/logrus"
)

const (
	BrewPkg = "leftbin/homebrew-tap/mactl"
)

func Setup() error {
	log.Info("installing mactl")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to install %s pkg using brew", BrewPkg)
	}
	log.Info("installed mactl")
	return nil
}

func Upgrade() error {
	log.Info("upgrading mactl")
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to upgrade %s pkg using brew", BrewPkg)
	}
	log.Info("upgraded mactl")
	return nil
}
