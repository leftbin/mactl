package awscli

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/installer/brew"
)

const (
	BrewPkg = "awscli"
)

func Setup() error {
	log.Info("installing awscli")
	err := brew.Install(BrewPkg)
	if err != nil {
		return errors.Wrapf(err, "failed to install %s pkg", BrewPkg)
	}
	log.Info("installed awscli")
	return nil
}

func Upgrade() error {
	log.Info("upgrading awscli")
	err := brew.Upgrade(BrewPkg)
	if err != nil {
		return errors.Wrapf(err, "failed to upgrade %s pkg", BrewPkg)
	}
	log.Info("upgraded awscli")
	return nil
}
