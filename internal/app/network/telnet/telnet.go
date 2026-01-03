package telnet

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	log "github.com/sirupsen/logrus"
)

const (
	BrewPkg = "telnet"
)

func Setup() error {
	log.Info("installing telnet")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to install %s pkg using brew", BrewPkg)
	}
	log.Info("installed telnet")
	return nil
}

func Upgrade() error {
	log.Info("upgrading telnet")
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to upgrade %s pkg using brew", BrewPkg)
	}
	log.Info("upgraded telnet")
	return nil
}
