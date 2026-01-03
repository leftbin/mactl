package packer

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	log "github.com/sirupsen/logrus"
)

const (
	BrewPkg = "packer"
)

func Setup() error {
	log.Info("installing packer")
	err := brew.Install(BrewPkg)
	if err != nil {
		return errors.Wrap(err, "failed to install packer")
	}
	log.Info("installed packer")
	return nil
}

func Upgrade() error {
	log.Info("upgrading packer")
	err := brew.Upgrade(BrewPkg)
	if err != nil {
		return errors.Wrap(err, "failed to upgrade packer")
	}
	log.Info("upgraded packer")
	return nil
}
