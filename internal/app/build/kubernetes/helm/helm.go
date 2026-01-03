package helm

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	log "github.com/sirupsen/logrus"
)

const (
	BrewPkg = "helm"
)

func Setup() error {
	log.Info("installing helm")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to install helm")
	}
	log.Info("installed helm")
	return nil
}

func Upgrade() error {
	log.Info("upgrading helm")
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to upgrade helm")
	}
	log.Info("upgraded helm")
	return nil
}
